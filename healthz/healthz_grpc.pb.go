// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: healthz/healthz.proto

package healthz

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

// HealthzClient is the client API for Healthz service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthzClient interface {
	// Get will get health status for a gNMI path.  If no status is available for
	// the requested path an error will be returned.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type healthzClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthzClient(cc grpc.ClientConnInterface) HealthzClient {
	return &healthzClient{cc}
}

func (c *healthzClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/gnoi.healthz.Healthz/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthzServer is the server API for Healthz service.
// All implementations must embed UnimplementedHealthzServer
// for forward compatibility
type HealthzServer interface {
	// Get will get health status for a gNMI path.  If no status is available for
	// the requested path an error will be returned.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedHealthzServer()
}

// UnimplementedHealthzServer must be embedded to have forward compatible implementations.
type UnimplementedHealthzServer struct {
}

func (UnimplementedHealthzServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHealthzServer) mustEmbedUnimplementedHealthzServer() {}

// UnsafeHealthzServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthzServer will
// result in compilation errors.
type UnsafeHealthzServer interface {
	mustEmbedUnimplementedHealthzServer()
}

func RegisterHealthzServer(s grpc.ServiceRegistrar, srv HealthzServer) {
	s.RegisterService(&Healthz_ServiceDesc, srv)
}

func _Healthz_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthzServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.healthz.Healthz/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthzServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Healthz_ServiceDesc is the grpc.ServiceDesc for Healthz service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Healthz_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.healthz.Healthz",
	HandlerType: (*HealthzServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Healthz_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "healthz/healthz.proto",
}
