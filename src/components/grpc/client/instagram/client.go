package shopeeclient

import (
	"context"
	"fmt"
	"io"
	"time"

	log "github.com/golang/glog"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/footprintai/shrimping/components/grpc/client"
	"github.com/footprintai/shrimping/components/grpc/compression"
	errors "github.com/footprintai/shrimping/components/grpc/errors"
	pb "github.com/footprintai/shrimping/components/grpc/proto/pb"
	"github.com/footprintai/shrimping/components/grpc/version"
)

func NewClient(serverAddr string, apiKey string, timeout time.Duration, cachePolicy string, callOptions ...grpc.CallOption) (*Client, error) {
	clientConn, err := client.NewGrpcClient(serverAddr)
	if err != nil {
		return nil, err
	}
	return &Client{
		pbClient:    pb.NewShrimpingInstagramClient(clientConn),
		verClient:   pb.NewVersioningClient(clientConn),
		callOpts:    callOptions,
		conn:        clientConn,
		apiKey:      apiKey,
		timeout:     timeout,
		cachePolicy: cachePolicy,
	}, nil
}

type Client struct {
	pbClient    pb.ShrimpingInstagramClient
	verClient   pb.VersioningClient
	callOpts    []grpc.CallOption
	conn        *client.GrpcClient
	timeout     time.Duration
	apiKey      string
	cachePolicy string
}

func (c *Client) Callback(cb func(*pb.InstagramResponse) error) error {
	ctx := context.Background()
	for {
		stream, err := c.pbClient.Callback(c.ctxWithToken(ctx), &pb.CallbackRequest{}, c.callOpts...)
		if err != nil {
			log.Errorf("get pbclient failed, err:%s", err)
			return errors.ParseError(err)
		}
		for {
			log.Info("callback: start receiving")
			resp, err := stream.Recv()
			if err == io.EOF {
				log.Info("callback: got eof, reconnecting")
				break
			} else if err != nil {
				log.Infof("callback: got err:%+v, reconnecting", err)
				break
			}
			if err := cb(decompress(resp)); err != nil {
				log.Errorf("callback failed, err:%+s\n", err)
				return errors.ParseError(err)
			}
		}
		if err := c.conn.WaitForReconnect(ctx); err != nil {
			return err
		}
	}
}

func (c *Client) Profile(usernames []string) ([]*pb.InstagramResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.pbClient.Profile(c.ctxWithToken(ctx), &pb.InstagramProfileRequest{
		Usernames:    usernames,
		CacheControl: c.cachePolicy,
	}, c.callOpts...)
	if err != nil {
		return nil, errors.ParseError(err)
	}
	return []*pb.InstagramResponse{decompress(resp)}, nil
}

func decompress(out *pb.InstagramResponse) *pb.InstagramResponse {
	if out.Compression != pb.InstagramObjectCompressionAlgorithm_GZIP {
		return out
	}
	compressor := &compression.GzipEncodeDecoder{}
	for _, rawProfile := range out.RawProfiles {
		if rawProfile != nil {
			rawProfile.RawBytes, _ = compressor.Decode(rawProfile.RawBytes)

		}
	}
	for _, rawPost := range out.RawPosts {
		if rawPost != nil {
			rawPost.RawBytes, _ = compressor.Decode(rawPost.RawBytes)
		}
	}
	for _, rawSearch := range out.RawTopSearchs {
		if rawSearch != nil {
			rawSearch.RawBytes, _ = compressor.Decode(rawSearch.RawBytes)
		}
	}
	return out
}

func (c *Client) ctxWithToken(ctx context.Context) context.Context {
	md := metadata.Pairs("authorization", fmt.Sprintf("X-API-Key %v", c.apiKey))
	nCtx := metautils.NiceMD(md).ToOutgoing(ctx)
	return nCtx
}

func (c *Client) Posts(shortcodes []string) ([]*pb.InstagramResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.pbClient.Posts(c.ctxWithToken(ctx), &pb.InstagramPostRequest{
		Shortcodes:   shortcodes,
		CacheControl: c.cachePolicy,
	}, c.callOpts...)
	if err != nil {
		return nil, errors.ParseError(err)
	}
	return []*pb.InstagramResponse{decompress(resp)}, nil
}

func (c *Client) TopSearch(hashtags []string) ([]*pb.InstagramResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	out, err := c.pbClient.TopSearch(c.ctxWithToken(ctx), &pb.InstagramTopSearchRequest{
		Hashtags: hashtags,
	}, c.callOpts...)
	if err != nil {
		return nil, errors.ParseError(err)
	}
	return []*pb.InstagramResponse{decompress(out)}, nil
}

func (c *Client) Version() (serverVersion string, isVersionCompatible bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	ver := version.GetVersion()
	resp, err := c.verClient.Version(
		c.ctxWithToken(ctx),
		&pb.VersionRequest{
			ClientVersion: ver,
		},
	)
	if err != nil {
		return "", false, err
	}
	return resp.ServerVersion, resp.Compatible, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
