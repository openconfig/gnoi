// Code generated by protoc-gen-go.
// source: mpls/mpls.proto
// DO NOT EDIT!

/*
Package gnoi_mpls is a generated protocol buffer package.

It is generated from these files:
	mpls/mpls.proto

It has these top-level messages:
	ClearLSPRequest
	ClearLSPResponse
	ClearLSPCountersRequest
	ClearLSPCountersResponse
	MPLSPingPWEDestination
	MPLSPingRSVPTEDestination
	MPLSPingRequest
	MPLSPingResponse
*/
package gnoi_mpls

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClearLSPRequest_Mode int32

const (
	ClearLSPRequest_DEFAULT ClearLSPRequest_Mode = 0
	// Reoptimize the LSP using the current bandwidth.
	ClearLSPRequest_NONAGGRESSIVE ClearLSPRequest_Mode = 0
	// Reoptimize the LSP using the current bandwidth. Only use IGP metric in
	// calcuation.
	ClearLSPRequest_AGGRESSIVE ClearLSPRequest_Mode = 1
	// Reset and restart all LSPs that originated from this routing device.
	ClearLSPRequest_RESET ClearLSPRequest_Mode = 2
	// Apply the highest bandwidth collected on a tunnel without waiting for
	// the current application period to end. Only use IGP metric in
	// calcuation.
	ClearLSPRequest_AUTOBW_AGGRESSIVE ClearLSPRequest_Mode = 3
	// Apply the highest bandwidth collected on a tunnel without waiting for
	// the current application period to end.
	ClearLSPRequest_AUTOBW_NONAGGRESSIVE ClearLSPRequest_Mode = 4
)

var ClearLSPRequest_Mode_name = map[int32]string{
	0: "DEFAULT",
	// Duplicate value: 0: "NONAGGRESSIVE",
	1: "AGGRESSIVE",
	2: "RESET",
	3: "AUTOBW_AGGRESSIVE",
	4: "AUTOBW_NONAGGRESSIVE",
}
var ClearLSPRequest_Mode_value = map[string]int32{
	"DEFAULT":              0,
	"NONAGGRESSIVE":        0,
	"AGGRESSIVE":           1,
	"RESET":                2,
	"AUTOBW_AGGRESSIVE":    3,
	"AUTOBW_NONAGGRESSIVE": 4,
}

