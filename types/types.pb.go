// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: github.com/openconfig/gnoi/types/types.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type L3Protocol int32

const (
	L3Protocol_UNSPECIFIED L3Protocol = 0
	L3Protocol_IPV4        L3Protocol = 1
	L3Protocol_IPV6        L3Protocol = 2
)

// Enum value maps for L3Protocol.
var (
	L3Protocol_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "IPV4",
		2: "IPV6",
	}
	L3Protocol_value = map[string]int32{
		"UNSPECIFIED": 0,
		"IPV4":        1,
		"IPV6":        2,
	}
)

func (x L3Protocol) Enum() *L3Protocol {
	p := new(L3Protocol)
	*p = x
	return p
}

func (x L3Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (L3Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_openconfig_gnoi_types_types_proto_enumTypes[0].Descriptor()
}

func (L3Protocol) Type() protoreflect.EnumType {
	return &file_github_com_openconfig_gnoi_types_types_proto_enumTypes[0]
}

func (x L3Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use L3Protocol.Descriptor instead.
func (L3Protocol) EnumDescriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_types_types_proto_rawDescGZIP(), []int{0}
}

type HashType_HashMethod int32

const (
	HashType_UNSPECIFIED HashType_HashMethod = 0
	HashType_SHA256      HashType_HashMethod = 1
	HashType_SHA512      HashType_HashMethod = 2
	HashType_MD5         HashType_HashMethod = 3
)

// Enum value maps for HashType_HashMethod.
var (
	HashType_HashMethod_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "SHA256",
		2: "SHA512",
		3: "MD5",
	}
	HashType_HashMethod_value = map[string]int32{
		"UNSPECIFIED": 0,
		"SHA256":      1,
		"SHA512":      2,
		"MD5":         3,
	}
)

func (x HashType_HashMethod) Enum() *HashType_HashMethod {
	p := new(HashType_HashMethod)
	*p = x
	return p
}

func (x HashType_HashMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HashType_HashMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_openconfig_gnoi_types_types_proto_enumTypes[1].Descriptor()
}

func (HashType_HashMethod) Type() protoreflect.EnumType {
	return &file_github_com_openconfig_gnoi_types_types_proto_enumTypes[1]
}

func (x HashType_HashMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HashType_HashMethod.Descriptor instead.
func (HashType_HashMethod) EnumDescriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_types_types_proto_rawDescGZIP(), []int{0, 0}
}

type HashType struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Method        HashType_HashMethod    `protobuf:"varint,1,opt,name=method,proto3,enum=gnoi.types.HashType_HashMethod" json:"method,omitempty"`
	Hash          []byte                 `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HashType) Reset() {
	*x = HashType{}
	mi := &file_github_com_openconfig_gnoi_types_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HashType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashType) ProtoMessage() {}

func (x *HashType) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnoi_types_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashType.ProtoReflect.Descriptor instead.
func (*HashType) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_types_types_proto_rawDescGZIP(), []int{0}
}

func (x *HashType) GetMethod() HashType_HashMethod {
	if x != nil {
		return x.Method
	}
	return HashType_UNSPECIFIED
}

func (x *HashType) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type Path struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Origin        string                 `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	Elem          []*PathElem            `protobuf:"bytes,3,rep,name=elem,proto3" json:"elem,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Path) Reset() {
	*x = Path{}
	mi := &file_github_com_openconfig_gnoi_types_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path) ProtoMessage() {}

func (x *Path) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnoi_types_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path.ProtoReflect.Descriptor instead.
func (*Path) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_types_types_proto_rawDescGZIP(), []int{1}
}

func (x *Path) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *Path) GetElem() []*PathElem {
	if x != nil {
		return x.Elem
	}
	return nil
}

type PathElem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key           map[string]string      `protobuf:"bytes,2,rep,name=key,proto3" json:"key,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PathElem) Reset() {
	*x = PathElem{}
	mi := &file_github_com_openconfig_gnoi_types_types_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PathElem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathElem) ProtoMessage() {}

func (x *PathElem) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnoi_types_types_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathElem.ProtoReflect.Descriptor instead.
func (*PathElem) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_types_types_proto_rawDescGZIP(), []int{2}
}

func (x *PathElem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PathElem) GetKey() map[string]string {
	if x != nil {
		return x.Key
	}
	return nil
}

