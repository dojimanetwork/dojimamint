// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/p2p/types.proto

package p2p

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type NetAddress struct {
	ID   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IP   string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (m *NetAddress) Reset()         { *m = NetAddress{} }
func (m *NetAddress) String() string { return proto.CompactTextString(m) }
func (*NetAddress) ProtoMessage()    {}
func (*NetAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8a29e659aeca578, []int{0}
}
func (m *NetAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetAddress.Merge(m, src)
}
func (m *NetAddress) XXX_Size() int {
	return m.Size()
}
func (m *NetAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_NetAddress.DiscardUnknown(m)
}

var xxx_messageInfo_NetAddress proto.InternalMessageInfo

func (m *NetAddress) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *NetAddress) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *NetAddress) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type ProtocolVersion struct {
	P2P   uint64 `protobuf:"varint,1,opt,name=p2p,proto3" json:"p2p,omitempty"`
	Block uint64 `protobuf:"varint,2,opt,name=block,proto3" json:"block,omitempty"`
	App   uint64 `protobuf:"varint,3,opt,name=app,proto3" json:"app,omitempty"`
}

func (m *ProtocolVersion) Reset()         { *m = ProtocolVersion{} }
func (m *ProtocolVersion) String() string { return proto.CompactTextString(m) }
func (*ProtocolVersion) ProtoMessage()    {}
func (*ProtocolVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8a29e659aeca578, []int{1}
}
func (m *ProtocolVersion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProtocolVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProtocolVersion.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProtocolVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolVersion.Merge(m, src)
}
func (m *ProtocolVersion) XXX_Size() int {
	return m.Size()
}
func (m *ProtocolVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolVersion proto.InternalMessageInfo

func (m *ProtocolVersion) GetP2P() uint64 {
	if m != nil {
		return m.P2P
	}
	return 0
}

func (m *ProtocolVersion) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *ProtocolVersion) GetApp() uint64 {
	if m != nil {
		return m.App
	}
	return 0
}

type DefaultNodeInfo struct {
	ProtocolVersion ProtocolVersion      `protobuf:"bytes,1,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version"`
	DefaultNodeID   string               `protobuf:"bytes,2,opt,name=default_node_id,json=defaultNodeId,proto3" json:"default_node_id,omitempty"`
	ListenAddr      string               `protobuf:"bytes,3,opt,name=listen_addr,json=listenAddr,proto3" json:"listen_addr,omitempty"`
	Network         string               `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	Version         string               `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Channels        []byte               `protobuf:"bytes,6,opt,name=channels,proto3" json:"channels,omitempty"`
	Moniker         string               `protobuf:"bytes,7,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Other           DefaultNodeInfoOther `protobuf:"bytes,8,opt,name=other,proto3" json:"other"`
}

func (m *DefaultNodeInfo) Reset()         { *m = DefaultNodeInfo{} }
func (m *DefaultNodeInfo) String() string { return proto.CompactTextString(m) }
func (*DefaultNodeInfo) ProtoMessage()    {}
func (*DefaultNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8a29e659aeca578, []int{2}
}
func (m *DefaultNodeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DefaultNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DefaultNodeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DefaultNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultNodeInfo.Merge(m, src)
}
func (m *DefaultNodeInfo) XXX_Size() int {
	return m.Size()
}
func (m *DefaultNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultNodeInfo proto.InternalMessageInfo

func (m *DefaultNodeInfo) GetProtocolVersion() ProtocolVersion {
	if m != nil {
		return m.ProtocolVersion
	}
	return ProtocolVersion{}
}

func (m *DefaultNodeInfo) GetDefaultNodeID() string {
	if m != nil {
		return m.DefaultNodeID
	}
	return ""
}

func (m *DefaultNodeInfo) GetListenAddr() string {
	if m != nil {
		return m.ListenAddr
	}
	return ""
}

func (m *DefaultNodeInfo) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *DefaultNodeInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DefaultNodeInfo) GetChannels() []byte {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *DefaultNodeInfo) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *DefaultNodeInfo) GetOther() DefaultNodeInfoOther {
	if m != nil {
		return m.Other
	}
	return DefaultNodeInfoOther{}
}

type DefaultNodeInfoOther struct {
	TxIndex    string `protobuf:"bytes,1,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	RPCAddress string `protobuf:"bytes,2,opt,name=rpc_address,json=rpcAddress,proto3" json:"rpc_address,omitempty"`
}

func (m *DefaultNodeInfoOther) Reset()         { *m = DefaultNodeInfoOther{} }
func (m *DefaultNodeInfoOther) String() string { return proto.CompactTextString(m) }
func (*DefaultNodeInfoOther) ProtoMessage()    {}
func (*DefaultNodeInfoOther) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8a29e659aeca578, []int{3}
}
func (m *DefaultNodeInfoOther) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DefaultNodeInfoOther) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DefaultNodeInfoOther.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DefaultNodeInfoOther) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultNodeInfoOther.Merge(m, src)
}
func (m *DefaultNodeInfoOther) XXX_Size() int {
	return m.Size()
}
func (m *DefaultNodeInfoOther) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultNodeInfoOther.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultNodeInfoOther proto.InternalMessageInfo