func (x ClearLSPRequest_Mode) String() string {
	return proto.EnumName(ClearLSPRequest_Mode_name, int32(x))
}
func (ClearLSPRequest_Mode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type MPLSPingRequest_ReplyMode int32

const (
	MPLSPingRequest_IPV4         MPLSPingRequest_ReplyMode = 0
	MPLSPingRequest_ROUTER_ALERT MPLSPingRequest_ReplyMode = 1
)

var MPLSPingRequest_ReplyMode_name = map[int32]string{
	0: "IPV4",
	1: "ROUTER_ALERT",
}
var MPLSPingRequest_ReplyMode_value = map[string]int32{
	"IPV4":         0,
	"ROUTER_ALERT": 1,
}

func (x MPLSPingRequest_ReplyMode) String() string {
	return proto.EnumName(MPLSPingRequest_ReplyMode_name, int32(x))
}
func (MPLSPingRequest_ReplyMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

type MPLSPingResponse_EchoResponseCode int32

const (
	// A successful echo response was received.
	MPLSPingResponse_SUCCESS MPLSPingResponse_EchoResponseCode = 0
	// The MPLS ping packet was not sent, for an unknown reason.
	MPLSPingResponse_NOT_SENT MPLSPingResponse_EchoResponseCode = 1
	// The local system timed out waiting for an LSP ping response.
	MPLSPingResponse_TIMEOUT MPLSPingResponse_EchoResponseCode = 2
)

var MPLSPingResponse_EchoResponseCode_name = map[int32]string{
	0: "SUCCESS",
	1: "NOT_SENT",
	2: "TIMEOUT",
}
var MPLSPingResponse_EchoResponseCode_value = map[string]int32{
	"SUCCESS":  0,
	"NOT_SENT": 1,
	"TIMEOUT":  2,
}

func (x MPLSPingResponse_EchoResponseCode) String() string {
	return proto.EnumName(MPLSPingResponse_EchoResponseCode_name, int32(x))
}
func (MPLSPingResponse_EchoResponseCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0}
}

// Request to clear a single tunnel on a target device.
type ClearLSPRequest struct {
	Name string               `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Mode ClearLSPRequest_Mode `protobuf:"varint,3,opt,name=mode,enum=gnoi.mpls.ClearLSPRequest_Mode" json:"mode,omitempty"`
}

func (m *ClearLSPRequest) Reset()                    { *m = ClearLSPRequest{} }
func (m *ClearLSPRequest) String() string            { return proto.CompactTextString(m) }
func (*ClearLSPRequest) ProtoMessage()               {}
func (*ClearLSPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClearLSPRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClearLSPRequest) GetMode() ClearLSPRequest_Mode {
	if m != nil {
		return m.Mode
	}
	return ClearLSPRequest_DEFAULT
}

type ClearLSPResponse struct {
}

func (m *ClearLSPResponse) Reset()                    { *m = ClearLSPResponse{} }
func (m *ClearLSPResponse) String() string            { return proto.CompactTextString(m) }
func (*ClearLSPResponse) ProtoMessage()               {}
func (*ClearLSPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Request to clear a single tunnel counters on a target device.
type ClearLSPCountersRequest struct {
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ClearLSPCountersRequest) Reset()                    { *m = ClearLSPCountersRequest{} }
func (m *ClearLSPCountersRequest) String() string            { return proto.CompactTextString(m) }
func (*ClearLSPCountersRequest) ProtoMessage()               {}
func (*ClearLSPCountersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClearLSPCountersRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ClearLSPCountersResponse struct {
}

func (m *ClearLSPCountersResponse) Reset()                    { *m = ClearLSPCountersResponse{} }
func (m *ClearLSPCountersResponse) String() string            { return proto.CompactTextString(m) }
func (*ClearLSPCountersResponse) ProtoMessage()               {}
func (*ClearLSPCountersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type MPLSPingPWEDestination struct {
	// The address of the egress LER that the MPLS ping should be sent on when
	// destined to a PWE service.
	Eler string `protobuf:"bytes,1,opt,name=eler" json:"eler,omitempty"`
	// The virtual circuit ID for the PWE via which the ping should be sent.
	Vcid uint32 `protobuf:"varint,2,opt,name=vcid" json:"vcid,omitempty"`
}

func (m *MPLSPingPWEDestination) Reset()                    { *m = MPLSPingPWEDestination{} }
func (m *MPLSPingPWEDestination) String() string            { return proto.CompactTextString(m) }
func (*MPLSPingPWEDestination) ProtoMessage()               {}
func (*MPLSPingPWEDestination) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MPLSPingPWEDestination) GetEler() string {
	if m != nil {
		return m.Eler
	}
	return ""
}

func (m *MPLSPingPWEDestination) GetVcid() uint32 {
	if m != nil {
		return m.Vcid
	}
	return 0
}

// MPLSPingRSVPTEDestination specifies the destination for an MPLS Ping in
// terms of an absolute specification of an RSVP-TE LSP. It can be used to
// identify an individual RSVP-TE session via which a ping should be sent.
type MPLSPingRSVPTEDestination struct {
	// The IPv4 or IPv6 address used by the system initiating (acting as the
	// head-end) of the RSVP-TE LSP.
	Src string `protobuf:"bytes,1,opt,name=src" json:"src,omitempty"`
	// The IPv4 or IPv6 address used by the system terminating (acting as the
	// tail-end) of the RSVP-TE LSP.
	Dst string `protobuf:"bytes,2,opt,name=dst" json:"dst,omitempty"`
	// The extended tunnel ID of the RSVP-TE LSP, expressed as an unsigned, 32b
	// integer.
	ExtendedTunnelId uint32 `protobuf:"varint,3,opt,name=extended_tunnel_id,json=extendedTunnelId" json:"extended_tunnel_id,omitempty"`
}

func (m *MPLSPingRSVPTEDestination) Reset()                    { *m = MPLSPingRSVPTEDestination{} }
func (m *MPLSPingRSVPTEDestination) String() string            { return proto.CompactTextString(m) }
func (*MPLSPingRSVPTEDestination) ProtoMessage()               {}
func (*MPLSPingRSVPTEDestination) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MPLSPingRSVPTEDestination) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *MPLSPingRSVPTEDestination) GetDst() string {
	if m != nil {
		return m.Dst
	}
	return ""
}

func (m *MPLSPingRSVPTEDestination) GetExtendedTunnelId() uint32 {
	if m != nil {
		return m.ExtendedTunnelId
	}
	return 0
}

// MPLSPingRequest specifies the parameters that should be used as input from
// the client, to a system that is initiating an RFC4379 MPLS ping request.
type MPLSPingRequest struct {
	// One field within the destination field should be specified to determine
	// the destination of the LSP ping.
	//
	// Types that are valid to be assigned to Destination:
	//	*MPLSPingRequest_LdpFec
	//	*MPLSPingRequest_Fec129Pwe
	//	*MPLSPingRequest_RsvpteLspName
	//	*MPLSPingRequest_RsvpteLsp
	Destination isMPLSPingRequest_Destination `protobuf_oneof:"destination"`
	ReplyMode   MPLSPingRequest_ReplyMode     `protobuf:"varint,6,opt,name=reply_mode,json=replyMode,enum=gnoi.mpls.MPLSPingRequest_ReplyMode" json:"reply_mode,omitempty"`
	// The number of MPLS echo request packets to send.
	Count uint32 `protobuf:"varint,7,opt,name=count" json:"count,omitempty"`
	// The size (in bytes) of each MPLS echo request packet.
	Size uint32 `protobuf:"varint,8,opt,name=size" json:"size,omitempty"`
	// The source IPv4 address that should be used in the request packet.
	SourceAddress string `protobuf:"bytes,9,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
	// The MPLS TTL that should be set in the packets sent.
	MplsTtl uint32 `protobuf:"varint,10,opt,name=mpls_ttl,json=mplsTtl" json:"mpls_ttl,omitempty"`
	// The value of the traffic class (TC, formerly known as EXP) bits that
	// should be set in the MPLS ping packets.
	TrafficClass uint32 `protobuf:"varint,11,opt,name=traffic_class,json=trafficClass" json:"traffic_class,omitempty"`
}

func (m *MPLSPingRequest) Reset()                    { *m = MPLSPingRequest{} }
func (m *MPLSPingRequest) String() string            { return proto.CompactTextString(m) }
func (*MPLSPingRequest) ProtoMessage()               {}
func (*MPLSPingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isMPLSPingRequest_Destination interface {
	isMPLSPingRequest_Destination()
}

type MPLSPingRequest_LdpFec struct {
	LdpFec string `protobuf:"bytes,1,opt,name=ldp_fec,json=ldpFec,oneof"`
}
type MPLSPingRequest_Fec129Pwe struct {
	Fec129Pwe *MPLSPingPWEDestination `protobuf:"bytes,2,opt,name=fec129_pwe,json=fec129Pwe,oneof"`
}
type MPLSPingRequest_RsvpteLspName struct {
	RsvpteLspName string `protobuf:"bytes,4,opt,name=rsvpte_lsp_name,json=rsvpteLspName,oneof"`
}
type MPLSPingRequest_RsvpteLsp struct {
	RsvpteLsp *MPLSPingRSVPTEDestination `protobuf:"bytes,5,opt,name=rsvpte_lsp,json=rsvpteLsp,oneof"`
}

func (*MPLSPingRequest_LdpFec) isMPLSPingRequest_Destination()        {}
func (*MPLSPingRequest_Fec129Pwe) isMPLSPingRequest_Destination()     {}
func (*MPLSPingRequest_RsvpteLspName) isMPLSPingRequest_Destination() {}
func (*MPLSPingRequest_RsvpteLsp) isMPLSPingRequest_Destination()     {}

func (m *MPLSPingRequest) GetDestination() isMPLSPingRequest_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *MPLSPingRequest) GetLdpFec() string {
	if x, ok := m.GetDestination().(*MPLSPingRequest_LdpFec); ok {
		return x.LdpFec
	}
	return ""
}

func (m *MPLSPingRequest) GetFec129Pwe() *MPLSPingPWEDestination {
	if x, ok := m.GetDestination().(*MPLSPingRequest_Fec129Pwe); ok {
		return x.Fec129Pwe
	}
	return nil
}

func (m *MPLSPingRequest) GetRsvpteLspName() string {
	if x, ok := m.GetDestination().(*MPLSPingRequest_RsvpteLspName); ok {
		return x.RsvpteLspName
	}
	return ""
}

func (m *MPLSPingRequest) GetRsvpteLsp() *MPLSPingRSVPTEDestination {
	if x, ok := m.GetDestination().(*MPLSPingRequest_RsvpteLsp); ok {
		return x.RsvpteLsp
	}
	return nil
}

func (m *MPLSPingRequest) GetReplyMode() MPLSPingRequest_ReplyMode {
	if m != nil {
		return m.ReplyMode
	}
	return MPLSPingRequest_IPV4
}

func (m *MPLSPingRequest) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *MPLSPingRequest) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *MPLSPingRequest) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *MPLSPingRequest) GetMplsTtl() uint32 {
	if m != nil {
		return m.MplsTtl
	}
	return 0
}

