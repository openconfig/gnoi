// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: dry_run/dry_run.proto

package dry_run

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

// DryRunClient is the client API for DryRun service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DryRunClient interface {
	Compare(ctx context.Context, in *CompareRequest, opts ...grpc.CallOption) (DryRun_CompareClient, error)
}

type dryRunClient struct {
	cc grpc.ClientConnInterface
}

func NewDryRunClient(cc grpc.ClientConnInterface) DryRunClient {
	return &dryRunClient{cc}
}

func (c *dryRunClient) Compare(ctx context.Context, in *CompareRequest, opts ...grpc.CallOption) (DryRun_CompareClient, error) {
	stream, err := c.cc.NewStream(ctx, &DryRun_ServiceDesc.Streams[0], "/gnoi.dry_run.DryRun/Compare", opts...)
	if err != nil {
		return nil, err
	}
	x := &dryRunCompareClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DryRun_CompareClient interface {
	Recv() (*CompareResponse, error)
	grpc.ClientStream
}

type dryRunCompareClient struct {
	grpc.ClientStream
}

func (x *dryRunCompareClient) Recv() (*CompareResponse, error) {
	m := new(CompareResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DryRunServer is the server API for DryRun service.
// All implementations must embed UnimplementedDryRunServer
// for forward compatibility
type DryRunServer interface {
	Compare(*CompareRequest, DryRun_CompareServer) error
	mustEmbedUnimplementedDryRunServer()
}

// UnimplementedDryRunServer must be embedded to have forward compatible implementations.
type UnimplementedDryRunServer struct {
}

func (UnimplementedDryRunServer) Compare(*CompareRequest, DryRun_CompareServer) error {
	return status.Errorf(codes.Unimplemented, "method Compare not implemented")
}
func (UnimplementedDryRunServer) mustEmbedUnimplementedDryRunServer() {}

// UnsafeDryRunServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DryRunServer will
// result in compilation errors.
type UnsafeDryRunServer interface {
	mustEmbedUnimplementedDryRunServer()
}

func RegisterDryRunServer(s grpc.ServiceRegistrar, srv DryRunServer) {
	s.RegisterService(&DryRun_ServiceDesc, srv)
}

func _DryRun_Compare_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CompareRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DryRunServer).Compare(m, &dryRunCompareServer{stream})
}

type DryRun_CompareServer interface {
	Send(*CompareResponse) error
	grpc.ServerStream
}

type dryRunCompareServer struct {
	grpc.ServerStream
}

func (x *dryRunCompareServer) Send(m *CompareResponse) error {
	return x.ServerStream.SendMsg(m)
}

// DryRun_ServiceDesc is the grpc.ServiceDesc for DryRun service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DryRun_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.dry_run.DryRun",
	HandlerType: (*DryRunServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Compare",
			Handler:       _DryRun_Compare_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dry_run/dry_run.proto",
}
