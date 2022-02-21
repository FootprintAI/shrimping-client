package instagramclient

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

func NewClient(serverAddr string, apiKey string, timeout time.Duration, callOptions ...grpc.CallOption) (*Client, error) {
	clientConn, err := client.NewGrpcClient(serverAddr)
	if err != nil {
		return nil, err
	}
	return &Client{
		pbClient:   pb.NewShritagramClient(clientConn),
		verClient:  pb.NewVersioningClient(clientConn),
		crudClient: pb.NewWebhookCRUDClient(clientConn),
		callOpts:   callOptions,
		conn:       clientConn,
		apiKey:     apiKey,
		timeout:    timeout,
	}, nil
}

type Client struct {
	pbClient   pb.ShritagramClient
	verClient  pb.VersioningClient
	crudClient pb.WebhookCRUDClient
	callOpts   []grpc.CallOption
	conn       *client.GrpcClient
	timeout    time.Duration
	apiKey     string
}

func (c *Client) TestWebhook() error {
	ctx := context.Background()
	_, err := c.crudClient.TestWebhook(ctx, &pb.TestWebhookRequest{})
	return err
}

func (c *Client) VerifySignature(host, hmackey string, bodydigest string, createdTs *int32) (signStr string, signature string, headers string, err error) {
	ctx := context.Background()
	resp, err := c.crudClient.VerifySignature(c.ctxWithToken(ctx), &pb.VerifySignatureRequest{
		Host:       host,
		Hmackey:    hmackey,
		BodyDigest: bodydigest,
		CreatedTs:  createdTs,
	})
	if err != nil {
		return "", "", "", err
	}
	return resp.SignStr, resp.Signature, resp.HttpHeaders, nil

}

func (c *Client) SubscribeWebhook(webhookHost string, forceHttps bool) (apikey, secret, host string, err error) {
	ctx := context.Background()
	resp, err := c.crudClient.SubscribeWebhook(c.ctxWithToken(ctx), &pb.SubscribeWebhookRequest{
		WebhookAddress: webhookHost,
		ForceHttps:     &forceHttps,
	})
	if err != nil {
		return "", "", "", err
	}
	return resp.Data.ApiKey, resp.Data.SecretKey, resp.Data.Address, nil
}

type WebhookResponse struct {
	ApiKey         string
	ApiSecret      string
	WebhookAddress string
}

func (c *Client) ListWebhook() ([]*WebhookResponse, error) {
	ctx := context.Background()
	resp, err := c.crudClient.ListWebhook(c.ctxWithToken(ctx), &pb.ListWebhookRequest{})
	if err != nil {
		return nil, err
	}
	var respList []*WebhookResponse
	for _, data := range resp.List {
		respList = append(respList, &WebhookResponse{
			ApiKey:         data.ApiKey,
			ApiSecret:      data.SecretKey,
			WebhookAddress: data.Address,
		})
	}
	return respList, nil
}

func (c *Client) DeleteWebhook(apikeys []string) error {
	ctx := context.Background()
	_, err := c.crudClient.DeleteWebhook(c.ctxWithToken(ctx), &pb.DeleteWebhookRequest{ApiKeys: apikeys})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Callback(cb func(*pb.ShritagramResponse) error) error {
	ctx := context.Background()
	for {
		stream, err := c.pbClient.Callback(c.ctxWithToken(ctx), &pb.CallbackRequest{}, c.callOpts...)
		if err != nil {
			log.Errorf("get pbclient failed, err:%s", err)
			return errors.ParseError(err)
		}
		for {
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

func (c *Client) Profile(usernames []string, cachePolicy string, priority string) ([]*pb.ShritagramResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.pbClient.Profile(c.ctxWithToken(ctx), &pb.ProfileRequest{
		Usernames:    usernames,
		CacheControl: cachePolicy,
		Priority:     priority,
	}, c.callOpts...)
	if err != nil {
		return nil, errors.ParseError(err)
	}
	return []*pb.ShritagramResponse{decompress(resp)}, nil
}

func decompress(out *pb.ShritagramResponse) *pb.ShritagramResponse {
	if out.Compression != pb.ObjectCompressionAlgorithm_GZIP {
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

func (c *Client) Posts(shortcodes []string, cachePolicy string, priority string) ([]*pb.ShritagramResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.pbClient.Posts(c.ctxWithToken(ctx), &pb.PostRequest{
		Shortcodes:   shortcodes,
		CacheControl: cachePolicy,
		Priority:     priority,
	}, c.callOpts...)
	if err != nil {
		return nil, errors.ParseError(err)
	}
	return []*pb.ShritagramResponse{decompress(resp)}, nil
}

func (c *Client) TopSearch(hashtags []string, priority string) ([]*pb.ShritagramResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	out, err := c.pbClient.TopSearch(c.ctxWithToken(ctx), &pb.TopSearchRequest{
		Hashtags: hashtags,
		Priority: priority,
	}, c.callOpts...)
	if err != nil {
		return nil, errors.ParseError(err)
	}
	return []*pb.ShritagramResponse{decompress(out)}, nil
}

func (c *Client) Version() (serverVersion string, isVersionCompatible bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	resp, err := c.verClient.Version(
		c.ctxWithToken(ctx),
		&pb.VersionRequest{
			ClientVersion: version.GetVersion(),
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
