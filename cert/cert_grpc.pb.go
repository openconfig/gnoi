//
// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// This file defines the gNOI API to be used for certificate installation and
// rotation.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: cert/cert.proto

package cert

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
	CertificateManagement_Rotate_FullMethodName                         = "/gnoi.certificate.CertificateManagement/Rotate"
	CertificateManagement_Install_FullMethodName                        = "/gnoi.certificate.CertificateManagement/Install"
	CertificateManagement_GenerateCSR_FullMethodName                    = "/gnoi.certificate.CertificateManagement/GenerateCSR"
	CertificateManagement_LoadCertificate_FullMethodName                = "/gnoi.certificate.CertificateManagement/LoadCertificate"
	CertificateManagement_LoadCertificateAuthorityBundle_FullMethodName = "/gnoi.certificate.CertificateManagement/LoadCertificateAuthorityBundle"
	CertificateManagement_GetCertificates_FullMethodName                = "/gnoi.certificate.CertificateManagement/GetCertificates"
	CertificateManagement_RevokeCertificates_FullMethodName             = "/gnoi.certificate.CertificateManagement/RevokeCertificates"
	CertificateManagement_CanGenerateCSR_FullMethodName                 = "/gnoi.certificate.CertificateManagement/CanGenerateCSR"
)

// CertificateManagementClient is the client API for CertificateManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CertificateManagementClient interface {
	// Rotate will replace an existing Certificate on the target by creating a
	// new CSR request and placing the new Certificate based on the CSR on the
	// target. If the stream is broken or any steps in the process fail the
	// target must rollback to the original Certificate.
	//
	// The following describes the sequence of messages that must be exchanged
	// in the Rotate() RPC.
	//
	// Sequence of expected messages:
	// Case 1: When Target generates the CSR.
	//
	//	Step 1: Start the stream
	//	  Client <---- Rotate() RPC stream begin ------> Target
	//
	//	Step 2: CSR
	//	  Client -----> GenerateCSRRequest----> Target
	//	  Client <----- GenerateCSRResponse <--- Target
	//
	//	Step 3: Certificate Signing
	//	  Client gets the certificate signed by the CA.
	//
	//	Step 4: Send Certificate to Target.
	//	  Client --> LoadCertificateRequest ----> Target
	//	  Client <-- LoadCertificateResponse <--- Target
	//
	//	Step 5: Test/Validation by the client.
	//	  This step should be to create a new connection to the target using
	//	  The new certificate and validate that the certificate works.
	//	  Once verfied, the client will then proceed to finalize the rotation.
	//	  If the new connection cannot be completed the client will cancel the
	//	  RPC thereby forcing the target to rollback the certificate.
	//
	//	Step 6: Final commit.
	//	  Client ---> FinalizeRequest ----> Target
	//
	// Case 2: When Client generates the CSR.
	//
	//	Step 1: Start the stream
	//	  Client <---- Rotate() RPC stream begin ----> Target
	//
	//	Step 2: CSR
	//	  Client generates its own certificate.
	//
	//	Step 3: Certificate Signing
	//	  Client gets the certificate signed by the CA.
	//
	//	Step 4: Send Certificate to Target.
	//	  Client ---> LoadCertificateRequest ----> Target
	//	  Client <--- LoadCertificateResponse <--- Target
	//
	//	Step 5: Test/Validation by the client.
	//
	//	Step 6: Final commit.
	//	  Client ---> FinalizeRequest ----> Target
	Rotate(ctx context.Context, opts ...grpc.CallOption) (CertificateManagement_RotateClient, error)
	// Install will put a new Certificate on the target by creating a new CSR
	// request and placing the new Certificate based on the CSR on the target.The
	// new Certificate will be associated with a new Certificate Id on the target.
	// If the target has a pre existing Certificate with the given Certificate Id,
	// the operation should fail.
	// If the stream is broken or any steps in the process fail the target must
	// revert any changes in state.
	//
	// The following describes the sequence of messages that must be exchanged
	// in the Install() RPC.
	//
	// Sequence of expected messages:
	// Case 1: When Target generates the CSR-------------------------:
	//
	//	Step 1: Start the stream
	//	  Client <---- Install() RPC stream begin ------> Target
	//
	//	Step 2: CSR
	//	  Client -----> GenerateCSRRequest() ----> Target
	//	  Client <---- GenerateCSRResponse() <---- Target
	//
	//	Step 3: Certificate Signing
	//	  Client gets the certificate signed by the CA.
	//
	//	Step 4: Send Certificate to Target.
	//	  Client -> LoadCertificateRequest() ----> Target
	//	  Client <- LoadCertificateResponse() <--- Target
	//
	// Case 2: When Client generates the CSR-------------------------:
	//
	//	Step 1: Start the stream
	//	  Client <---- Install() RPC stream begin ------> Target
	//
	//	Step 2: CSR
	//	  Client generates its own certificate.
	//
	//	Step 3: Certificate Signing
	//	  Client gets the certificate signed by the CA.
	//
	//	Step 4: Send Certificate to Target.
	//	  Client -> LoadCertificateRequest() ----> Target
	//	  Client <- LoadCertificateResponse() <--- Target
	Install(ctx context.Context, opts ...grpc.CallOption) (CertificateManagement_InstallClient, error)
	// When credentials are generated on the device, generates a keypair and
	// returns the Certificate Signing Request (CSR). The CSR has the public key,
	// which when signed by the CA, becomes the Certificate.
	GenerateCSR(ctx context.Context, in *GenerateCSRRequest, opts ...grpc.CallOption) (*GenerateCSRResponse, error)
	// Loads a certificate signed by a Certificate Authority (CA).
	LoadCertificate(ctx context.Context, in *LoadCertificateRequest, opts ...grpc.CallOption) (*LoadCertificateResponse, error)
	// Loads a bundle of CA certificates.
	LoadCertificateAuthorityBundle(ctx context.Context, in *LoadCertificateAuthorityBundleRequest, opts ...grpc.CallOption) (*LoadCertificateAuthorityBundleResponse, error)
	// An RPC to get the certificates on the target.
	GetCertificates(ctx context.Context, in *GetCertificatesRequest, opts ...grpc.CallOption) (*GetCertificatesResponse, error)
	// An RPC to revoke specific certificates.
	// If a certificate is not present on the target, the request should silently
	// succeed. Revoking a certificate should render the existing certificate
	// unusable by any endpoints.
	RevokeCertificates(ctx context.Context, in *RevokeCertificatesRequest, opts ...grpc.CallOption) (*RevokeCertificatesResponse, error)
	// An RPC to ask a target if it can generate a Certificate.
	CanGenerateCSR(ctx context.Context, in *CanGenerateCSRRequest, opts ...grpc.CallOption) (*CanGenerateCSRResponse, error)
}

type certificateManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateManagementClient(cc grpc.ClientConnInterface) CertificateManagementClient {
	return &certificateManagementClient{cc}
}

func (c *certificateManagementClient) Rotate(ctx context.Context, opts ...grpc.CallOption) (CertificateManagement_RotateClient, error) {
	stream, err := c.cc.NewStream(ctx, &CertificateManagement_ServiceDesc.Streams[0], CertificateManagement_Rotate_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &certificateManagementRotateClient{stream}
	return x, nil
}

type CertificateManagement_RotateClient interface {
	Send(*RotateCertificateRequest) error
	Recv() (*RotateCertificateResponse, error)
	grpc.ClientStream
}

type certificateManagementRotateClient struct {
	grpc.ClientStream
}

func (x *certificateManagementRotateClient) Send(m *RotateCertificateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *certificateManagementRotateClient) Recv() (*RotateCertificateResponse, error) {
	m := new(RotateCertificateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *certificateManagementClient) Install(ctx context.Context, opts ...grpc.CallOption) (CertificateManagement_InstallClient, error) {
	stream, err := c.cc.NewStream(ctx, &CertificateManagement_ServiceDesc.Streams[1], CertificateManagement_Install_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &certificateManagementInstallClient{stream}
	return x, nil
}

type CertificateManagement_InstallClient interface {
	Send(*InstallCertificateRequest) error
	Recv() (*InstallCertificateResponse, error)
	grpc.ClientStream
}

type certificateManagementInstallClient struct {
	grpc.ClientStream
}

func (x *certificateManagementInstallClient) Send(m *InstallCertificateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *certificateManagementInstallClient) Recv() (*InstallCertificateResponse, error) {
	m := new(InstallCertificateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *certificateManagementClient) GenerateCSR(ctx context.Context, in *GenerateCSRRequest, opts ...grpc.CallOption) (*GenerateCSRResponse, error) {
	out := new(GenerateCSRResponse)
	err := c.cc.Invoke(ctx, CertificateManagement_GenerateCSR_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateManagementClient) LoadCertificate(ctx context.Context, in *LoadCertificateRequest, opts ...grpc.CallOption) (*LoadCertificateResponse, error) {
	out := new(LoadCertificateResponse)
	err := c.cc.Invoke(ctx, CertificateManagement_LoadCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateManagementClient) LoadCertificateAuthorityBundle(ctx context.Context, in *LoadCertificateAuthorityBundleRequest, opts ...grpc.CallOption) (*LoadCertificateAuthorityBundleResponse, error) {
	out := new(LoadCertificateAuthorityBundleResponse)
	err := c.cc.Invoke(ctx, CertificateManagement_LoadCertificateAuthorityBundle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateManagementClient) GetCertificates(ctx context.Context, in *GetCertificatesRequest, opts ...grpc.CallOption) (*GetCertificatesResponse, error) {
	out := new(GetCertificatesResponse)
	err := c.cc.Invoke(ctx, CertificateManagement_GetCertificates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateManagementClient) RevokeCertificates(ctx context.Context, in *RevokeCertificatesRequest, opts ...grpc.CallOption) (*RevokeCertificatesResponse, error) {
	out := new(RevokeCertificatesResponse)
	err := c.cc.Invoke(ctx, CertificateManagement_RevokeCertificates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateManagementClient) CanGenerateCSR(ctx context.Context, in *CanGenerateCSRRequest, opts ...grpc.CallOption) (*CanGenerateCSRResponse, error) {
	out := new(CanGenerateCSRResponse)
	err := c.cc.Invoke(ctx, CertificateManagement_CanGenerateCSR_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateManagementServer is the server API for CertificateManagement service.
// All implementations must embed UnimplementedCertificateManagementServer
// for forward compatibility
type CertificateManagementServer interface {
	// Rotate will replace an existing Certificate on the target by creating a
	// new CSR request and placing the new Certificate based on the CSR on the
	// target. If the stream is broken or any steps in the process fail the
	// target must rollback to the original Certificate.
	//
	// The following describes the sequence of messages that must be exchanged
	// in the Rotate() RPC.
	//
	// Sequence of expected messages:
	// Case 1: When Target generates the CSR.
	//
	//	Step 1: Start the stream
	//	  Client <---- Rotate() RPC stream begin ------> Target
	//
	//	Step 2: CSR
	//	  Client -----> GenerateCSRRequest----> Target
	//	  Client <----- GenerateCSRResponse <--- Target
	//
	//	Step 3: Certificate Signing
	//	  Client gets the certificate signed by the CA.
	//
	//	Step 4: Send Certificate to Target.
	//	  Client --> LoadCertificateRequest ----> Target
	//	  Client <-- LoadCertificateResponse <--- Target
	//
	//	Step 5: Test/Validation by the client.
	//	  This step should be to create a new connection to the target using
	//	  The new certificate and validate that the certificate works.
	//	  Once verfied, the client will then proceed to finalize the rotation.
	//	  If the new connection cannot be completed the client will cancel the
	//	  RPC thereby forcing the target to rollback the certificate.
	//
	//	Step 6: Final commit.
	//	  Client ---> FinalizeRequest ----> Target
	//
	// Case 2: When Client generates the CSR.
	//
	//	Step 1: Start the stream
	//	  Client <---- Rotate() RPC stream begin ----> Target
	//
	//	Step 2: CSR
	//	  Client generates its own certificate.
	//
	//	Step 3: Certificate Signing
	//	  Client gets the certificate signed by the CA.
	//
	//	Step 4: Send Certificate to Target.
	//	  Client ---> LoadCertificateRequest ----> Target
	//	  Client <--- LoadCertificateResponse <--- Target
	//
	//	Step 5: Test/Validation by the client.
	//
	//	Step 6: Final commit.
	//	  Client ---> FinalizeRequest ----> Target
	Rotate(CertificateManagement_RotateServer) error
	// Install will put a new Certificate on the target by creating a new CSR
	// request and placing the new Certificate based on the CSR on the target.The
	// new Certificate will be associated with a new Certificate Id on the target.
	// If the target has a pre existing Certificate with the given Certificate Id,
	// the operation should fail.
	// If the stream is broken or any steps in the process fail the target must
	// revert any changes in state.
	//
	// The following describes the sequence of messages that must be exchanged
	// in the Install() RPC.
	//
	// Sequence of expected messages:
	// Case 1: When Target generates the CSR-------------------------:
	//
	//	Step 1: Start the stream
	//	  Client <---- Install() RPC stream begin ------> Target
	//
	//	Step 2: CSR
	//	  Client -----> GenerateCSRRequest() ----> Target
	//	  Client <---- GenerateCSRResponse() <---- Target
	//
	//	Step 3: Certificate Signing
	//	  Client gets the certificate signed by the CA.
	//
	//	Step 4: Send Certificate to Target.
	//	  Client -> LoadCertificateRequest() ----> Target
	//	  Client <- LoadCertificateResponse() <--- Target
	//
	// Case 2: When Client generates the CSR-------------------------:
	//
	//	Step 1: Start the stream
	//	  Client <---- Install() RPC stream begin ------> Target
	//
	//	Step 2: CSR
	//	  Client generates its own certificate.
	//
	//	Step 3: Certificate Signing
	//	  Client gets the certificate signed by the CA.
	//
	//	Step 4: Send Certificate to Target.
	//	  Client -> LoadCertificateRequest() ----> Target
	//	  Client <- LoadCertificateResponse() <--- Target
	Install(CertificateManagement_InstallServer) error
	// When credentials are generated on the device, generates a keypair and
	// returns the Certificate Signing Request (CSR). The CSR has the public key,
	// which when signed by the CA, becomes the Certificate.
	GenerateCSR(context.Context, *GenerateCSRRequest) (*GenerateCSRResponse, error)
	// Loads a certificate signed by a Certificate Authority (CA).
	LoadCertificate(context.Context, *LoadCertificateRequest) (*LoadCertificateResponse, error)
	// Loads a bundle of CA certificates.
	LoadCertificateAuthorityBundle(context.Context, *LoadCertificateAuthorityBundleRequest) (*LoadCertificateAuthorityBundleResponse, error)
	// An RPC to get the certificates on the target.
	GetCertificates(context.Context, *GetCertificatesRequest) (*GetCertificatesResponse, error)
	// An RPC to revoke specific certificates.
	// If a certificate is not present on the target, the request should silently
	// succeed. Revoking a certificate should render the existing certificate
	// unusable by any endpoints.
	RevokeCertificates(context.Context, *RevokeCertificatesRequest) (*RevokeCertificatesResponse, error)
	// An RPC to ask a target if it can generate a Certificate.
	CanGenerateCSR(context.Context, *CanGenerateCSRRequest) (*CanGenerateCSRResponse, error)
	mustEmbedUnimplementedCertificateManagementServer()
}

// UnimplementedCertificateManagementServer must be embedded to have forward compatible implementations.
type UnimplementedCertificateManagementServer struct {
}

func (UnimplementedCertificateManagementServer) Rotate(CertificateManagement_RotateServer) error {
	return status.Errorf(codes.Unimplemented, "method Rotate not implemented")
}
func (UnimplementedCertificateManagementServer) Install(CertificateManagement_InstallServer) error {
	return status.Errorf(codes.Unimplemented, "method Install not implemented")
}
func (UnimplementedCertificateManagementServer) GenerateCSR(context.Context, *GenerateCSRRequest) (*GenerateCSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateCSR not implemented")
}
func (UnimplementedCertificateManagementServer) LoadCertificate(context.Context, *LoadCertificateRequest) (*LoadCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadCertificate not implemented")
}
func (UnimplementedCertificateManagementServer) LoadCertificateAuthorityBundle(context.Context, *LoadCertificateAuthorityBundleRequest) (*LoadCertificateAuthorityBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadCertificateAuthorityBundle not implemented")
}
func (UnimplementedCertificateManagementServer) GetCertificates(context.Context, *GetCertificatesRequest) (*GetCertificatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificates not implemented")
}
func (UnimplementedCertificateManagementServer) RevokeCertificates(context.Context, *RevokeCertificatesRequest) (*RevokeCertificatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeCertificates not implemented")
}
func (UnimplementedCertificateManagementServer) CanGenerateCSR(context.Context, *CanGenerateCSRRequest) (*CanGenerateCSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanGenerateCSR not implemented")
}
func (UnimplementedCertificateManagementServer) mustEmbedUnimplementedCertificateManagementServer() {}

// UnsafeCertificateManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateManagementServer will
// result in compilation errors.
type UnsafeCertificateManagementServer interface {
	mustEmbedUnimplementedCertificateManagementServer()
}

func RegisterCertificateManagementServer(s grpc.ServiceRegistrar, srv CertificateManagementServer) {
	s.RegisterService(&CertificateManagement_ServiceDesc, srv)
}

func _CertificateManagement_Rotate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CertificateManagementServer).Rotate(&certificateManagementRotateServer{stream})
}

type CertificateManagement_RotateServer interface {
	Send(*RotateCertificateResponse) error
	Recv() (*RotateCertificateRequest, error)
	grpc.ServerStream
}

type certificateManagementRotateServer struct {
	grpc.ServerStream
}

func (x *certificateManagementRotateServer) Send(m *RotateCertificateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *certificateManagementRotateServer) Recv() (*RotateCertificateRequest, error) {
	m := new(RotateCertificateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CertificateManagement_Install_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CertificateManagementServer).Install(&certificateManagementInstallServer{stream})
}

type CertificateManagement_InstallServer interface {
	Send(*InstallCertificateResponse) error
	Recv() (*InstallCertificateRequest, error)
	grpc.ServerStream
}

type certificateManagementInstallServer struct {
	grpc.ServerStream
}

func (x *certificateManagementInstallServer) Send(m *InstallCertificateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *certificateManagementInstallServer) Recv() (*InstallCertificateRequest, error) {
	m := new(InstallCertificateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CertificateManagement_GenerateCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateCSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateManagementServer).GenerateCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateManagement_GenerateCSR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateManagementServer).GenerateCSR(ctx, req.(*GenerateCSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateManagement_LoadCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateManagementServer).LoadCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateManagement_LoadCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateManagementServer).LoadCertificate(ctx, req.(*LoadCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateManagement_LoadCertificateAuthorityBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadCertificateAuthorityBundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateManagementServer).LoadCertificateAuthorityBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateManagement_LoadCertificateAuthorityBundle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateManagementServer).LoadCertificateAuthorityBundle(ctx, req.(*LoadCertificateAuthorityBundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateManagement_GetCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateManagementServer).GetCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateManagement_GetCertificates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateManagementServer).GetCertificates(ctx, req.(*GetCertificatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateManagement_RevokeCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeCertificatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateManagementServer).RevokeCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateManagement_RevokeCertificates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateManagementServer).RevokeCertificates(ctx, req.(*RevokeCertificatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateManagement_CanGenerateCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CanGenerateCSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateManagementServer).CanGenerateCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateManagement_CanGenerateCSR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateManagementServer).CanGenerateCSR(ctx, req.(*CanGenerateCSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificateManagement_ServiceDesc is the grpc.ServiceDesc for CertificateManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.certificate.CertificateManagement",
	HandlerType: (*CertificateManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateCSR",
			Handler:    _CertificateManagement_GenerateCSR_Handler,
		},
		{
			MethodName: "LoadCertificate",
			Handler:    _CertificateManagement_LoadCertificate_Handler,
		},
		{
			MethodName: "LoadCertificateAuthorityBundle",
			Handler:    _CertificateManagement_LoadCertificateAuthorityBundle_Handler,
		},
		{
			MethodName: "GetCertificates",
			Handler:    _CertificateManagement_GetCertificates_Handler,
		},
		{
			MethodName: "RevokeCertificates",
			Handler:    _CertificateManagement_RevokeCertificates_Handler,
		},
		{
			MethodName: "CanGenerateCSR",
			Handler:    _CertificateManagement_CanGenerateCSR_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Rotate",
			Handler:       _CertificateManagement_Rotate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Install",
			Handler:       _CertificateManagement_Install_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cert/cert.proto",
}
