// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: packet_link_qualification/packet_link_qualification.proto

package packet_link_qualification

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

// PacketLinkQualClient is the client API for PacketLinkQual service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PacketLinkQualClient interface {
	// Starts packet based link qualification on a set of interfaces. Each start
	// operation is uniquely identified by an ID supplied by the caller. This ID
	// must be used by other operations (e.g., stop, get results, etc.) to select
	// the test associated with the ID. Caller specifies the minimum wait time
	// before warming up the link, number of packets, and preparation time to warm
	// up the link, and the actual duration of the qualification. It also
	// specifies if the interface under test is used as a near-end/ far-end from
	// the test perspective. On receiving the request, the device performs the
	// setup, and qualifies the link for the supplied test duration. The device
	// keeps the results until the corresponding stop request. Each start response
	// matches with a corresponding unique operation ID.
	//
	// If there is any error starting the qualification, the RPC is returned with
	// a non-zero code in the corresponding status, and a message to denote the
	// details.
	StartPacketQualification(ctx context.Context, in *StartPacketQualificationRequest, opts ...grpc.CallOption) (*StartPacketQualificationResponse, error)
	// Stops packet based link qualification on a set of interfaces. The caller
	// uses the operation ID it previously used when starting the operation to
	// stop it. Note that a qualification operation is considered completed if the
	// device has a record/history of it. Also note that it is OK to receive a
	// stop request for an interface which has completed qualification, as long as
	// the recorded operation ID matches the one specified by the request. Devices
	// should clean up at stop request. Each stop response matches with a
	// corresponding unique operation ID supplied in the request.
	//
	// If there is any error starting the qualification, the RPC is returned with
	// a non-zero code in the corresponding status, and a message to denote the
	// details.
	StopPacketQualification(ctx context.Context, in *StopPacketQualificationRequest, opts ...grpc.CallOption) (*StopPacketQualificationResponse, error)
	// Gets packet based link qualification results for a set of operation IDs.
	// Note that the results will be incomplete during the operation. The caller
	// uses the operation ID it previously used when starting the operation to
	// query it. The device is expected to keep the results until the
	// corresponding stop request. Each get response matches with a corresponding
	// unique operation ID supplied in the request.
	//
	// If there is any error starting the qualification, the RPC is returned with
	// a non-zero code in the corresponding status, and a message to denote the
	// details.
	GetPacketQualificationResult(ctx context.Context, in *GetPacketQualificationResultRequest, opts ...grpc.CallOption) (*GetPacketQualificationResultResponse, error)
}

type packetLinkQualClient struct {
	cc grpc.ClientConnInterface
}

func NewPacketLinkQualClient(cc grpc.ClientConnInterface) PacketLinkQualClient {
	return &packetLinkQualClient{cc}
}