func (m *DefaultNodeInfoOther) GetTxIndex() string {
	if m != nil {
		return m.TxIndex
	}
	return ""
}

func (m *DefaultNodeInfoOther) GetRPCAddress() string {
	if m != nil {
		return m.RPCAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*NetAddress)(nil), "dojimamint.p2p.NetAddress")
	proto.RegisterType((*ProtocolVersion)(nil), "dojimamint.p2p.ProtocolVersion")
	proto.RegisterType((*DefaultNodeInfo)(nil), "dojimamint.p2p.DefaultNodeInfo")
	proto.RegisterType((*DefaultNodeInfoOther)(nil), "dojimamint.p2p.DefaultNodeInfoOther")
}

func init() { proto.RegisterFile("tendermint/p2p/types.proto", fileDescriptor_c8a29e659aeca578) }

var fileDescriptor_c8a29e659aeca578 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x3d, 0x8f, 0xda, 0x40,
	0x10, 0xc5, 0xc6, 0x7c, 0xdc, 0x10, 0x8e, 0xcb, 0x0a, 0x45, 0x3e, 0x0a, 0x1b, 0xa1, 0x14, 0x54,
	0x20, 0x11, 0xa5, 0x48, 0x97, 0x10, 0x1a, 0x9a, 0x3b, 0x6b, 0x15, 0xa5, 0x48, 0x83, 0xc0, 0xbb,
	0x07, 0x1b, 0xcc, 0xee, 0x6a, 0xbd, 0x97, 0x90, 0x7f, 0x91, 0x9f, 0x75, 0xe5, 0x95, 0xa9, 0xac,
	0xc8, 0x94, 0xf9, 0x13, 0xd1, 0xae, 0x7d, 0x91, 0x0f, 0xa5, 0x9b, 0x37, 0x5f, 0x6f, 0xe6, 0xe9,
	0xc1, 0x40, 0x53, 0x4e, 0xa8, 0x3a, 0x30, 0xae, 0xa7, 0x72, 0x26, 0xa7, 0xfa, 0x87, 0xa4, 0xe9,
	0x44, 0x2a, 0xa1, 0x05, 0xba, 0x24, 0xe2, 0x2b, 0x3b, 0xac, 0x4d, 0x6d, 0x22, 0x67, 0x72, 0xd0,
	0xdf, 0x8a, 0xad, 0xb0, 0xa5, 0xa9, 0x89, 0x8a, 0xae, 0x51, 0x04, 0x70, 0x43, 0xf5, 0x07, 0x42,
	0x14, 0x4d, 0x53, 0xf4, 0x0a, 0x5c, 0x46, 0x7c, 0x67, 0xe8, 0x8c, 0x2f, 0xe6, 0xcd, 0x3c, 0x0b,
	0xdd, 0xe5, 0x02, 0xbb, 0x8c, 0xd8, 0xbc, 0xf4, 0xdd, 0x4a, 0x3e, 0xc2, 0x2e, 0x93, 0x08, 0x81,
	0x27, 0x85, 0xd2, 0x7e, 0x7d, 0xe8, 0x8c, 0xbb, 0xd8, 0xc6, 0xa3, 0x4f, 0xd0, 0x8b, 0xcc, 0xea,
	0x58, 0x24, 0x9f, 0xa9, 0x4a, 0x99, 0xe0, 0xe8, 0x1a, 0xea, 0x72, 0x26, 0xed, 0x5e, 0x6f, 0xde,
	0xca, 0xb3, 0xb0, 0x1e, 0xcd, 0x22, 0x6c, 0x72, 0xa8, 0x0f, 0x8d, 0x4d, 0x22, 0xe2, 0xbd, 0x5d,
	0xee, 0xe1, 0x02, 0xa0, 0x2b, 0xa8, 0xaf, 0xa5, 0xb4, 0x6b, 0x3d, 0x6c, 0xc2, 0xd1, 0x1f, 0x17,
	0x7a, 0x0b, 0x7a, 0xb7, 0xbe, 0x4f, 0xf4, 0x8d, 0x20, 0x74, 0xc9, 0xef, 0x04, 0x8a, 0xe0, 0x4a,
	0x96, 0x4c, 0xab, 0x6f, 0x05, 0x95, 0xe5, 0xe8, 0xcc, 0xc2, 0xc9, 0xf3, 0xe7, 0x27, 0x67, 0x17,
	0xcd, 0xbd, 0x87, 0x2c, 0xac, 0xe1, 0x9e, 0x3c, 0x3b, 0xf4, 0x1d, 0xf4, 0x48, 0x41, 0xb2, 0xe2,
	0x82, 0xd0, 0x15, 0x23, 0xe5, 0xd3, 0x2f, 0xf3, 0x2c, 0xec, 0x56, 0xf9, 0x17, 0xb8, 0x4b, 0x2a,
	0x90, 0xa0, 0x10, 0x3a, 0x09, 0x4b, 0x35, 0xe5, 0xab, 0x35, 0x21, 0xca, 0x9e, 0x7e, 0x81, 0xa1,
	0x48, 0x19, 0x79, 0x91, 0x0f, 0x2d, 0x4e, 0xf5, 0x77, 0xa1, 0xf6, 0xbe, 0x67, 0x8b, 0x4f, 0xd0,
	0x54, 0x9e, 0xce, 0x6f, 0x14, 0x95, 0x12, 0xa2, 0x01, 0xb4, 0xe3, 0xdd, 0x9a, 0x73, 0x9a, 0xa4,
	0x7e, 0x73, 0xe8, 0x8c, 0x5f, 0xe0, 0x7f, 0xd8, 0x4c, 0x1d, 0x04, 0x67, 0x7b, 0xaa, 0xfc, 0x56,
	0x31, 0x55, 0x42, 0xf4, 0x1e, 0x1a, 0x42, 0xef, 0xa8, 0xf2, 0xdb, 0x56, 0x8c, 0xd7, 0xe7, 0x62,
	0x9c, 0xe9, 0x78, 0x6b, 0x7a, 0x4b, 0x45, 0x8a, 0xc1, 0xd1, 0x06, 0xfa, 0xff, 0x6b, 0x42, 0xd7,
	0xd0, 0xd6, 0xc7, 0x15, 0xe3, 0x84, 0x1e, 0x0b, 0x97, 0xe0, 0x96, 0x3e, 0x2e, 0x0d, 0x44, 0x53,
	0xe8, 0x28, 0x19, 0xdb, 0xe7, 0x69, 0x9a, 0x96, 0xb2, 0x5d, 0xe6, 0x59, 0x08, 0x38, 0xfa, 0x58,
	0xfa, 0x0b, 0x83, 0x92, 0x71, 0x19, 0xcf, 0x6f, 0x1f, 0xf2, 0xc0, 0x79, 0xcc, 0x03, 0xe7, 0x77,
	0x1e, 0x38, 0x3f, 0x4f, 0x41, 0xed, 0xf1, 0x14, 0xd4, 0x7e, 0x9d, 0x82, 0xda, 0x97, 0xb7, 0x5b,
	0xa6, 0x77, 0xf7, 0x9b, 0x49, 0x2c, 0x0e, 0xd3, 0x8a, 0xc1, 0xab, 0x5e, 0xb7, 0x36, 0x7e, 0x6e,
	0xfe, 0x4d, 0xd3, 0x66, 0xdf, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x21, 0xa2, 0xe9, 0x0e, 0x15,
	0x03, 0x00, 0x00,
}

