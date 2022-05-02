// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: factory_reset/factory_reset.proto

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

// FactoryResetClient is the client API for FactoryReset service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FactoryResetClient interface {
	// The Start RPC allows the Client to instruct the Target to immediately
	// clean all existing state and boot the Target in the same condition as it is
	// shipped from factory. State includes storage, configuration, logs,
	// certificates and licenses.
	//
	// Optionally allows rolling back the OS to the same version shipped from
	// factory.
	//
	// Optionally allows for the Target to zero-fill permanent storage where state
	// data is stored.
	//
	// If any of the optional flags is set but not supported, a gRPC Status with
	// code INVALID_ARGUMENT must be returned with the details value set to a
	// properly populated ResetError message.
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
	err := c.cc.Invoke(ctx, "/gnoi.factory_reset.FactoryReset/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FactoryResetServer is the server API for FactoryReset service.
// All implementations must embed UnimplementedFactoryResetServer
// for forward compatibility
type FactoryResetServer interface {
	// The Start RPC allows the Client to instruct the Target to immediately
	// clean all existing state and boot the Target in the same condition as it is
	// shipped from factory. State includes storage, configuration, logs,
	// certificates and licenses.
	//
	// Optionally allows rolling back the OS to the same version shipped from
	// factory.
	//
	// Optionally allows for the Target to zero-fill permanent storage where state
	// data is stored.
	//
	// If any of the optional flags is set but not supported, a gRPC Status with
	// code INVALID_ARGUMENT must be returned with the details value set to a
	// properly populated ResetError message.
	Start(context.Context, *StartRequest) (*StartResponse, error)
	mustEmbedUnimplementedFactoryResetServer()
}

// UnimplementedFactoryResetServer must be embedded to have forward compatible implementations.
type UnimplementedFactoryResetServer struct {
}

func (UnimplementedFactoryResetServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedFactoryResetServer) mustEmbedUnimplementedFactoryResetServer() {}

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
		FullMethod: "/gnoi.factory_reset.FactoryReset/Start",
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
	Metadata: "factory_reset/factory_reset.proto",
}