func (m *MPLSPingRequest) GetTrafficClass() uint32 {
	if m != nil {
		return m.TrafficClass
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MPLSPingRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MPLSPingRequest_OneofMarshaler, _MPLSPingRequest_OneofUnmarshaler, _MPLSPingRequest_OneofSizer, []interface{}{
		(*MPLSPingRequest_LdpFec)(nil),
		(*MPLSPingRequest_Fec129Pwe)(nil),
		(*MPLSPingRequest_RsvpteLspName)(nil),
		(*MPLSPingRequest_RsvpteLsp)(nil),
	}
}

func _MPLSPingRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MPLSPingRequest)
	// destination
	switch x := m.Destination.(type) {
	case *MPLSPingRequest_LdpFec:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.LdpFec)
	case *MPLSPingRequest_Fec129Pwe:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Fec129Pwe); err != nil {
			return err
		}
	case *MPLSPingRequest_RsvpteLspName:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.RsvpteLspName)
	case *MPLSPingRequest_RsvpteLsp:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RsvpteLsp); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("MPLSPingRequest.Destination has unexpected type %T", x)
	}
	return nil
}

func _MPLSPingRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MPLSPingRequest)
	switch tag {
	case 1: // destination.ldp_fec
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Destination = &MPLSPingRequest_LdpFec{x}
		return true, err
	case 2: // destination.fec129_pwe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MPLSPingPWEDestination)
		err := b.DecodeMessage(msg)
		m.Destination = &MPLSPingRequest_Fec129Pwe{msg}
		return true, err
	case 4: // destination.rsvpte_lsp_name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Destination = &MPLSPingRequest_RsvpteLspName{x}
		return true, err
	case 5: // destination.rsvpte_lsp
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MPLSPingRSVPTEDestination)
		err := b.DecodeMessage(msg)
		m.Destination = &MPLSPingRequest_RsvpteLsp{msg}
		return true, err
	default:
		return false, nil
	}
}

