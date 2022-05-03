// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: otdr/otdr.proto

package otdr

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

// OTDRClient is the client API for OTDR service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OTDRClient interface {
	// Initiate triggers an optical time domain reflectometer (OTDR) trace to
	// be run on a specified port. The results of the trace may be sent back in
	// the response or saved on the device to be retrieved later. If the
	// information is saved on the device a path to the file must be returned
	// which can later be used by File.Get or File.TransferToRemote. The RPC
	// should initiate the OTDR trace and wait until the OTDR trace has completed
	// and the device has processed the results before returning. If the RPC is
	// cancelled while in operation, the running OTDR trace should stop.
	Initiate(ctx context.Context, in *InitiateRequest, opts ...grpc.CallOption) (OTDR_InitiateClient, error)
}

type oTDRClient struct {
	cc grpc.ClientConnInterface
}

func NewOTDRClient(cc grpc.ClientConnInterface) OTDRClient {
	return &oTDRClient{cc}
}

func (c *oTDRClient) Initiate(ctx context.Context, in *InitiateRequest, opts ...grpc.CallOption) (OTDR_InitiateClient, error) {
	stream, err := c.cc.NewStream(ctx, &OTDR_ServiceDesc.Streams[0], "/gnoi.optical.OTDR/Initiate", opts...)
	if err != nil {
		return nil, err
	}
	x := &oTDRInitiateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OTDR_InitiateClient interface {
	Recv() (*InitiateResponse, error)
	grpc.ClientStream
}

type oTDRInitiateClient struct {
	grpc.ClientStream
}

func (x *oTDRInitiateClient) Recv() (*InitiateResponse, error) {
	m := new(InitiateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OTDRServer is the server API for OTDR service.
// All implementations must embed UnimplementedOTDRServer
// for forward compatibility
type OTDRServer interface {
	// Initiate triggers an optical time domain reflectometer (OTDR) trace to
	// be run on a specified port. The results of the trace may be sent back in
	// the response or saved on the device to be retrieved later. If the
	// information is saved on the device a path to the file must be returned
	// which can later be used by File.Get or File.TransferToRemote. The RPC
	// should initiate the OTDR trace and wait until the OTDR trace has completed
	// and the device has processed the results before returning. If the RPC is
	// cancelled while in operation, the running OTDR trace should stop.
	Initiate(*InitiateRequest, OTDR_InitiateServer) error
	mustEmbedUnimplementedOTDRServer()
}

// UnimplementedOTDRServer must be embedded to have forward compatible implementations.
type UnimplementedOTDRServer struct {
}

func (UnimplementedOTDRServer) Initiate(*InitiateRequest, OTDR_InitiateServer) error {
	return status.Errorf(codes.Unimplemented, "method Initiate not implemented")
}
func (UnimplementedOTDRServer) mustEmbedUnimplementedOTDRServer() {}

// UnsafeOTDRServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OTDRServer will
// result in compilation errors.
type UnsafeOTDRServer interface {
	mustEmbedUnimplementedOTDRServer()
}

func RegisterOTDRServer(s grpc.ServiceRegistrar, srv OTDRServer) {
	s.RegisterService(&OTDR_ServiceDesc, srv)
}

func _OTDR_Initiate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InitiateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OTDRServer).Initiate(m, &oTDRInitiateServer{stream})
}

type OTDR_InitiateServer interface {
	Send(*InitiateResponse) error
	grpc.ServerStream
}

type oTDRInitiateServer struct {
	grpc.ServerStream
}

func (x *oTDRInitiateServer) Send(m *InitiateResponse) error {
	return x.ServerStream.SendMsg(m)
}

// OTDR_ServiceDesc is the grpc.ServiceDesc for OTDR service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OTDR_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.optical.OTDR",
	HandlerType: (*OTDRServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Initiate",
			Handler:       _OTDR_Initiate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "otdr/otdr.proto",
}
