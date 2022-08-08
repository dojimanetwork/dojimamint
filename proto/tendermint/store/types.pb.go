// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tendermint/store/types.proto

package store

import (
	fmt "fmt"
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

type BlockStoreState struct {
	Base   int64 `protobuf:"varint,1,opt,name=base,proto3" json:"base,omitempty"`
	Height int64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *BlockStoreState) Reset()         { *m = BlockStoreState{} }
func (m *BlockStoreState) String() string { return proto.CompactTextString(m) }
func (*BlockStoreState) ProtoMessage()    {}
func (*BlockStoreState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9e53a0a74267f7, []int{0}
}
func (m *BlockStoreState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockStoreState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockStoreState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockStoreState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockStoreState.Merge(m, src)
}
func (m *BlockStoreState) XXX_Size() int {
	return m.Size()
}
func (m *BlockStoreState) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockStoreState.DiscardUnknown(m)
}

var xxx_messageInfo_BlockStoreState proto.InternalMessageInfo

func (m *BlockStoreState) GetBase() int64 {
	if m != nil {
		return m.Base
	}
	return 0
}

func (m *BlockStoreState) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*BlockStoreState)(nil), "tendermint.store.BlockStoreState")
}

func init() { proto.RegisterFile("tendermint/store/types.proto", fileDescriptor_ff9e53a0a74267f7) }

var fileDescriptor_ff9e53a0a74267f7 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x49, 0xcd, 0x4b,
	0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0x2f, 0xa9, 0x2c,
	0x48, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x40, 0xc8, 0xea, 0x81, 0x65, 0x95,
	0x6c, 0xb9, 0xf8, 0x9d, 0x72, 0xf2, 0x93, 0xb3, 0x83, 0x41, 0xbc, 0xe0, 0x92, 0xc4, 0x92, 0x54,
	0x21, 0x21, 0x2e, 0x96, 0xa4, 0xc4, 0xe2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x30,
	0x5b, 0x48, 0x8c, 0x8b, 0x2d, 0x23, 0x35, 0x33, 0x3d, 0xa3, 0x44, 0x82, 0x09, 0x2c, 0x0a, 0xe5,
	0x39, 0x85, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13,
	0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x55, 0x7a, 0x66,
	0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x4a, 0x7e, 0x56, 0x66, 0x6e, 0x62, 0x5e,
	0x6a, 0x49, 0x79, 0x7e, 0x51, 0x36, 0x94, 0x07, 0x76, 0x21, 0xd8, 0x55, 0xfa, 0xe8, 0x4e, 0x4e,
	0x62, 0x03, 0x8b, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x89, 0xc3, 0x1c, 0xd9, 0xcd, 0x00,
	0x00, 0x00,
}

func (m *BlockStoreState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockStoreState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockStoreState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if m.Base != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Base))
		i--
		dAtA[i] = 0x8
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
func (m *BlockStoreState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Base != 0 {
		n += 1 + sovTypes(uint64(m.Base))
	}
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockStoreState) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: BlockStoreState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockStoreState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Base", wireType)
			}
			m.Base = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Base |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
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
