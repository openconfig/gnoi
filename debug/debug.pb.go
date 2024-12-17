// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.1
// source: github.com/openconfig/gnoi/debug/debug.proto

package debug

import (
	_ "github.com/openconfig/gnoi/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DebugRequest_Mode int32

const (
	DebugRequest_MODE_UNSPECIFIED DebugRequest_Mode = 0
	DebugRequest_MODE_SHELL       DebugRequest_Mode = 1
	DebugRequest_MODE_CLI         DebugRequest_Mode = 2
)

// Enum value maps for DebugRequest_Mode.
var (
	DebugRequest_Mode_name = map[int32]string{
		0: "MODE_UNSPECIFIED",
		1: "MODE_SHELL",
		2: "MODE_CLI",
	}
	DebugRequest_Mode_value = map[string]int32{
		"MODE_UNSPECIFIED": 0,
		"MODE_SHELL":       1,
		"MODE_CLI":         2,
	}
)

func (x DebugRequest_Mode) Enum() *DebugRequest_Mode {
	p := new(DebugRequest_Mode)
	*p = x
	return p
}

func (x DebugRequest_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DebugRequest_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_openconfig_gnoi_debug_debug_proto_enumTypes[0].Descriptor()
}

func (DebugRequest_Mode) Type() protoreflect.EnumType {
	return &file_github_com_openconfig_gnoi_debug_debug_proto_enumTypes[0]
}

func (x DebugRequest_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DebugRequest_Mode.Descriptor instead.
func (DebugRequest_Mode) EnumDescriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_debug_debug_proto_rawDescGZIP(), []int{0, 0}
}

type DebugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode        DebugRequest_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=gnoi.debug.DebugRequest_Mode" json:"mode,omitempty"`
	Command     []byte            `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	ByteLimit   int64             `protobuf:"varint,3,opt,name=byte_limit,json=byteLimit,proto3" json:"byte_limit,omitempty"`
	Timeout     int64             `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	RoleAccount string            `protobuf:"bytes,5,opt,name=role_account,json=roleAccount,proto3" json:"role_account,omitempty"`
}

func (x *DebugRequest) Reset() {
	*x = DebugRequest{}
	mi := &file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DebugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugRequest) ProtoMessage() {}

func (x *DebugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugRequest.ProtoReflect.Descriptor instead.
func (*DebugRequest) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_debug_debug_proto_rawDescGZIP(), []int{0}
}

func (x *DebugRequest) GetMode() DebugRequest_Mode {
	if x != nil {
		return x.Mode
	}
	return DebugRequest_MODE_UNSPECIFIED
}

func (x *DebugRequest) GetCommand() []byte {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *DebugRequest) GetByteLimit() int64 {
	if x != nil {
		return x.ByteLimit
	}
	return 0
}

func (x *DebugRequest) GetTimeout() int64 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *DebugRequest) GetRoleAccount() string {
	if x != nil {
		return x.RoleAccount
	}
	return ""
}

type DebugResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*DebugResponse_Request
	//	*DebugResponse_Data
	//	*DebugResponse_Status
	Response isDebugResponse_Response `protobuf_oneof:"response"`
}

func (x *DebugResponse) Reset() {
	*x = DebugResponse{}
	mi := &file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DebugResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugResponse) ProtoMessage() {}

func (x *DebugResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugResponse.ProtoReflect.Descriptor instead.
func (*DebugResponse) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_debug_debug_proto_rawDescGZIP(), []int{1}
}

func (m *DebugResponse) GetResponse() isDebugResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *DebugResponse) GetRequest() *DebugRequest {
	if x, ok := x.GetResponse().(*DebugResponse_Request); ok {
		return x.Request
	}
	return nil
}

func (x *DebugResponse) GetData() []byte {
	if x, ok := x.GetResponse().(*DebugResponse_Data); ok {
		return x.Data
	}
	return nil
}

func (x *DebugResponse) GetStatus() *DebugStatus {
	if x, ok := x.GetResponse().(*DebugResponse_Status); ok {
		return x.Status
	}
	return nil
}

type isDebugResponse_Response interface {
	isDebugResponse_Response()
}

type DebugResponse_Request struct {
	Request *DebugRequest `protobuf:"bytes,100,opt,name=request,proto3,oneof"`
}

type DebugResponse_Data struct {
	Data []byte `protobuf:"bytes,101,opt,name=data,proto3,oneof"`
}

type DebugResponse_Status struct {
	Status *DebugStatus `protobuf:"bytes,102,opt,name=status,proto3,oneof"`
}

func (*DebugResponse_Request) isDebugResponse_Response() {}

func (*DebugResponse_Data) isDebugResponse_Response() {}

