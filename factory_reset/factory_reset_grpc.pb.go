// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: github.com/openconfig/gnoi/factory_reset/factory_reset.proto

package factory_reset

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
	FactoryReset_Start_FullMethodName = "/gnoi.factory_reset.FactoryReset/Start"
)

// FactoryResetClient is the client API for FactoryReset service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FactoryResetClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
}

type factoryResetClient struct {
	cc grpc.ClientConnInterface
}

func NewFactoryResetClient(cc grpc.ClientConnInterface) FactoryResetClient {
	return &factoryResetClient{cc}
}

func (c *factoryResetClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, FactoryReset_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FactoryResetServer is the server API for FactoryReset service.
// All implementations should embed UnimplementedFactoryResetServer
// for forward compatibility
type FactoryResetServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
}

// UnimplementedFactoryResetServer should be embedded to have forward compatible implementations.
type UnimplementedFactoryResetServer struct {
}

func (UnimplementedFactoryResetServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}

// UnsafeFactoryResetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FactoryResetServer will
// result in compilation errors.
type UnsafeFactoryResetServer interface {
	mustEmbedUnimplementedFactoryResetServer()
}

func RegisterFactoryResetServer(s grpc.ServiceRegistrar, srv FactoryResetServer) {
	s.RegisterService(&FactoryReset_ServiceDesc, srv)
}

func _FactoryReset_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FactoryResetServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FactoryReset_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FactoryResetServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FactoryReset_ServiceDesc is the grpc.ServiceDesc for FactoryReset service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FactoryReset_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.factory_reset.FactoryReset",
	HandlerType: (*FactoryResetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _FactoryReset_Start_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openconfig/gnoi/factory_reset/factory_reset.proto",
}
