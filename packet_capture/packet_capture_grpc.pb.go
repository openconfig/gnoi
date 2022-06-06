// Code generated by protoc-gen-go-grpc. DO NOT EDIT
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.3
// source: packet_capture/packet_capture.proto

package pcap

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

// PacketCaptureClient is the client API for PacketCapture service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PacketCaptureClient interface {
	// Pcap starts packet capture on the target and streams back
	// the results.
	Pcap(ctx context.Context, in *PcapRequest, opts ...grpc.CallOption) (PacketCapture_PcapClient, error)
}

type packetCaptureClient struct {
	cc grpc.ClientConnInterface
}

func NewPacketCaptureClient(cc grpc.ClientConnInterface) PacketCaptureClient {
	return &packetCaptureClient{cc}
}

func (c *packetCaptureClient) Pcap(ctx context.Context, in *PcapRequest, opts ...grpc.CallOption) (PacketCapture_PcapClient, error) {
	stream, err := c.cc.NewStream(ctx, &PacketCapture_ServiceDesc.Streams[0], "/gnoi.pcap.PacketCapture/Pcap", opts...)
	if err != nil {
		return nil, err
	}
	x := &packetCapturePcapClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PacketCapture_PcapClient interface {
	Recv() (*PcapResponse, error)
	grpc.ClientStream
}

type packetCapturePcapClient struct {
	grpc.ClientStream
}

func (x *packetCapturePcapClient) Recv() (*PcapResponse, error) {
	m := new(PcapResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PacketCaptureServer is the server API for PacketCapture service.
// All implementations must embed UnimplementedPacketCaptureServer
// for forward compatibility
type PacketCaptureServer interface {
	// Pcap starts packet capture on the target and streams back
	// the results.
	Pcap(*PcapRequest, PacketCapture_PcapServer) error
	mustEmbedUnimplementedPacketCaptureServer()
}

// UnimplementedPacketCaptureServer must be embedded to have forward compatible implementations.
type UnimplementedPacketCaptureServer struct {
}

func (UnimplementedPacketCaptureServer) Pcap(*PcapRequest, PacketCapture_PcapServer) error {
	return status.Errorf(codes.Unimplemented, "method Pcap not implemented")
}
func (UnimplementedPacketCaptureServer) mustEmbedUnimplementedPacketCaptureServer() {}

// UnsafePacketCaptureServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PacketCaptureServer will
// result in compilation errors.
type UnsafePacketCaptureServer interface {
	mustEmbedUnimplementedPacketCaptureServer()
}

func RegisterPacketCaptureServer(s grpc.ServiceRegistrar, srv PacketCaptureServer) {
	s.RegisterService(&PacketCapture_ServiceDesc, srv)
}

func _PacketCapture_Pcap_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PcapRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PacketCaptureServer).Pcap(m, &packetCapturePcapServer{stream})
}

type PacketCapture_PcapServer interface {
	Send(*PcapResponse) error
	grpc.ServerStream
}

type packetCapturePcapServer struct {
	grpc.ServerStream
}

func (x *packetCapturePcapServer) Send(m *PcapResponse) error {
	return x.ServerStream.SendMsg(m)
}

// PacketCapture_ServiceDesc is the grpc.ServiceDesc for PacketCapture service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PacketCapture_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.pcap.PacketCapture",
	HandlerType: (*PacketCaptureServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pcap",
			Handler:       _PacketCapture_Pcap_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "packet_capture/packet_capture.proto",
}
