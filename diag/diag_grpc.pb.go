// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: diag/diag.proto

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

// DiagClient is the client API for Diag service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiagClient interface {
	// Starts BERT operation on a set of ports. Each BERT operation is uniquely
	// identified by an ID, which is given by the caller. The caller can then
	// use this ID (as well as the list of the ports) to stop the BERT operation
	// and/or get the BERT results. This RPC is expected to return an error status
	// in the following situations:
	// - When BERT operation is supported on none of the ports specified by
	//   the request.
	// - When BERT is already in progress on any port specified by the request.
	// - In case of any low-level HW/SW internal errors.
	// The RPC returns an OK status of none of these situations is encountered.
	StartBERT(ctx context.Context, in *StartBERTRequest, opts ...grpc.CallOption) (*StartBERTResponse, error)
	// Stops an already in-progress BERT operation on a set of ports. The caller
	// uses the BERT operation ID it previously used when starting the operation
	// to stop it. The RPC is expected to return an error status in the following
	// situations:
	// - When there is at least one BERT operation in progress on a port which
	//   cannot be stopped in the middle of the operation (either due to lack of
	//   support or internal problems).
	// - When no BERT operation, which matches the given BERT operation ID, is in
	//   progress or completed on any of the ports specified by the request.
	// - When the BERT operation ID does not match the in progress or completed
	//   BERT operation on any of the ports specified by the request.
	// The RPC returns an OK status of none of these situations is encountered.
	// Note that a BERT operation is considered completed if the device has a
	// record/history of it. Also note that it is OK to receive a stop request for
	// a port which has completed BERT, as long as the recorded BERT operation ID
	// matches the one specified by the request.
	StopBERT(ctx context.Context, in *StopBERTRequest, opts ...grpc.CallOption) (*StopBERTResponse, error)
	// Gets BERT results during the BERT operation or after it completes. The
	// caller uses the BERT operation ID it previously used when starting the
	// operation to query it. The device is expected to keep the results for
	// last N BERT operations for some amount of time, as specified by the
	// product requirement. This RPC is expected to return error status in the
	// following situations:
	// - When no BERT operation, which matches the given BERT operation ID, is in
	//   progress or completed on any of the ports specified by the request.
	// - When the BERT operation ID does not match the in progress or completed
	//   BERT operation on any of the ports specified by the request.
	// The RPC returns an OK status of none of these situations is encountered.
	// Note that a BERT operation is considered completed if device has a
	// record of it.
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
	err := c.cc.Invoke(ctx, "/gnoi.diag.Diag/StartBERT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diagClient) StopBERT(ctx context.Context, in *StopBERTRequest, opts ...grpc.CallOption) (*StopBERTResponse, error) {
	out := new(StopBERTResponse)
	err := c.cc.Invoke(ctx, "/gnoi.diag.Diag/StopBERT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diagClient) GetBERTResult(ctx context.Context, in *GetBERTResultRequest, opts ...grpc.CallOption) (*GetBERTResultResponse, error) {
	out := new(GetBERTResultResponse)
	err := c.cc.Invoke(ctx, "/gnoi.diag.Diag/GetBERTResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiagServer is the server API for Diag service.
// All implementations must embed UnimplementedDiagServer
// for forward compatibility
type DiagServer interface {
	// Starts BERT operation on a set of ports. Each BERT operation is uniquely
	// identified by an ID, which is given by the caller. The caller can then
	// use this ID (as well as the list of the ports) to stop the BERT operation
	// and/or get the BERT results. This RPC is expected to return an error status
	// in the following situations:
	// - When BERT operation is supported on none of the ports specified by
	//   the request.
	// - When BERT is already in progress on any port specified by the request.
	// - In case of any low-level HW/SW internal errors.
	// The RPC returns an OK status of none of these situations is encountered.
	StartBERT(context.Context, *StartBERTRequest) (*StartBERTResponse, error)
	// Stops an already in-progress BERT operation on a set of ports. The caller
	// uses the BERT operation ID it previously used when starting the operation
	// to stop it. The RPC is expected to return an error status in the following
	// situations:
	// - When there is at least one BERT operation in progress on a port which
	//   cannot be stopped in the middle of the operation (either due to lack of
	//   support or internal problems).
	// - When no BERT operation, which matches the given BERT operation ID, is in
	//   progress or completed on any of the ports specified by the request.
	// - When the BERT operation ID does not match the in progress or completed
	//   BERT operation on any of the ports specified by the request.
	// The RPC returns an OK status of none of these situations is encountered.
	// Note that a BERT operation is considered completed if the device has a
	// record/history of it. Also note that it is OK to receive a stop request for
	// a port which has completed BERT, as long as the recorded BERT operation ID
	// matches the one specified by the request.
	StopBERT(context.Context, *StopBERTRequest) (*StopBERTResponse, error)
	// Gets BERT results during the BERT operation or after it completes. The
	// caller uses the BERT operation ID it previously used when starting the
	// operation to query it. The device is expected to keep the results for
	// last N BERT operations for some amount of time, as specified by the
	// product requirement. This RPC is expected to return error status in the
	// following situations:
	// - When no BERT operation, which matches the given BERT operation ID, is in
	//   progress or completed on any of the ports specified by the request.
	// - When the BERT operation ID does not match the in progress or completed
	//   BERT operation on any of the ports specified by the request.
	// The RPC returns an OK status of none of these situations is encountered.
	// Note that a BERT operation is considered completed if device has a
	// record of it.
	GetBERTResult(context.Context, *GetBERTResultRequest) (*GetBERTResultResponse, error)
	mustEmbedUnimplementedDiagServer()
}

// UnimplementedDiagServer must be embedded to have forward compatible implementations.
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
func (UnimplementedDiagServer) mustEmbedUnimplementedDiagServer() {}

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
		FullMethod: "/gnoi.diag.Diag/StartBERT",
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
		FullMethod: "/gnoi.diag.Diag/StopBERT",
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
		FullMethod: "/gnoi.diag.Diag/GetBERTResult",
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
	Metadata: "diag/diag.proto",
}
