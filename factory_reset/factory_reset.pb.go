// This file defines a gNOI API used for factory reseting a Target.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: factory_reset/factory_reset.proto

package factory_reset

import (
	_ "github.com/openconfig/gnoi/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Instructs the Target to rollback the OS to the same version as it shipped
	// from factory.
	FactoryOs bool `protobuf:"varint,1,opt,name=factory_os,json=factoryOs,proto3" json:"factory_os,omitempty"`
	// Instructs the Target to zero fill persistent storage state data.
	ZeroFill bool `protobuf:"varint,2,opt,name=zero_fill,json=zeroFill,proto3" json:"zero_fill,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_factory_reset_factory_reset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_factory_reset_factory_reset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_factory_reset_factory_reset_proto_rawDescGZIP(), []int{0}
}

func (x *StartRequest) GetFactoryOs() bool {
	if x != nil {
		return x.FactoryOs
	}
	return false
}

func (x *StartRequest) GetZeroFill() bool {
	if x != nil {
		return x.ZeroFill
	}
	return false
}

type ResetSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetSuccess) Reset() {
	*x = ResetSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_factory_reset_factory_reset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetSuccess) ProtoMessage() {}

func (x *ResetSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_factory_reset_factory_reset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetSuccess.ProtoReflect.Descriptor instead.
func (*ResetSuccess) Descriptor() ([]byte, []int) {
	return file_factory_reset_factory_reset_proto_rawDescGZIP(), []int{1}
}

// Message also used in gRPC status.details field.
type ResetError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Factory OS reset is not supported.
	FactoryOsUnsupported bool `protobuf:"varint,1,opt,name=factory_os_unsupported,json=factoryOsUnsupported,proto3" json:"factory_os_unsupported,omitempty"`
	// Zero fill is not supported.
	ZeroFillUnsupported bool `protobuf:"varint,2,opt,name=zero_fill_unsupported,json=zeroFillUnsupported,proto3" json:"zero_fill_unsupported,omitempty"`
	// Unspecified error, must provide detail message.
	Other  bool   `protobuf:"varint,3,opt,name=other,proto3" json:"other,omitempty"`
	Detail string `protobuf:"bytes,4,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *ResetError) Reset() {
	*x = ResetError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_factory_reset_factory_reset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetError) ProtoMessage() {}

func (x *ResetError) ProtoReflect() protoreflect.Message {
	mi := &file_factory_reset_factory_reset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetError.ProtoReflect.Descriptor instead.
func (*ResetError) Descriptor() ([]byte, []int) {
	return file_factory_reset_factory_reset_proto_rawDescGZIP(), []int{2}
}

func (x *ResetError) GetFactoryOsUnsupported() bool {
	if x != nil {
		return x.FactoryOsUnsupported
	}
	return false
}

func (x *ResetError) GetZeroFillUnsupported() bool {
	if x != nil {
		return x.ZeroFillUnsupported
	}
	return false
}

func (x *ResetError) GetOther() bool {
	if x != nil {
		return x.Other
	}
	return false
}

func (x *ResetError) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*StartResponse_ResetSuccess
	//	*StartResponse_ResetError
	Response isStartResponse_Response `protobuf_oneof:"response"`
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_factory_reset_factory_reset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_factory_reset_factory_reset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_factory_reset_factory_reset_proto_rawDescGZIP(), []int{3}
}

func (m *StartResponse) GetResponse() isStartResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *StartResponse) GetResetSuccess() *ResetSuccess {
	if x, ok := x.GetResponse().(*StartResponse_ResetSuccess); ok {
		return x.ResetSuccess
	}
	return nil
}

func (x *StartResponse) GetResetError() *ResetError {
	if x, ok := x.GetResponse().(*StartResponse_ResetError); ok {
		return x.ResetError
	}
	return nil
}

type isStartResponse_Response interface {
	isStartResponse_Response()
}

type StartResponse_ResetSuccess struct {
	// Reset will be executed.
	ResetSuccess *ResetSuccess `protobuf:"bytes,1,opt,name=reset_success,json=resetSuccess,proto3,oneof"`
}

