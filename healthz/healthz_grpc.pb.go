// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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
	// List returns all events for the provided component path.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Acknowledge will set the acknowledged field for the event.
	// This is an idempotent operation.
	Acknowledge(ctx context.Context, in *AcknowlegeRequest, opts ...grpc.CallOption) (*AcknowledgeResponse, error)
	// Artifact will stream the artifact contents for the provided artifact id.
	Artifact(ctx context.Context, in *ArtifactRequest, opts ...grpc.CallOption) (Healthz_ArtifactClient, error)
	// Check will invoke the healthz on the provided component path. This RPC
	// can be expensive depending on the vendor implementation.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
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

func (c *healthzClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/gnoi.healthz.Healthz/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthzClient) Acknowledge(ctx context.Context, in *AcknowlegeRequest, opts ...grpc.CallOption) (*AcknowledgeResponse, error) {
	out := new(AcknowledgeResponse)
	err := c.cc.Invoke(ctx, "/gnoi.healthz.Healthz/Acknowledge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthzClient) Artifact(ctx context.Context, in *ArtifactRequest, opts ...grpc.CallOption) (Healthz_ArtifactClient, error) {
	stream, err := c.cc.NewStream(ctx, &Healthz_ServiceDesc.Streams[0], "/gnoi.healthz.Healthz/Artifact", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthzArtifactClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Healthz_ArtifactClient interface {
	Recv() (*ArtifactResponse, error)
	grpc.ClientStream
}

type healthzArtifactClient struct {
	grpc.ClientStream
}

func (x *healthzArtifactClient) Recv() (*ArtifactResponse, error) {
	m := new(ArtifactResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *healthzClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/gnoi.healthz.Healthz/Check", in, out, opts...)
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
	// List returns all events for the provided component path.
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Acknowledge will set the acknowledged field for the event.
	// This is an idempotent operation.
	Acknowledge(context.Context, *AcknowlegeRequest) (*AcknowledgeResponse, error)
	// Artifact will stream the artifact contents for the provided artifact id.
	Artifact(*ArtifactRequest, Healthz_ArtifactServer) error
	// Check will invoke the healthz on the provided component path. This RPC
	// can be expensive depending on the vendor implementation.
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	mustEmbedUnimplementedHealthzServer()
}

// UnimplementedHealthzServer must be embedded to have forward compatible implementations.
type UnimplementedHealthzServer struct {
}

func (UnimplementedHealthzServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHealthzServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedHealthzServer) Acknowledge(context.Context, *AcknowlegeRequest) (*AcknowledgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Acknowledge not implemented")
}
func (UnimplementedHealthzServer) Artifact(*ArtifactRequest, Healthz_ArtifactServer) error {
	return status.Errorf(codes.Unimplemented, "method Artifact not implemented")
}
func (UnimplementedHealthzServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
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

func _Healthz_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthzServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.healthz.Healthz/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthzServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Healthz_Acknowledge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcknowlegeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthzServer).Acknowledge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.healthz.Healthz/Acknowledge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthzServer).Acknowledge(ctx, req.(*AcknowlegeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Healthz_Artifact_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ArtifactRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HealthzServer).Artifact(m, &healthzArtifactServer{stream})
}

type Healthz_ArtifactServer interface {
	Send(*ArtifactResponse) error
	grpc.ServerStream
}

type healthzArtifactServer struct {
	grpc.ServerStream
}

func (x *healthzArtifactServer) Send(m *ArtifactResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Healthz_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthzServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.healthz.Healthz/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthzServer).Check(ctx, req.(*CheckRequest))
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
		{
			MethodName: "List",
			Handler:    _Healthz_List_Handler,
		},
		{
			MethodName: "Acknowledge",
			Handler:    _Healthz_Acknowledge_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Healthz_Check_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Artifact",
			Handler:       _Healthz_Artifact_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "healthz/healthz.proto",
}
