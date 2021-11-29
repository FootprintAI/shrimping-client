package shopeeclient

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/footprintai/shrimping/components/grpc/client"
	"github.com/footprintai/shrimping/components/grpc/compression"
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

func (c *Client) Profile(usernames []string) ([]*pb.RawInstgramObject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	stream, err := c.pbClient.Profile(c.ctxWithToken(ctx), &pb.InstagramRequest{
		Usernames:    usernames,
		CacheControl: c.cachePolicy,
	}, c.callOpts...)
	if err != nil {
		return nil, err
	}
	var out []*pb.RawInstgramObject
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		out = append(out, decompress(resp).RawObjects...)
	}
	return out, nil
}

func decompress(out *pb.InstagramResponse) *pb.InstagramResponse {
	if out.Compression != pb.InstagramObjectCompressionAlgorithm_GZIP {
		return out
	}
	compressor := &compression.GzipEncodeDecoder{}
	for _, rawObject := range out.RawObjects {
		if rawObject != nil {
			if rawObject.RawProfile != nil {
				rawObject.RawProfile.RawBytes, _ = compressor.Decode(rawObject.RawProfile.RawBytes)
			}
			for _, rawPost := range rawObject.RawPosts {
				if rawPost != nil {
					rawPost.RawBytes, _ = compressor.Decode(rawPost.RawBytes)
				}
			}
		}
	}
	return out
}

func (c *Client) ctxWithToken(ctx context.Context) context.Context {
	md := metadata.Pairs("authorization", fmt.Sprintf("api-key %v", c.apiKey))
	nCtx := metautils.NiceMD(md).ToOutgoing(ctx)
	return nCtx
}

func (c *Client) Posts(usernames []string) ([]*pb.RawInstgramObject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	stream, err := c.pbClient.Posts(c.ctxWithToken(ctx), &pb.InstagramRequest{
		Usernames:    usernames,
		CacheControl: c.cachePolicy,
	}, c.callOpts...)
	if err != nil {
		return nil, err
	}
	var out []*pb.RawInstgramObject
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		out = append(out, decompress(resp).RawObjects...)
	}
	return out, nil
}

func (c *Client) TopSearch(keywords []string) ([]*pb.RawInstgramTopSearchObject, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	out, err := c.pbClient.TopSearch(c.ctxWithToken(ctx), &pb.InstagramTopSearchRequest{
		Keywords: keywords,
	}, c.callOpts...)
	if err != nil {
		return nil, err
	}
	return decompressTopSearch(out).RawObjects, nil
}

func decompressTopSearch(out *pb.InstagramTopSearchResponse) *pb.InstagramTopSearchResponse {
	if out.Compression != pb.InstagramObjectCompressionAlgorithm_GZIP {
		return out
	}
	compressor := &compression.GzipEncodeDecoder{}
	for _, rawObject := range out.RawObjects {
		if rawObject != nil {
			rawObject.RawBytes, _ = compressor.Decode(rawObject.RawBytes)
		}
	}
	return out
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