type StartResponse_ResetError struct {
	// Reset will not be executed.
	ResetError *ResetError `protobuf:"bytes,2,opt,name=reset_error,json=resetError,proto3,oneof"`
}

func (*StartResponse_ResetSuccess) isStartResponse_Response() {}

func (*StartResponse_ResetError) isStartResponse_Response() {}

var File_factory_reset_factory_reset_proto protoreflect.FileDescriptor

var file_factory_reset_factory_reset_proto_rawDesc = []byte{
	0x0a, 0x21, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2f,
	0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67,
	0x6e, 0x6f, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x66, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x4f, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x66, 0x69, 0x6c,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x7a, 0x65, 0x72, 0x6f, 0x46, 0x69, 0x6c,
	0x6c, 0x22, 0x0e, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0xa4, 0x01, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x34, 0x0a, 0x16, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6f, 0x73, 0x5f, 0x75,
	0x6e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x14, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4f, 0x73, 0x55, 0x6e, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x7a, 0x65, 0x72, 0x6f, 0x5f, 0x66,
	0x69, 0x6c, 0x6c, 0x5f, 0x75, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x7a, 0x65, 0x72, 0x6f, 0x46, 0x69, 0x6c, 0x6c, 0x55,
	0x6e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0xa7, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x74, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x41, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e,
	0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x52, 0x65,
	0x73, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x5c, 0x0a, 0x0c, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x12, 0x4c, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x20, 0x2e, 0x67, 0x6e,
	0x6f, 0x69, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x32, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f,
	0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f, 0x66,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0xd2, 0x3e, 0x05, 0x30,
	0x2e, 0x31, 0x2e, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_factory_reset_factory_reset_proto_rawDescOnce sync.Once
	file_factory_reset_factory_reset_proto_rawDescData = file_factory_reset_factory_reset_proto_rawDesc
)

func file_factory_reset_factory_reset_proto_rawDescGZIP() []byte {
	file_factory_reset_factory_reset_proto_rawDescOnce.Do(func() {
		file_factory_reset_factory_reset_proto_rawDescData = protoimpl.X.CompressGZIP(file_factory_reset_factory_reset_proto_rawDescData)
	})
	return file_factory_reset_factory_reset_proto_rawDescData
}

var file_factory_reset_factory_reset_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_factory_reset_factory_reset_proto_goTypes = []interface{}{
	(*StartRequest)(nil),  // 0: gnoi.factory_reset.StartRequest
	(*ResetSuccess)(nil),  // 1: gnoi.factory_reset.ResetSuccess
	(*ResetError)(nil),    // 2: gnoi.factory_reset.ResetError
	(*StartResponse)(nil), // 3: gnoi.factory_reset.StartResponse
}
var file_factory_reset_factory_reset_proto_depIdxs = []int32{
	1, // 0: gnoi.factory_reset.StartResponse.reset_success:type_name -> gnoi.factory_reset.ResetSuccess
	2, // 1: gnoi.factory_reset.StartResponse.reset_error:type_name -> gnoi.factory_reset.ResetError
	0, // 2: gnoi.factory_reset.FactoryReset.Start:input_type -> gnoi.factory_reset.StartRequest
	3, // 3: gnoi.factory_reset.FactoryReset.Start:output_type -> gnoi.factory_reset.StartResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_factory_reset_factory_reset_proto_init() }
func file_factory_reset_factory_reset_proto_init() {
	if File_factory_reset_factory_reset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_factory_reset_factory_reset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_factory_reset_factory_reset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetSuccess); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_factory_reset_factory_reset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_factory_reset_factory_reset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_factory_reset_factory_reset_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*StartResponse_ResetSuccess)(nil),
		(*StartResponse_ResetError)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_factory_reset_factory_reset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_factory_reset_factory_reset_proto_goTypes,
		DependencyIndexes: file_factory_reset_factory_reset_proto_depIdxs,
		MessageInfos:      file_factory_reset_factory_reset_proto_msgTypes,
	}.Build()
	File_factory_reset_factory_reset_proto = out.File
	file_factory_reset_factory_reset_proto_rawDesc = nil
	file_factory_reset_factory_reset_proto_goTypes = nil
	file_factory_reset_factory_reset_proto_depIdxs = nil
}
