package client

import (
	"context"
	"sync"

	log "github.com/golang/glog"
	"google.golang.org/grpc"
)

func NewGrpcLazyClinet(serverAddr string) (*GrpcLazyClient, error) {
	log.Infof("use lazy client for addr:%s, the actual client may not initialized yet!\n", serverAddr)
	return &GrpcLazyClient{
		serverAddr: serverAddr,
	}, nil
}

type GrpcLazyClient struct {
	*GrpcClient

	serverAddr string
	once       sync.Once
}

func (l *GrpcLazyClient) initOnce() error {
	var err error
	l.once.Do(func() {
		l.GrpcClient, err = NewGrpcClient(l.serverAddr)
	})
	if err != nil {
		return err
	}
	return nil
}

func (l *GrpcLazyClient) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	if err := l.initOnce(); err != nil {
		return err
	}
	return l.GrpcClient.Invoke(ctx, method, args, reply, opts...)
}

func (l *GrpcLazyClient) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	if err := l.initOnce(); err != nil {
		return nil, err
	}
	return l.GrpcClient.NewStream(ctx, desc, method, opts...)
}

func (l *GrpcLazyClient) Close() error {
	if l.GrpcClient == nil {
		return nil
	}
	return l.GrpcClient.Close()
}