func (c *packetLinkQualClient) StartPacketQualification(ctx context.Context, in *StartPacketQualificationRequest, opts ...grpc.CallOption) (*StartPacketQualificationResponse, error) {
	out := new(StartPacketQualificationResponse)
	err := c.cc.Invoke(ctx, "/gnoi.packet_link_qualification.PacketLinkQual/StartPacketQualification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packetLinkQualClient) StopPacketQualification(ctx context.Context, in *StopPacketQualificationRequest, opts ...grpc.CallOption) (*StopPacketQualificationResponse, error) {
	out := new(StopPacketQualificationResponse)
	err := c.cc.Invoke(ctx, "/gnoi.packet_link_qualification.PacketLinkQual/StopPacketQualification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *packetLinkQualClient) GetPacketQualificationResult(ctx context.Context, in *GetPacketQualificationResultRequest, opts ...grpc.CallOption) (*GetPacketQualificationResultResponse, error) {
	out := new(GetPacketQualificationResultResponse)
	err := c.cc.Invoke(ctx, "/gnoi.packet_link_qualification.PacketLinkQual/GetPacketQualificationResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PacketLinkQualServer is the server API for PacketLinkQual service.
// All implementations must embed UnimplementedPacketLinkQualServer
// for forward compatibility
type PacketLinkQualServer interface {
	// Starts packet based link qualification on a set of interfaces. Each start
	// operation is uniquely identified by an ID supplied by the caller. This ID
	// must be used by other operations (e.g., stop, get results, etc.) to select
	// the test associated with the ID. Caller specifies the minimum wait time
	// before warming up the link, number of packets, and preparation time to warm
	// up the link, and the actual duration of the qualification. It also
	// specifies if the interface under test is used as a near-end/ far-end from
	// the test perspective. On receiving the request, the device performs the
	// setup, and qualifies the link for the supplied test duration. The device
	// keeps the results until the corresponding stop request. Each start response
	// matches with a corresponding unique operation ID.
	//
	// If there is any error starting the qualification, the RPC is returned with
	// a non-zero code in the corresponding status, and a message to denote the
	// details.
	StartPacketQualification(context.Context, *StartPacketQualificationRequest) (*StartPacketQualificationResponse, error)
	// Stops packet based link qualification on a set of interfaces. The caller
	// uses the operation ID it previously used when starting the operation to
	// stop it. Note that a qualification operation is considered completed if the
	// device has a record/history of it. Also note that it is OK to receive a
	// stop request for an interface which has completed qualification, as long as
	// the recorded operation ID matches the one specified by the request. Devices
	// should clean up at stop request. Each stop response matches with a
	// corresponding unique operation ID supplied in the request.
	//
	// If there is any error starting the qualification, the RPC is returned with
	// a non-zero code in the corresponding status, and a message to denote the
	// details.
	StopPacketQualification(context.Context, *StopPacketQualificationRequest) (*StopPacketQualificationResponse, error)
	// Gets packet based link qualification results for a set of operation IDs.
	// Note that the results will be incomplete during the operation. The caller
	// uses the operation ID it previously used when starting the operation to
	// query it. The device is expected to keep the results until the
	// corresponding stop request. Each get response matches with a corresponding
	// unique operation ID supplied in the request.
	//
	// If there is any error starting the qualification, the RPC is returned with
	// a non-zero code in the corresponding status, and a message to denote the
	// details.
	GetPacketQualificationResult(context.Context, *GetPacketQualificationResultRequest) (*GetPacketQualificationResultResponse, error)
	mustEmbedUnimplementedPacketLinkQualServer()
}

// UnimplementedPacketLinkQualServer must be embedded to have forward compatible implementations.
type UnimplementedPacketLinkQualServer struct {
}

func (UnimplementedPacketLinkQualServer) StartPacketQualification(context.Context, *StartPacketQualificationRequest) (*StartPacketQualificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPacketQualification not implemented")
}
func (UnimplementedPacketLinkQualServer) StopPacketQualification(context.Context, *StopPacketQualificationRequest) (*StopPacketQualificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopPacketQualification not implemented")
}
func (UnimplementedPacketLinkQualServer) GetPacketQualificationResult(context.Context, *GetPacketQualificationResultRequest) (*GetPacketQualificationResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPacketQualificationResult not implemented")
}
func (UnimplementedPacketLinkQualServer) mustEmbedUnimplementedPacketLinkQualServer() {}

// UnsafePacketLinkQualServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PacketLinkQualServer will
// result in compilation errors.
type UnsafePacketLinkQualServer interface {
	mustEmbedUnimplementedPacketLinkQualServer()
}

func RegisterPacketLinkQualServer(s grpc.ServiceRegistrar, srv PacketLinkQualServer) {
	s.RegisterService(&PacketLinkQual_ServiceDesc, srv)
}

func _PacketLinkQual_StartPacketQualification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPacketQualificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PacketLinkQualServer).StartPacketQualification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.packet_link_qualification.PacketLinkQual/StartPacketQualification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PacketLinkQualServer).StartPacketQualification(ctx, req.(*StartPacketQualificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PacketLinkQual_StopPacketQualification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopPacketQualificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PacketLinkQualServer).StopPacketQualification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.packet_link_qualification.PacketLinkQual/StopPacketQualification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PacketLinkQualServer).StopPacketQualification(ctx, req.(*StopPacketQualificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PacketLinkQual_GetPacketQualificationResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPacketQualificationResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PacketLinkQualServer).GetPacketQualificationResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.packet_link_qualification.PacketLinkQual/GetPacketQualificationResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PacketLinkQualServer).GetPacketQualificationResult(ctx, req.(*GetPacketQualificationResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PacketLinkQual_ServiceDesc is the grpc.ServiceDesc for PacketLinkQual service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PacketLinkQual_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.packet_link_qualification.PacketLinkQual",
	HandlerType: (*PacketLinkQualServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartPacketQualification",
			Handler:    _PacketLinkQual_StartPacketQualification_Handler,
		},
		{
			MethodName: "StopPacketQualification",
			Handler:    _PacketLinkQual_StopPacketQualification_Handler,
		},
		{
			MethodName: "GetPacketQualificationResult",
			Handler:    _PacketLinkQual_GetPacketQualificationResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "packet_link_qualification/packet_link_qualification.proto",
}