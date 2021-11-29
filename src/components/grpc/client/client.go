package client

import (
	"time"

	//log "github.com/golang/glog"

	"google.golang.org/grpc"

	"github.com/footprintai/shrimping/components/grpc/client/credentials"
)

func NewGrpcClient(serverAddr string) (*GrpcClient, error) {
	creds := credentials.NewTLSCredentials()
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(creds),
		grpc.FailOnNonTempDialError(true),
		grpc.WithTimeout(10 * time.Second),
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
