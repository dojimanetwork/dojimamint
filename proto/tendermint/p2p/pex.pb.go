// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/p2p/pex.proto

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

type PexRequest struct {
}

func (m *PexRequest) Reset()         { *m = PexRequest{} }
func (m *PexRequest) String() string { return proto.CompactTextString(m) }
func (*PexRequest) ProtoMessage()    {}
func (*PexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81c2f011fd13be57, []int{0}
}
func (m *PexRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PexRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PexRequest.Merge(m, src)
}
func (m *PexRequest) XXX_Size() int {
	return m.Size()
}
func (m *PexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PexRequest proto.InternalMessageInfo

type PexAddrs struct {
	Addrs []NetAddress `protobuf:"bytes,1,rep,name=addrs,proto3" json:"addrs"`
}

func (m *PexAddrs) Reset()         { *m = PexAddrs{} }
func (m *PexAddrs) String() string { return proto.CompactTextString(m) }
func (*PexAddrs) ProtoMessage()    {}
func (*PexAddrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_81c2f011fd13be57, []int{1}
}
func (m *PexAddrs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PexAddrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PexAddrs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PexAddrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PexAddrs.Merge(m, src)
}
func (m *PexAddrs) XXX_Size() int {
	return m.Size()
}
func (m *PexAddrs) XXX_DiscardUnknown() {
	xxx_messageInfo_PexAddrs.DiscardUnknown(m)
}

var xxx_messageInfo_PexAddrs proto.InternalMessageInfo

func (m *PexAddrs) GetAddrs() []NetAddress {
	if m != nil {
		return m.Addrs
	}
	return nil
}

type Message struct {
	// Types that are valid to be assigned to Sum:
	//	*Message_PexRequest
	//	*Message_PexAddrs
	Sum isMessage_Sum `protobuf_oneof:"sum"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_81c2f011fd13be57, []int{2}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

type isMessage_Sum interface {
	isMessage_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Message_PexRequest struct {
	PexRequest *PexRequest `protobuf:"bytes,1,opt,name=pex_request,json=pexRequest,proto3,oneof" json:"pex_request,omitempty"`
}
type Message_PexAddrs struct {
	PexAddrs *PexAddrs `protobuf:"bytes,2,opt,name=pex_addrs,json=pexAddrs,proto3,oneof" json:"pex_addrs,omitempty"`
}

func (*Message_PexRequest) isMessage_Sum() {}
func (*Message_PexAddrs) isMessage_Sum()   {}

func (m *Message) GetSum() isMessage_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Message) GetPexRequest() *PexRequest {
	if x, ok := m.GetSum().(*Message_PexRequest); ok {
		return x.PexRequest
	}
	return nil
}

func (m *Message) GetPexAddrs() *PexAddrs {
	if x, ok := m.GetSum().(*Message_PexAddrs); ok {
		return x.PexAddrs
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_PexRequest)(nil),
		(*Message_PexAddrs)(nil),
	}
}

func init() {
	proto.RegisterType((*PexRequest)(nil), "dojimamint.p2p.PexRequest")
	proto.RegisterType((*PexAddrs)(nil), "dojimamint.p2p.PexAddrs")
	proto.RegisterType((*Message)(nil), "dojimamint.p2p.Message")
}

func init() { proto.RegisterFile("tendermint/p2p/pex.proto", fileDescriptor_81c2f011fd13be57) }

var fileDescriptor_81c2f011fd13be57 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0x30, 0x2a, 0xd0, 0x2f, 0x48, 0xad, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4b, 0xc9, 0xcf, 0xca, 0xcc, 0x4d, 0x04, 0xc9, 0xe8, 0x15,
	0x18, 0x15, 0x48, 0x49, 0xa1, 0xa9, 0x2c, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0xa8, 0x95, 0x12, 0x49,
	0xcf, 0x4f, 0xcf, 0x07, 0x33, 0xf5, 0x41, 0x2c, 0x88, 0xa8, 0x12, 0x0f, 0x17, 0x57, 0x40, 0x6a,
	0x45, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x92, 0x13, 0x17, 0x47, 0x40, 0x6a, 0x85, 0x63,
	0x4a, 0x4a, 0x51, 0xb1, 0x90, 0x19, 0x17, 0x6b, 0x22, 0x88, 0x21, 0xc1, 0xa8, 0xc0, 0xac, 0xc1,
	0x6d, 0x24, 0xa5, 0x87, 0x6a, 0x97, 0x9e, 0x5f, 0x6a, 0x09, 0x48, 0x61, 0x6a, 0x71, 0xb1, 0x13,
	0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x10, 0xe5, 0x4a, 0x1d, 0x8c, 0x5c, 0xec, 0xbe, 0xa9, 0xc5,
	0xc5, 0x89, 0xe9, 0xa9, 0x42, 0xb6, 0x5c, 0xdc, 0x05, 0xa9, 0x15, 0xf1, 0x45, 0x10, 0xe3, 0x25,
	0x18, 0x15, 0x18, 0xb1, 0x99, 0x84, 0x70, 0x80, 0x07, 0x43, 0x10, 0x57, 0x01, 0x9c, 0x27, 0x64,
	0xce, 0xc5, 0x09, 0xd2, 0x0e, 0x71, 0x06, 0x13, 0x58, 0xb3, 0x04, 0x16, 0xcd, 0x60, 0xf7, 0x7a,
	0x30, 0x04, 0x71, 0x14, 0x40, 0xd9, 0x4e, 0xac, 0x5c, 0xcc, 0xc5, 0xa5, 0xb9, 0x4e, 0xfe, 0x27,
	0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c,
	0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x9a, 0x9e, 0x59, 0x92, 0x51, 0x9a,
	0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x8f, 0x14, 0x66, 0xc8, 0xc1, 0x07, 0x0e, 0x29, 0xd4, 0xf0, 0x4c,
	0x62, 0x03, 0x8b, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x66, 0x2e, 0xa4, 0xb0, 0x92, 0x01,
	0x00, 0x00,
}

func (m *PexRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PexRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PexRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *PexAddrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PexAddrs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PexAddrs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addrs) > 0 {
		for iNdEx := len(m.Addrs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Addrs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPex(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Message_PexRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_PexRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PexRequest != nil {
		{
			size, err := m.PexRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPex(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *Message_PexAddrs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message_PexAddrs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PexAddrs != nil {
		{
			size, err := m.PexAddrs.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPex(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func encodeVarintPex(dAtA []byte, offset int, v uint64) int {
	offset -= sovPex(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PexRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *PexAddrs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Addrs) > 0 {
		for _, e := range m.Addrs {
			l = e.Size()
			n += 1 + l + sovPex(uint64(l))
		}
	}
	return n
}

func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Message_PexRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PexRequest != nil {
		l = m.PexRequest.Size()
		n += 1 + l + sovPex(uint64(l))
	}
	return n
}
func (m *Message_PexAddrs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PexAddrs != nil {
		l = m.PexAddrs.Size()
		n += 1 + l + sovPex(uint64(l))
	}
	return n
}

func sovPex(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPex(x uint64) (n int) {
	return sovPex(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PexRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPex
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
			return fmt.Errorf("proto: PexRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PexRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPex
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPex
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
func (m *PexAddrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPex
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
			return fmt.Errorf("proto: PexAddrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PexAddrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPex
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
				return ErrInvalidLengthPex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addrs = append(m.Addrs, NetAddress{})
			if err := m.Addrs[len(m.Addrs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPex
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPex
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
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPex
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PexRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPex
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
				return ErrInvalidLengthPex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PexRequest{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_PexRequest{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PexAddrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPex
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
				return ErrInvalidLengthPex
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &PexAddrs{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Message_PexAddrs{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPex
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPex
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
func skipPex(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPex
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
					return 0, ErrIntOverflowPex
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
					return 0, ErrIntOverflowPex
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
				return 0, ErrInvalidLengthPex
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPex
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPex
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPex        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPex          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPex = fmt.Errorf("proto: unexpected end of group")
)