func _MPLSPingRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MPLSPingRequest)
	// destination
	switch x := m.Destination.(type) {
	case *MPLSPingRequest_LdpFec:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.LdpFec)))
		n += len(x.LdpFec)
	case *MPLSPingRequest_Fec129Pwe:
		s := proto.Size(x.Fec129Pwe)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *MPLSPingRequest_RsvpteLspName:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.RsvpteLspName)))
		n += len(x.RsvpteLspName)
	case *MPLSPingRequest_RsvpteLsp:
		s := proto.Size(x.RsvpteLsp)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// MPLSPingResponse (MPLSPong?) is sent from the target to the client based on
// each MPLS Echo Response packet it receives associated with an input MPLSPing
// RPC.
type MPLSPingResponse struct {
	Seq uint32 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	// The response that was received from the egress LER.
	Response MPLSPingResponse_EchoResponseCode `protobuf:"varint,2,opt,name=response,enum=gnoi.mpls.MPLSPingResponse_EchoResponseCode" json:"response,omitempty"`
	// The time (in nanoseconds) between transmission of the Echo Response,
	// and the echo reply.
	ResponseTime uint64 `protobuf:"varint,3,opt,name=response_time,json=responseTime" json:"response_time,omitempty"`
}

func (m *MPLSPingResponse) Reset()                    { *m = MPLSPingResponse{} }
func (m *MPLSPingResponse) String() string            { return proto.CompactTextString(m) }
func (*MPLSPingResponse) ProtoMessage()               {}
func (*MPLSPingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MPLSPingResponse) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *MPLSPingResponse) GetResponse() MPLSPingResponse_EchoResponseCode {
	if m != nil {
		return m.Response
	}
	return MPLSPingResponse_SUCCESS
}