type Credentials struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Username string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// Types that are valid to be assigned to Password:
	//
	//	*Credentials_Cleartext
	//	*Credentials_Hashed
	Password      isCredentials_Password `protobuf_oneof:"password"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Credentials) Reset() {
	*x = Credentials{}
	mi := &file_github_com_openconfig_gnoi_types_types_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credentials) ProtoMessage() {}

func (x *Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_openconfig_gnoi_types_types_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credentials.ProtoReflect.Descriptor instead.
func (*Credentials) Descriptor() ([]byte, []int) {
	return file_github_com_openconfig_gnoi_types_types_proto_rawDescGZIP(), []int{3}
}

func (x *Credentials) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Credentials) GetPassword() isCredentials_Password {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *Credentials) GetCleartext() string {
	if x != nil {
		if x, ok := x.Password.(*Credentials_Cleartext); ok {
			return x.Cleartext
		}
	}
	return ""
}

func (x *Credentials) GetHashed() *HashType {
	if x != nil {
		if x, ok := x.Password.(*Credentials_Hashed); ok {
			return x.Hashed
		}
	}
	return nil
}

type isCredentials_Password interface {
	isCredentials_Password()
}

type Credentials_Cleartext struct {
	Cleartext string `protobuf:"bytes,2,opt,name=cleartext,proto3,oneof"`
}

type Credentials_Hashed struct {
	Hashed *HashType `protobuf:"bytes,3,opt,name=hashed,proto3,oneof"`
}

func (*Credentials_Cleartext) isCredentials_Password() {}

func (*Credentials_Hashed) isCredentials_Password() {}

var file_github_com_openconfig_gnoi_types_types_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1002,
		Name:          "gnoi.types.gnoi_version",
		Tag:           "bytes,1002,opt,name=gnoi_version",
		Filename:      "github.com/openconfig/gnoi/types/types.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional string gnoi_version = 1002;
	E_GnoiVersion = &file_github_com_openconfig_gnoi_types_types_proto_extTypes[0]
)

var File_github_com_openconfig_gnoi_types_types_proto protoreflect.FileDescriptor

var file_github_com_openconfig_gnoi_types_types_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65,
	0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a,
	0x08, 0x48, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x67, 0x6e, 0x6f, 0x69,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x48, 0x61, 0x73, 0x68, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x3e, 0x0a, 0x0a, 0x48, 0x61, 0x73, 0x68, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x35, 0x31, 0x32, 0x10, 0x02, 0x12, 0x07, 0x0a,
	0x03, 0x4d, 0x44, 0x35, 0x10, 0x03, 0x22, 0x48, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x04, 0x65, 0x6c, 0x65, 0x6d, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x45, 0x6c, 0x65, 0x6d, 0x52, 0x04, 0x65, 0x6c, 0x65, 0x6d,
	0x22, 0x87, 0x01, 0x0a, 0x08, 0x50, 0x61, 0x74, 0x68, 0x45, 0x6c, 0x65, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2f, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x74, 0x68,
	0x45, 0x6c, 0x65, 0x6d, 0x2e, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x1a, 0x36, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x85, 0x01, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x09, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6c, 0x65,
	0x61, 0x72, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6e, 0x6f, 0x69, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x06,
	0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x2a, 0x31, 0x0a, 0x0a, 0x4c, 0x33, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x50, 0x56, 0x34, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x49,
	0x50, 0x56, 0x36, 0x10, 0x02, 0x3a, 0x40, 0x0a, 0x0c, 0x67, 0x6e, 0x6f, 0x69, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x6e, 0x6f, 0x69,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x67, 0x6e, 0x6f, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_openconfig_gnoi_types_types_proto_rawDescOnce sync.Once
	file_github_com_openconfig_gnoi_types_types_proto_rawDescData = file_github_com_openconfig_gnoi_types_types_proto_rawDesc
)

func file_github_com_openconfig_gnoi_types_types_proto_rawDescGZIP() []byte {
	file_github_com_openconfig_gnoi_types_types_proto_rawDescOnce.Do(func() {
		file_github_com_openconfig_gnoi_types_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_openconfig_gnoi_types_types_proto_rawDescData)
	})
	return file_github_com_openconfig_gnoi_types_types_proto_rawDescData
}

var file_github_com_openconfig_gnoi_types_types_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_github_com_openconfig_gnoi_types_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_openconfig_gnoi_types_types_proto_goTypes = []any{
	(L3Protocol)(0),                  // 0: gnoi.types.L3Protocol
	(HashType_HashMethod)(0),         // 1: gnoi.types.HashType.HashMethod
	(*HashType)(nil),                 // 2: gnoi.types.HashType
	(*Path)(nil),                     // 3: gnoi.types.Path
	(*PathElem)(nil),                 // 4: gnoi.types.PathElem
	(*Credentials)(nil),              // 5: gnoi.types.Credentials
	nil,                              // 6: gnoi.types.PathElem.KeyEntry
	(*descriptorpb.FileOptions)(nil), // 7: google.protobuf.FileOptions
}
var file_github_com_openconfig_gnoi_types_types_proto_depIdxs = []int32{
	1, // 0: gnoi.types.HashType.method:type_name -> gnoi.types.HashType.HashMethod
	4, // 1: gnoi.types.Path.elem:type_name -> gnoi.types.PathElem
	6, // 2: gnoi.types.PathElem.key:type_name -> gnoi.types.PathElem.KeyEntry
	2, // 3: gnoi.types.Credentials.hashed:type_name -> gnoi.types.HashType
	7, // 4: gnoi.types.gnoi_version:extendee -> google.protobuf.FileOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	4, // [4:5] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_github_com_openconfig_gnoi_types_types_proto_init() }
func file_github_com_openconfig_gnoi_types_types_proto_init() {
	if File_github_com_openconfig_gnoi_types_types_proto != nil {
		return
	}
	file_github_com_openconfig_gnoi_types_types_proto_msgTypes[3].OneofWrappers = []any{
		(*Credentials_Cleartext)(nil),
		(*Credentials_Hashed)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_openconfig_gnoi_types_types_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_github_com_openconfig_gnoi_types_types_proto_goTypes,
		DependencyIndexes: file_github_com_openconfig_gnoi_types_types_proto_depIdxs,
		EnumInfos:         file_github_com_openconfig_gnoi_types_types_proto_enumTypes,
		MessageInfos:      file_github_com_openconfig_gnoi_types_types_proto_msgTypes,
		ExtensionInfos:    file_github_com_openconfig_gnoi_types_types_proto_extTypes,
	}.Build()
	File_github_com_openconfig_gnoi_types_types_proto = out.File
	file_github_com_openconfig_gnoi_types_types_proto_rawDesc = nil
	file_github_com_openconfig_gnoi_types_types_proto_goTypes = nil
	file_github_com_openconfig_gnoi_types_types_proto_depIdxs = nil
}
