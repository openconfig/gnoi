// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.1
// source: github.com/openconfig/gnoi/diag/diag.proto

package diag

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
	Diag_StartBERT_FullMethodName     = "/gnoi.diag.Diag/StartBERT"
	Diag_StopBERT_FullMethodName      = "/gnoi.diag.Diag/StopBERT"
	Diag_GetBERTResult_FullMethodName = "/gnoi.diag.Diag/GetBERTResult"
)

// DiagClient is the client API for Diag service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiagClient interface {
	StartBERT(ctx context.Context, in *StartBERTRequest, opts ...grpc.CallOption) (*StartBERTResponse, error)
	StopBERT(ctx context.Context, in *StopBERTRequest, opts ...grpc.CallOption) (*StopBERTResponse, error)
	GetBERTResult(ctx context.Context, in *GetBERTResultRequest, opts ...grpc.CallOption) (*GetBERTResultResponse, error)
}

type diagClient struct {
	cc grpc.ClientConnInterface
}

func NewDiagClient(cc grpc.ClientConnInterface) DiagClient {
	return &diagClient{cc}
}

func (c *diagClient) StartBERT(ctx context.Context, in *StartBERTRequest, opts ...grpc.CallOption) (*StartBERTResponse, error) {
	out := new(StartBERTResponse)
	err := c.cc.Invoke(ctx, Diag_StartBERT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diagClient) StopBERT(ctx context.Context, in *StopBERTRequest, opts ...grpc.CallOption) (*StopBERTResponse, error) {
	out := new(StopBERTResponse)
	err := c.cc.Invoke(ctx, Diag_StopBERT_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diagClient) GetBERTResult(ctx context.Context, in *GetBERTResultRequest, opts ...grpc.CallOption) (*GetBERTResultResponse, error) {
	out := new(GetBERTResultResponse)
	err := c.cc.Invoke(ctx, Diag_GetBERTResult_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiagServer is the server API for Diag service.
// All implementations should embed UnimplementedDiagServer
// for forward compatibility
type DiagServer interface {
	StartBERT(context.Context, *StartBERTRequest) (*StartBERTResponse, error)
	StopBERT(context.Context, *StopBERTRequest) (*StopBERTResponse, error)
	GetBERTResult(context.Context, *GetBERTResultRequest) (*GetBERTResultResponse, error)
}

// UnimplementedDiagServer should be embedded to have forward compatible implementations.
type UnimplementedDiagServer struct {
}

func (UnimplementedDiagServer) StartBERT(context.Context, *StartBERTRequest) (*StartBERTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartBERT not implemented")
}
func (UnimplementedDiagServer) StopBERT(context.Context, *StopBERTRequest) (*StopBERTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopBERT not implemented")
}
func (UnimplementedDiagServer) GetBERTResult(context.Context, *GetBERTResultRequest) (*GetBERTResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBERTResult not implemented")
}

// UnsafeDiagServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiagServer will
// result in compilation errors.
type UnsafeDiagServer interface {
	mustEmbedUnimplementedDiagServer()
}

func RegisterDiagServer(s grpc.ServiceRegistrar, srv DiagServer) {
	s.RegisterService(&Diag_ServiceDesc, srv)
}

func _Diag_StartBERT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartBERTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiagServer).StartBERT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Diag_StartBERT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiagServer).StartBERT(ctx, req.(*StartBERTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Diag_StopBERT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopBERTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiagServer).StopBERT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Diag_StopBERT_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiagServer).StopBERT(ctx, req.(*StopBERTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Diag_GetBERTResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBERTResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiagServer).GetBERTResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Diag_GetBERTResult_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiagServer).GetBERTResult(ctx, req.(*GetBERTResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Diag_ServiceDesc is the grpc.ServiceDesc for Diag service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Diag_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.diag.Diag",
	HandlerType: (*DiagServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartBERT",
			Handler:    _Diag_StartBERT_Handler,
		},
		{
			MethodName: "StopBERT",
			Handler:    _Diag_StopBERT_Handler,
		},
		{
			MethodName: "GetBERTResult",
			Handler:    _Diag_GetBERTResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openconfig/gnoi/diag/diag.proto",
}
