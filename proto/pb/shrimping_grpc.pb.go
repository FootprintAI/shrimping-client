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

// ShrimpingClient is the client API for Shrimping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShrimpingClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	LookupItemByIDs(ctx context.Context, in *LookupItemByIDsRequest, opts ...grpc.CallOption) (*RawResponse, error)
	LookupItemBySales(ctx context.Context, in *LookupItemBySalesRequest, opts ...grpc.CallOption) (*RawResponse, error)
}

type shrimpingClient struct {
	cc grpc.ClientConnInterface
}

func NewShrimpingClient(cc grpc.ClientConnInterface) ShrimpingClient {
	return &shrimpingClient{cc}
}

func (c *shrimpingClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/pb.shrimping/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shrimpingClient) LookupItemByIDs(ctx context.Context, in *LookupItemByIDsRequest, opts ...grpc.CallOption) (*RawResponse, error) {
	out := new(RawResponse)
	err := c.cc.Invoke(ctx, "/pb.shrimping/LookupItemByIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shrimpingClient) LookupItemBySales(ctx context.Context, in *LookupItemBySalesRequest, opts ...grpc.CallOption) (*RawResponse, error) {
	out := new(RawResponse)
	err := c.cc.Invoke(ctx, "/pb.shrimping/LookupItemBySales", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShrimpingServer is the server API for Shrimping service.
// All implementations must embed UnimplementedShrimpingServer
// for forward compatibility
type ShrimpingServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	LookupItemByIDs(context.Context, *LookupItemByIDsRequest) (*RawResponse, error)
	LookupItemBySales(context.Context, *LookupItemBySalesRequest) (*RawResponse, error)
	mustEmbedUnimplementedShrimpingServer()
}

// UnimplementedShrimpingServer must be embedded to have forward compatible implementations.
type UnimplementedShrimpingServer struct {
}

func (UnimplementedShrimpingServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedShrimpingServer) LookupItemByIDs(context.Context, *LookupItemByIDsRequest) (*RawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupItemByIDs not implemented")
}
func (UnimplementedShrimpingServer) LookupItemBySales(context.Context, *LookupItemBySalesRequest) (*RawResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LookupItemBySales not implemented")
}
func (UnimplementedShrimpingServer) mustEmbedUnimplementedShrimpingServer() {}

// UnsafeShrimpingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShrimpingServer will
// result in compilation errors.
type UnsafeShrimpingServer interface {
	mustEmbedUnimplementedShrimpingServer()
}

func RegisterShrimpingServer(s grpc.ServiceRegistrar, srv ShrimpingServer) {
	s.RegisterService(&Shrimping_ServiceDesc, srv)
}

func _Shrimping_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShrimpingServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.shrimping/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShrimpingServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shrimping_LookupItemByIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupItemByIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShrimpingServer).LookupItemByIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.shrimping/LookupItemByIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShrimpingServer).LookupItemByIDs(ctx, req.(*LookupItemByIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shrimping_LookupItemBySales_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupItemBySalesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShrimpingServer).LookupItemBySales(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.shrimping/LookupItemBySales",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShrimpingServer).LookupItemBySales(ctx, req.(*LookupItemBySalesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Shrimping_ServiceDesc is the grpc.ServiceDesc for Shrimping service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shrimping_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.shrimping",
	HandlerType: (*ShrimpingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Shrimping_Login_Handler,
		},
		{
			MethodName: "LookupItemByIDs",
			Handler:    _Shrimping_LookupItemByIDs_Handler,
		},
		{
			MethodName: "LookupItemBySales",
			Handler:    _Shrimping_LookupItemBySales_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shrimping.proto",
}