func (m *MPLSPingResponse) GetResponseTime() uint64 {
	if m != nil {
		return m.ResponseTime
	}
	return 0
}

func init() {
	proto.RegisterType((*ClearLSPRequest)(nil), "gnoi.mpls.ClearLSPRequest")
	proto.RegisterType((*ClearLSPResponse)(nil), "gnoi.mpls.ClearLSPResponse")
	proto.RegisterType((*ClearLSPCountersRequest)(nil), "gnoi.mpls.ClearLSPCountersRequest")
	proto.RegisterType((*ClearLSPCountersResponse)(nil), "gnoi.mpls.ClearLSPCountersResponse")
	proto.RegisterType((*MPLSPingPWEDestination)(nil), "gnoi.mpls.MPLSPingPWEDestination")
	proto.RegisterType((*MPLSPingRSVPTEDestination)(nil), "gnoi.mpls.MPLSPingRSVPTEDestination")
	proto.RegisterType((*MPLSPingRequest)(nil), "gnoi.mpls.MPLSPingRequest")
	proto.RegisterType((*MPLSPingResponse)(nil), "gnoi.mpls.MPLSPingResponse")
	proto.RegisterEnum("gnoi.mpls.ClearLSPRequest_Mode", ClearLSPRequest_Mode_name, ClearLSPRequest_Mode_value)
	proto.RegisterEnum("gnoi.mpls.MPLSPingRequest_ReplyMode", MPLSPingRequest_ReplyMode_name, MPLSPingRequest_ReplyMode_value)
	proto.RegisterEnum("gnoi.mpls.MPLSPingResponse_EchoResponseCode", MPLSPingResponse_EchoResponseCode_name, MPLSPingResponse_EchoResponseCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MPLS service

type MPLSClient interface {
	// ClearLSP clears a single tunnel (requests for it's route to be
	// recalculated).
	ClearLSP(ctx context.Context, in *ClearLSPRequest, opts ...grpc.CallOption) (*ClearLSPResponse, error)
	// ClearLSPCounters will clear the MPLS counters for the provided LSP.
	ClearLSPCounters(ctx context.Context, in *ClearLSPCountersRequest, opts ...grpc.CallOption) (*ClearLSPCountersResponse, error)
	// An MPLS ping, specified as per RFC4379.
	MPLSPing(ctx context.Context, in *MPLSPingRequest, opts ...grpc.CallOption) (MPLS_MPLSPingClient, error)
}

type mPLSClient struct {
	cc *grpc.ClientConn
}

func NewMPLSClient(cc *grpc.ClientConn) MPLSClient {
	return &mPLSClient{cc}
}

func (c *mPLSClient) ClearLSP(ctx context.Context, in *ClearLSPRequest, opts ...grpc.CallOption) (*ClearLSPResponse, error) {
	out := new(ClearLSPResponse)
	err := grpc.Invoke(ctx, "/gnoi.mpls.MPLS/ClearLSP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPLSClient) ClearLSPCounters(ctx context.Context, in *ClearLSPCountersRequest, opts ...grpc.CallOption) (*ClearLSPCountersResponse, error) {
	out := new(ClearLSPCountersResponse)
	err := grpc.Invoke(ctx, "/gnoi.mpls.MPLS/ClearLSPCounters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mPLSClient) MPLSPing(ctx context.Context, in *MPLSPingRequest, opts ...grpc.CallOption) (MPLS_MPLSPingClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MPLS_serviceDesc.Streams[0], c.cc, "/gnoi.mpls.MPLS/MPLSPing", opts...)
	if err != nil {
		return nil, err
	}
	x := &mPLSMPLSPingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MPLS_MPLSPingClient interface {
	Recv() (*MPLSPingResponse, error)
	grpc.ClientStream
}

type mPLSMPLSPingClient struct {
	grpc.ClientStream
}

func (x *mPLSMPLSPingClient) Recv() (*MPLSPingResponse, error) {
	m := new(MPLSPingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MPLS service

type MPLSServer interface {
	// ClearLSP clears a single tunnel (requests for it's route to be
	// recalculated).
	ClearLSP(context.Context, *ClearLSPRequest) (*ClearLSPResponse, error)
	// ClearLSPCounters will clear the MPLS counters for the provided LSP.
	ClearLSPCounters(context.Context, *ClearLSPCountersRequest) (*ClearLSPCountersResponse, error)
	// An MPLS ping, specified as per RFC4379.
	MPLSPing(*MPLSPingRequest, MPLS_MPLSPingServer) error
}

func RegisterMPLSServer(s *grpc.Server, srv MPLSServer) {
	s.RegisterService(&_MPLS_serviceDesc, srv)
}

func _MPLS_ClearLSP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearLSPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPLSServer).ClearLSP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.mpls.MPLS/ClearLSP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPLSServer).ClearLSP(ctx, req.(*ClearLSPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPLS_ClearLSPCounters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearLSPCountersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MPLSServer).ClearLSPCounters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.mpls.MPLS/ClearLSPCounters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MPLSServer).ClearLSPCounters(ctx, req.(*ClearLSPCountersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MPLS_MPLSPing_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MPLSPingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MPLSServer).MPLSPing(m, &mPLSMPLSPingServer{stream})
}

type MPLS_MPLSPingServer interface {
	Send(*MPLSPingResponse) error
	grpc.ServerStream
}

type mPLSMPLSPingServer struct {
	grpc.ServerStream
}

func (x *mPLSMPLSPingServer) Send(m *MPLSPingResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _MPLS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.mpls.MPLS",
	HandlerType: (*MPLSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClearLSP",
			Handler:    _MPLS_ClearLSP_Handler,
		},
		{
			MethodName: "ClearLSPCounters",
			Handler:    _MPLS_ClearLSPCounters_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MPLSPing",
			Handler:       _MPLS_MPLSPing_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mpls/mpls.proto",
}

func init() { proto.RegisterFile("mpls/mpls.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 769 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x5f, 0x6f, 0xda, 0x48,
	0x10, 0xb7, 0xc1, 0x09, 0x30, 0x60, 0x70, 0x56, 0xb9, 0x3b, 0x87, 0x7b, 0xb8, 0x9c, 0x73, 0xa7,
	0xcb, 0xe9, 0x72, 0xa4, 0x25, 0x7d, 0x69, 0xfb, 0x52, 0x20, 0x4e, 0x82, 0x44, 0x00, 0xad, 0x4d,
	0x22, 0x55, 0xaa, 0x2c, 0x62, 0x2f, 0xc4, 0x92, 0xb1, 0x1d, 0xaf, 0x49, 0xda, 0x7e, 0xc9, 0x4a,
	0xfd, 0x22, 0x7d, 0xea, 0x7b, 0xb5, 0x6b, 0x1b, 0x02, 0xa1, 0x79, 0xb1, 0x66, 0x7f, 0x33, 0xf3,
	0x9b, 0xbf, 0x1e, 0xa8, 0xcd, 0x42, 0x8f, 0x1e, 0xb3, 0x4f, 0x23, 0x8c, 0x82, 0x38, 0x40, 0xa5,
	0xa9, 0x1f, 0xb8, 0x0d, 0x06, 0x68, 0x5f, 0x45, 0xa8, 0x75, 0x3c, 0x32, 0x8e, 0x7a, 0xc6, 0x10,
	0x93, 0xbb, 0x39, 0xa1, 0x31, 0x42, 0x20, 0xf9, 0xe3, 0x19, 0x51, 0x73, 0xfb, 0xe2, 0x61, 0x09,
	0x73, 0x19, 0x9d, 0x80, 0x34, 0x0b, 0x1c, 0xa2, 0xe6, 0xf7, 0xc5, 0xc3, 0x6a, 0xf3, 0x8f, 0xc6,
	0x82, 0xa1, 0xb1, 0xe6, 0xdd, 0xb8, 0x0c, 0x1c, 0x82, 0xb9, 0xb1, 0x76, 0x0f, 0x12, 0x7b, 0xa1,
	0x32, 0x14, 0x4e, 0xf5, 0xb3, 0xd6, 0xa8, 0x67, 0x2a, 0x02, 0xda, 0x01, 0xb9, 0x3f, 0xe8, 0xb7,
	0xce, 0xcf, 0xb1, 0x6e, 0x18, 0xdd, 0x2b, 0x5d, 0x11, 0x50, 0x15, 0xe0, 0xd1, 0x5b, 0x44, 0x25,
	0xd8, 0xc2, 0xba, 0xa1, 0x9b, 0x4a, 0x0e, 0xfd, 0x02, 0x3b, 0xad, 0x91, 0x39, 0x68, 0x5f, 0x5b,
	0x8f, 0x2c, 0xf2, 0x48, 0x85, 0xdd, 0x14, 0x5e, 0xe5, 0x92, 0xea, 0x39, 0x45, 0xd4, 0x10, 0x28,
	0xcb, 0xac, 0x68, 0x18, 0xf8, 0x94, 0x68, 0xff, 0xc3, 0x6f, 0x19, 0xd6, 0x09, 0xe6, 0x7e, 0x4c,
	0x22, 0xfa, 0x4c, 0xbd, 0x5a, 0x1d, 0xd4, 0xa7, 0xe6, 0x29, 0xd5, 0x3b, 0xf8, 0xf5, 0x72, 0xd8,
	0x33, 0x86, 0xae, 0x3f, 0x1d, 0x5e, 0xeb, 0xa7, 0x84, 0xc6, 0xae, 0x3f, 0x8e, 0xdd, 0xc0, 0x67,
	0x4c, 0xc4, 0x23, 0x91, 0x2a, 0x26, 0x4c, 0x4c, 0x66, 0xd8, 0xbd, 0xed, 0x3a, 0x9c, 0x5d, 0xc6,
	0x5c, 0xd6, 0x66, 0xb0, 0x97, 0x31, 0x60, 0xe3, 0x6a, 0x68, 0xae, 0x90, 0x28, 0x90, 0xa7, 0x91,
	0x9d, 0x72, 0x30, 0x91, 0x21, 0x0e, 0x8d, 0xd3, 0xfc, 0x98, 0x88, 0x8e, 0x00, 0x91, 0x8f, 0x31,
	0xf1, 0x1d, 0xe2, 0x58, 0xf1, 0xdc, 0xf7, 0x89, 0x67, 0xb9, 0x0e, 0x1f, 0x8e, 0x8c, 0x95, 0x4c,
	0x63, 0x72, 0x45, 0xd7, 0xd1, 0xbe, 0xe5, 0xa1, 0xb6, 0x88, 0x97, 0x16, 0xbd, 0x07, 0x05, 0xcf,
	0x09, 0xad, 0x09, 0x49, 0x23, 0x5d, 0x08, 0x78, 0xdb, 0x73, 0xc2, 0x33, 0x62, 0xa3, 0x36, 0xc0,
	0x84, 0xd8, 0x2f, 0x9b, 0xaf, 0xad, 0xf0, 0x21, 0xe9, 0x4a, 0xb9, 0xf9, 0xe7, 0xa3, 0x89, 0x6f,
	0x2e, 0xfe, 0x42, 0xc0, 0xa5, 0xc4, 0x6d, 0xf8, 0x40, 0xd0, 0x21, 0xd4, 0x22, 0x7a, 0x1f, 0xc6,
	0xc4, 0xf2, 0x68, 0x68, 0xf1, 0xf6, 0x4a, 0x69, 0x18, 0x39, 0x51, 0xf4, 0x68, 0xd8, 0x67, 0x9b,
	0xa5, 0x03, 0x2c, 0x2d, 0xd5, 0x2d, 0x1e, 0xed, 0xaf, 0x0d, 0xd1, 0x9e, 0x34, 0x8a, 0x05, 0x5c,
	0x50, 0xa1, 0x0e, 0x40, 0x44, 0x42, 0xef, 0x93, 0xc5, 0xd7, 0x74, 0x9b, 0xaf, 0xe9, 0x46, 0x9a,
	0x74, 0x4d, 0x31, 0x33, 0xe6, 0xbb, 0x5a, 0x8a, 0x32, 0x11, 0xed, 0xc2, 0x96, 0xcd, 0xa6, 0xad,
	0x16, 0x78, 0x27, 0x93, 0x07, 0x9b, 0x20, 0x75, 0x3f, 0x13, 0xb5, 0x98, 0x4c, 0x90, 0xc9, 0xe8,
	0x6f, 0xa8, 0xd2, 0x60, 0x1e, 0xd9, 0xc4, 0x1a, 0x3b, 0x4e, 0x44, 0x28, 0x55, 0x4b, 0x7c, 0x3a,
	0x72, 0x82, 0xb6, 0x12, 0x10, 0xed, 0x41, 0x91, 0x45, 0xb7, 0xe2, 0xd8, 0x53, 0x81, 0xbb, 0x17,
	0xd8, 0xdb, 0x8c, 0x3d, 0x74, 0x00, 0x72, 0x1c, 0x8d, 0x27, 0x13, 0xd7, 0xb6, 0x6c, 0x6f, 0x4c,
	0xa9, 0x5a, 0xe6, 0xfa, 0x4a, 0x0a, 0x76, 0x18, 0xa6, 0xfd, 0x03, 0xa5, 0x45, 0xa2, 0xa8, 0x08,
	0x52, 0x77, 0x78, 0xf5, 0x4a, 0x11, 0x90, 0x02, 0x15, 0x3c, 0x18, 0x99, 0x3a, 0xb6, 0x5a, 0x3d,
	0x1d, 0x9b, 0x8a, 0xd8, 0x96, 0xa1, 0xec, 0x2c, 0x5b, 0xa3, 0x7d, 0x11, 0x41, 0x59, 0x56, 0x9c,
	0xec, 0x2d, 0x5f, 0x2c, 0x72, 0xc7, 0xc7, 0x2d, 0x63, 0x26, 0xa2, 0x0b, 0x28, 0x46, 0xa9, 0x96,
	0xcf, 0xb9, 0xda, 0x3c, 0xda, 0xd8, 0xb2, 0xc4, 0xa4, 0xa1, 0xdb, 0xb7, 0x41, 0xf6, 0xe8, 0xb0,
	0xd6, 0x2d, 0xbc, 0x59, 0x35, 0x99, 0x6c, 0xc5, 0xee, 0x2c, 0x39, 0x14, 0x12, 0xae, 0x64, 0xa0,
	0xe9, 0xce, 0x88, 0xf6, 0x06, 0x94, 0x75, 0x0a, 0x76, 0x1b, 0x8c, 0x51, 0xa7, 0xa3, 0x1b, 0x86,
	0x22, 0xa0, 0x0a, 0x14, 0xfb, 0x03, 0xd3, 0x32, 0xf4, 0xbe, 0xa9, 0x88, 0x4c, 0x65, 0x76, 0x2f,
	0xf5, 0xc1, 0xc8, 0x54, 0x72, 0xcd, 0xef, 0x22, 0x48, 0x2c, 0x21, 0xa4, 0x43, 0x31, 0xfb, 0x33,
	0x51, 0xfd, 0xe7, 0x77, 0xa8, 0xfe, 0xfb, 0x46, 0x5d, 0xfa, 0x0b, 0x0b, 0xe8, 0xc3, 0xf2, 0x46,
	0x64, 0x3f, 0x38, 0xd2, 0x36, 0xb8, 0xac, 0x1d, 0x8b, 0xfa, 0xc1, 0xb3, 0x36, 0x0b, 0xfa, 0x73,
	0x28, 0x66, 0xed, 0x5b, 0xc9, 0x72, 0x6d, 0x0d, 0x57, 0xb2, 0x5c, 0xef, 0xb7, 0x26, 0xbc, 0x10,
	0xdb, 0xff, 0xbd, 0xff, 0x77, 0xea, 0xc6, 0xb7, 0xf3, 0x9b, 0x86, 0x1d, 0xcc, 0x8e, 0x83, 0x90,
	0xf8, 0x76, 0xe0, 0x4f, 0xdc, 0xe9, 0x31, 0xf3, 0xe3, 0x47, 0xfd, 0xed, 0x82, 0xe1, 0x66, 0x9b,
	0xdf, 0xf7, 0x93, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x40, 0x2e, 0x99, 0xf2, 0x05, 0x00,
	0x00,
}
