// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShrimpingInstagramClient is the client API for ShrimpingInstagram service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShrimpingInstagramClient interface {
	Profile(ctx context.Context, in *InstagramRequest, opts ...grpc.CallOption) (ShrimpingInstagram_ProfileClient, error)
	Posts(ctx context.Context, in *InstagramRequest, opts ...grpc.CallOption) (ShrimpingInstagram_PostsClient, error)
}

type shrimpingInstagramClient struct {
	cc grpc.ClientConnInterface
}

func NewShrimpingInstagramClient(cc grpc.ClientConnInterface) ShrimpingInstagramClient {
	return &shrimpingInstagramClient{cc}
}

func (c *shrimpingInstagramClient) Profile(ctx context.Context, in *InstagramRequest, opts ...grpc.CallOption) (ShrimpingInstagram_ProfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &ShrimpingInstagram_ServiceDesc.Streams[0], "/pb.shrimpingInstagram/Profile", opts...)
	if err != nil {
		return nil, err
	}
	x := &shrimpingInstagramProfileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ShrimpingInstagram_ProfileClient interface {
	Recv() (*InstagramResponse, error)
	grpc.ClientStream
}

type shrimpingInstagramProfileClient struct {
	grpc.ClientStream
}

func (x *shrimpingInstagramProfileClient) Recv() (*InstagramResponse, error) {
	m := new(InstagramResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *shrimpingInstagramClient) Posts(ctx context.Context, in *InstagramRequest, opts ...grpc.CallOption) (ShrimpingInstagram_PostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ShrimpingInstagram_ServiceDesc.Streams[1], "/pb.shrimpingInstagram/Posts", opts...)
	if err != nil {
		return nil, err
	}
	x := &shrimpingInstagramPostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ShrimpingInstagram_PostsClient interface {
	Recv() (*InstagramResponse, error)
	grpc.ClientStream
}

type shrimpingInstagramPostsClient struct {
	grpc.ClientStream
}

func (x *shrimpingInstagramPostsClient) Recv() (*InstagramResponse, error) {
	m := new(InstagramResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ShrimpingInstagramServer is the server API for ShrimpingInstagram service.
// All implementations must embed UnimplementedShrimpingInstagramServer
// for forward compatibility
type ShrimpingInstagramServer interface {
	Profile(*InstagramRequest, ShrimpingInstagram_ProfileServer) error
	Posts(*InstagramRequest, ShrimpingInstagram_PostsServer) error
	mustEmbedUnimplementedShrimpingInstagramServer()
}

// UnimplementedShrimpingInstagramServer must be embedded to have forward compatible implementations.
type UnimplementedShrimpingInstagramServer struct {
}

func (UnimplementedShrimpingInstagramServer) Profile(*InstagramRequest, ShrimpingInstagram_ProfileServer) error {
	return status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (UnimplementedShrimpingInstagramServer) Posts(*InstagramRequest, ShrimpingInstagram_PostsServer) error {
	return status.Errorf(codes.Unimplemented, "method Posts not implemented")
}
func (UnimplementedShrimpingInstagramServer) mustEmbedUnimplementedShrimpingInstagramServer() {}

// UnsafeShrimpingInstagramServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShrimpingInstagramServer will
// result in compilation errors.
type UnsafeShrimpingInstagramServer interface {
	mustEmbedUnimplementedShrimpingInstagramServer()
}

func RegisterShrimpingInstagramServer(s grpc.ServiceRegistrar, srv ShrimpingInstagramServer) {
	s.RegisterService(&ShrimpingInstagram_ServiceDesc, srv)
}

func _ShrimpingInstagram_Profile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InstagramRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShrimpingInstagramServer).Profile(m, &shrimpingInstagramProfileServer{stream})
}

type ShrimpingInstagram_ProfileServer interface {
	Send(*InstagramResponse) error
	grpc.ServerStream
}

type shrimpingInstagramProfileServer struct {
	grpc.ServerStream
}

func (x *shrimpingInstagramProfileServer) Send(m *InstagramResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ShrimpingInstagram_Posts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InstagramRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShrimpingInstagramServer).Posts(m, &shrimpingInstagramPostsServer{stream})
}

type ShrimpingInstagram_PostsServer interface {
	Send(*InstagramResponse) error
	grpc.ServerStream
}

type shrimpingInstagramPostsServer struct {
	grpc.ServerStream
}

func (x *shrimpingInstagramPostsServer) Send(m *InstagramResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ShrimpingInstagram_ServiceDesc is the grpc.ServiceDesc for ShrimpingInstagram service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShrimpingInstagram_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.shrimpingInstagram",
	HandlerType: (*ShrimpingInstagramServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Profile",
			Handler:       _ShrimpingInstagram_Profile_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Posts",
			Handler:       _ShrimpingInstagram_Posts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "shritagram.proto",
}
