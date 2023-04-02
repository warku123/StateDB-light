// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.7
// source: sdbrpc.proto

package sdbrpc

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

const (
	Sdbrpc_Ping_FullMethodName = "/sdbrpc.Sdbrpc/Ping"
)

// SdbrpcClient is the client API for Sdbrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SdbrpcClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type sdbrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewSdbrpcClient(cc grpc.ClientConnInterface) SdbrpcClient {
	return &sdbrpcClient{cc}
}

func (c *sdbrpcClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Sdbrpc_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SdbrpcServer is the server API for Sdbrpc service.
// All implementations must embed UnimplementedSdbrpcServer
// for forward compatibility
type SdbrpcServer interface {
	Ping(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedSdbrpcServer()
}

// UnimplementedSdbrpcServer must be embedded to have forward compatible implementations.
type UnimplementedSdbrpcServer struct {
}

func (UnimplementedSdbrpcServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSdbrpcServer) mustEmbedUnimplementedSdbrpcServer() {}

// UnsafeSdbrpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SdbrpcServer will
// result in compilation errors.
type UnsafeSdbrpcServer interface {
	mustEmbedUnimplementedSdbrpcServer()
}

func RegisterSdbrpcServer(s grpc.ServiceRegistrar, srv SdbrpcServer) {
	s.RegisterService(&Sdbrpc_ServiceDesc, srv)
}

func _Sdbrpc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SdbrpcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sdbrpc_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SdbrpcServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Sdbrpc_ServiceDesc is the grpc.ServiceDesc for Sdbrpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sdbrpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sdbrpc.Sdbrpc",
	HandlerType: (*SdbrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Sdbrpc_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sdbrpc.proto",
}