func (m *NetAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NetAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Port != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x18
	}
	if len(m.IP) > 0 {
		i -= len(m.IP)
		copy(dAtA[i:], m.IP)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.IP)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProtocolVersion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProtocolVersion) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProtocolVersion) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.App != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.App))
		i--
		dAtA[i] = 0x18
	}
	if m.Block != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x10
	}
	if m.P2P != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.P2P))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DefaultNodeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DefaultNodeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefaultNodeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Other.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Channels) > 0 {
		i -= len(m.Channels)
		copy(dAtA[i:], m.Channels)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Channels)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ListenAddr) > 0 {
		i -= len(m.ListenAddr)
		copy(dAtA[i:], m.ListenAddr)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ListenAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DefaultNodeID) > 0 {
		i -= len(m.DefaultNodeID)
		copy(dAtA[i:], m.DefaultNodeID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DefaultNodeID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.ProtocolVersion.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DefaultNodeInfoOther) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DefaultNodeInfoOther) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefaultNodeInfoOther) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RPCAddress) > 0 {
		i -= len(m.RPCAddress)
		copy(dAtA[i:], m.RPCAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.RPCAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TxIndex) > 0 {
		i -= len(m.TxIndex)
		copy(dAtA[i:], m.TxIndex)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.TxIndex)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NetAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.IP)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovTypes(uint64(m.Port))
	}
	return n
}

func (m *ProtocolVersion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.P2P != 0 {
		n += 1 + sovTypes(uint64(m.P2P))
	}
	if m.Block != 0 {
		n += 1 + sovTypes(uint64(m.Block))
	}
	if m.App != 0 {
		n += 1 + sovTypes(uint64(m.App))
	}
	return n
}

func (m *DefaultNodeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ProtocolVersion.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.DefaultNodeID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ListenAddr)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Channels)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Other.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *DefaultNodeInfoOther) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxIndex)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.RPCAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NetAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IP = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProtocolVersion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProtocolVersion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProtocolVersion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field P2P", wireType)
			}
			m.P2P = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.P2P |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field App", wireType)
			}
			m.App = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.App |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DefaultNodeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DefaultNodeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DefaultNodeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolVersion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ProtocolVersion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultNodeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultNodeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListenAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ListenAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channels", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channels = append(m.Channels[:0], dAtA[iNdEx:postIndex]...)
			if m.Channels == nil {
				m.Channels = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Other", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Other.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DefaultNodeInfoOther) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DefaultNodeInfoOther: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DefaultNodeInfoOther: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RPCAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RPCAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