func (*DebugResponse_Status) isDebugResponse_Response() {}

type DebugStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details []*anypb.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *DebugStatus) Reset() {
	*x = DebugStatus{}
	mi := &file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DebugStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DebugStatus) ProtoMessage() {}

func (x *DebugStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DebugStatus.ProtoReflect.Descriptor instead.
func (*DebugStatus) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_debug_debug_proto_rawDescGZIP(), []int{2}
}

func (x *DebugStatus) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DebugStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DebugStatus) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_github_com_openconfig_gnoi_debug_debug_proto protoreflect.FileDescriptor

var file_github_com_openconfig_gnoi_debug_debug_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x1a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e,
	0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x6f,
	0x6c, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x72, 0x6f, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3a, 0x0a,
	0x04, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4d,
	0x4f, 0x44, 0x45, 0x5f, 0x53, 0x48, 0x45, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4d,
	0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x10, 0x02, 0x22, 0x9a, 0x01, 0x0a, 0x0d, 0x44, 0x65,
	0x62, 0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67,
	0x6e, 0x6f, 0x69, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x64,
	0x65, 0x62, 0x75, 0x67, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x0a, 0x0b, 0x44, 0x65, 0x62, 0x75, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x32, 0x47, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x3e, 0x0a, 0x05,
	0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x18, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x2e, 0x44, 0x65, 0x62, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x44, 0x65, 0x62,
	0x75, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x2a, 0xd2, 0x3e,
	0x05, 0x30, 0x2e, 0x31, 0x2e, 0x30, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e,
	0x6f, 0x69, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_openconfig_gnoi_debug_debug_proto_rawDescOnce sync.Once
	file_github_com_openconfig_gnoi_debug_debug_proto_rawDescData = file_github_com_openconfig_gnoi_debug_debug_proto_rawDesc
)

func file_github_com_openconfig_gnoi_debug_debug_proto_rawDescGZIP() []byte {
	file_github_com_openconfig_gnoi_debug_debug_proto_rawDescOnce.Do(func() {
		file_github_com_openconfig_gnoi_debug_debug_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_openconfig_gnoi_debug_debug_proto_rawDescData)
	})
	return file_github_com_openconfig_gnoi_debug_debug_proto_rawDescData
}

var file_github_com_openconfig_gnoi_debug_debug_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_openconfig_gnoi_debug_debug_proto_goTypes = []any{
	(DebugRequest_Mode)(0), // 0: gnoi.debug.DebugRequest.Mode
	(*DebugRequest)(nil),   // 1: gnoi.debug.DebugRequest
	(*DebugResponse)(nil),  // 2: gnoi.debug.DebugResponse
	(*DebugStatus)(nil),    // 3: gnoi.debug.DebugStatus
	(*anypb.Any)(nil),      // 4: google.protobuf.Any
}
var file_github_com_openconfig_gnoi_debug_debug_proto_depIdxs = []int32{
	0, // 0: gnoi.debug.DebugRequest.mode:type_name -> gnoi.debug.DebugRequest.Mode
	1, // 1: gnoi.debug.DebugResponse.request:type_name -> gnoi.debug.DebugRequest
	3, // 2: gnoi.debug.DebugResponse.status:type_name -> gnoi.debug.DebugStatus
	4, // 3: gnoi.debug.DebugStatus.details:type_name -> google.protobuf.Any
	1, // 4: gnoi.debug.Debug.Debug:input_type -> gnoi.debug.DebugRequest
	2, // 5: gnoi.debug.Debug.Debug:output_type -> gnoi.debug.DebugResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_openconfig_gnoi_debug_debug_proto_init() }
func file_github_com_openconfig_gnoi_debug_debug_proto_init() {
	if File_github_com_openconfig_gnoi_debug_debug_proto != nil {
		return
	}
	file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes[1].OneofWrappers = []any{
		(*DebugResponse_Request)(nil),
		(*DebugResponse_Data)(nil),
		(*DebugResponse_Status)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_openconfig_gnoi_debug_debug_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_openconfig_gnoi_debug_debug_proto_goTypes,
		DependencyIndexes: file_github_com_openconfig_gnoi_debug_debug_proto_depIdxs,
		EnumInfos:         file_github_com_openconfig_gnoi_debug_debug_proto_enumTypes,
		MessageInfos:      file_github_com_openconfig_gnoi_debug_debug_proto_msgTypes,
	}.Build()
	File_github_com_openconfig_gnoi_debug_debug_proto = out.File
	file_github_com_openconfig_gnoi_debug_debug_proto_rawDesc = nil
	file_github_com_openconfig_gnoi_debug_debug_proto_goTypes = nil
	file_github_com_openconfig_gnoi_debug_debug_proto_depIdxs = nil
}
