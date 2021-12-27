package client

import (
	"context"
	"fmt"
	"time"

	//log "github.com/golang/glog"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"

	"github.com/footprintai/shrimping/components/grpc/client/credentials"
)

type GrpcConnCloser interface {
	Close() error
}

func NewGrpcClient(serverAddr string) (*GrpcClient, error) {
	creds := credentials.NewTLSCredentials()
	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100 * time.Millisecond)),
		grpc_retry.WithCodes(codes.NotFound, codes.Aborted, codes.Unavailable),
	}
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(creds),
		//grpc.FailOnNonTempDialError(true), // this will disable non-temporary error retry, like server is donw
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(retryOpts...)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)),
	}
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, err
	}
	return &GrpcClient{ClientConn: conn}, nil
}

type GrpcClient struct {
	*grpc.ClientConn
}

func (g *GrpcClient) WaitForReconnect(ctx context.Context) error {
	// try to connect and wait for its change
	for {
		g.ClientConn.Connect()
		s := g.ClientConn.GetState()
		if s != connectivity.Ready {
			if !g.ClientConn.WaitForStateChange(ctx, s) {
				return fmt.Errorf("grpc: timeout when waiting for state change")
			}
		} else {
			return nil
		}
	}
}
