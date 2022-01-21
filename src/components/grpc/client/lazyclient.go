package client

import (
	"context"
	"flag"
	"sync"

	log "github.com/golang/glog"
	"google.golang.org/grpc"
)

var (
	disableLazyInit = flag.Bool("disableLazyInit", false, "Disable lazy init, default: false")
)

func NewGrpcLazyClinet(serverAddr string) (*GrpcLazyClient, error) {
	lazyClient := &GrpcLazyClient{
		serverAddr: serverAddr,
	}
	if *disableLazyInit {
		if err := lazyClient.initOnce(); err != nil {
			return nil, err
		}
	} else {
		log.Infof("use lazy client for addr:%s, the actual client may not initialized yet!\n", serverAddr)
	}
	return lazyClient, nil
}

type GrpcLazyClient struct {
	*GrpcClient

	serverAddr string
	once       sync.Once
}

func (l *GrpcLazyClient) initOnce() error {
	var err error
	l.once.Do(func() {
		if l.GrpcClient == nil {
			l.GrpcClient, err = NewGrpcClient(l.serverAddr)
		}
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
