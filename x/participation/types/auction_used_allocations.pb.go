// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: participation/auction_used_allocations.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Allocations used by a user for a specific auction
type AuctionUsedAllocations struct {
	Address        string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	AuctionID      uint64                                 `protobuf:"varint,2,opt,name=auctionID,proto3" json:"auctionID,omitempty"`
	Withdrawn      bool                                   `protobuf:"varint,3,opt,name=withdrawn,proto3" json:"withdrawn,omitempty"`
	NumAllocations github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=numAllocations,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"numAllocations"`
}

func (m *AuctionUsedAllocations) Reset()         { *m = AuctionUsedAllocations{} }
func (m *AuctionUsedAllocations) String() string { return proto.CompactTextString(m) }
func (*AuctionUsedAllocations) ProtoMessage()    {}
func (*AuctionUsedAllocations) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dcb8de466150226, []int{0}
}
func (m *AuctionUsedAllocations) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuctionUsedAllocations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuctionUsedAllocations.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuctionUsedAllocations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuctionUsedAllocations.Merge(m, src)
}
func (m *AuctionUsedAllocations) XXX_Size() int {
	return m.Size()
}
func (m *AuctionUsedAllocations) XXX_DiscardUnknown() {
	xxx_messageInfo_AuctionUsedAllocations.DiscardUnknown(m)
}

var xxx_messageInfo_AuctionUsedAllocations proto.InternalMessageInfo

func (m *AuctionUsedAllocations) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AuctionUsedAllocations) GetAuctionID() uint64 {
	if m != nil {
		return m.AuctionID
	}
	return 0
}

func (m *AuctionUsedAllocations) GetWithdrawn() bool {
	if m != nil {
		return m.Withdrawn
	}
	return false
}

func init() {
	proto.RegisterType((*AuctionUsedAllocations)(nil), "tendermint.spn.participation.AuctionUsedAllocations")
}

func init() {
	proto.RegisterFile("participation/auction_used_allocations.proto", fileDescriptor_0dcb8de466150226)
}

var fileDescriptor_0dcb8de466150226 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0x31, 0x4f, 0xc2, 0x40,
	0x14, 0xee, 0x29, 0x51, 0xe9, 0xe0, 0xd0, 0x10, 0x53, 0x09, 0x29, 0xc4, 0xc1, 0x30, 0x48, 0x9b,
	0xc8, 0xea, 0x02, 0x71, 0x61, 0x70, 0xa9, 0x71, 0x71, 0x21, 0x47, 0xef, 0x52, 0x2e, 0xb4, 0x77,
	0x97, 0x7b, 0xaf, 0x41, 0xff, 0x85, 0x3f, 0x86, 0x1f, 0xc1, 0x48, 0x98, 0x8c, 0x03, 0x31, 0xb0,
	0xfb, 0x1b, 0x0c, 0x6d, 0x4d, 0xc1, 0xe9, 0xdd, 0x7d, 0xdf, 0xfb, 0xde, 0xf7, 0xdd, 0x3d, 0xfb,
	0x4e, 0x53, 0x83, 0x22, 0x12, 0x9a, 0xa2, 0x50, 0x32, 0xa0, 0x59, 0xb4, 0xaf, 0xe3, 0x0c, 0x38,
	0x1b, 0xd3, 0x24, 0x51, 0x51, 0x8e, 0x83, 0xaf, 0x8d, 0x42, 0xe5, 0xb4, 0x90, 0x4b, 0xc6, 0x4d,
	0x2a, 0x24, 0xfa, 0xa0, 0xa5, 0x7f, 0x24, 0x6e, 0x36, 0x62, 0x15, 0xab, 0xbc, 0x31, 0xd8, 0x9f,
	0x0a, 0x4d, 0xf3, 0x3a, 0x52, 0x90, 0x2a, 0x18, 0x17, 0x44, 0x71, 0x29, 0xa8, 0x9b, 0x1f, 0x62,
	0x5f, 0x0d, 0x0a, 0xc7, 0x17, 0xe0, 0x6c, 0x50, 0xf9, 0x39, 0xf7, 0xf6, 0x39, 0x65, 0xcc, 0x70,
	0x00, 0x97, 0x74, 0x48, 0xb7, 0x3e, 0x74, 0xd7, 0x8b, 0x5e, 0xa3, 0x54, 0x0f, 0x0a, 0xe6, 0x19,
	0x8d, 0x90, 0x71, 0xf8, 0xd7, 0xe8, 0xb4, 0xec, 0x7a, 0x99, 0x7f, 0xf4, 0xe8, 0x9e, 0x74, 0x48,
	0xb7, 0x16, 0x56, 0xc0, 0x9e, 0x9d, 0x0b, 0x9c, 0x32, 0x43, 0xe7, 0xd2, 0x3d, 0xed, 0x90, 0xee,
	0x45, 0x58, 0x01, 0x0e, 0xb3, 0x2f, 0x65, 0x96, 0x1e, 0x24, 0x70, 0x6b, 0xb9, 0xed, 0xc3, 0x72,
	0xd3, 0xb6, 0xbe, 0x36, 0xed, 0xdb, 0x58, 0xe0, 0x34, 0x9b, 0xf8, 0x91, 0x4a, 0xcb, 0x37, 0x94,
	0xa5, 0x07, 0x6c, 0x16, 0xe0, 0xbb, 0xe6, 0xe0, 0x8f, 0x24, 0xae, 0x17, 0x3d, 0xbb, 0x0c, 0x39,
	0x92, 0x18, 0xfe, 0x9b, 0x39, 0x7c, 0x5a, 0x6e, 0x3d, 0xb2, 0xda, 0x7a, 0xe4, 0x7b, 0xeb, 0x91,
	0x8f, 0x9d, 0x67, 0xad, 0x76, 0x9e, 0xf5, 0xb9, 0xf3, 0xac, 0xd7, 0xfe, 0xc1, 0x7c, 0xd0, 0x3c,
	0x49, 0x60, 0x4a, 0x35, 0x0f, 0x24, 0xc7, 0xb9, 0x32, 0xb3, 0xe0, 0x2d, 0x38, 0xde, 0x53, 0x6e,
	0x38, 0x39, 0xcb, 0xbf, 0xb1, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x00, 0xd3, 0xfe, 0xc5,
	0x01, 0x00, 0x00,
}

func (m *AuctionUsedAllocations) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuctionUsedAllocations) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuctionUsedAllocations) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.NumAllocations.Size()
		i -= size
		if _, err := m.NumAllocations.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuctionUsedAllocations(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Withdrawn {
		i--
		if m.Withdrawn {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.AuctionID != 0 {
		i = encodeVarintAuctionUsedAllocations(dAtA, i, uint64(m.AuctionID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAuctionUsedAllocations(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuctionUsedAllocations(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuctionUsedAllocations(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuctionUsedAllocations) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAuctionUsedAllocations(uint64(l))
	}
	if m.AuctionID != 0 {
		n += 1 + sovAuctionUsedAllocations(uint64(m.AuctionID))
	}
	if m.Withdrawn {
		n += 2
	}
	l = m.NumAllocations.Size()
	n += 1 + l + sovAuctionUsedAllocations(uint64(l))
	return n
}

func sovAuctionUsedAllocations(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuctionUsedAllocations(x uint64) (n int) {
	return sovAuctionUsedAllocations(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuctionUsedAllocations) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuctionUsedAllocations
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
			return fmt.Errorf("proto: AuctionUsedAllocations: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuctionUsedAllocations: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuctionUsedAllocations
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
				return ErrInvalidLengthAuctionUsedAllocations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuctionUsedAllocations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionID", wireType)
			}
			m.AuctionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuctionUsedAllocations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdrawn", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuctionUsedAllocations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Withdrawn = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumAllocations", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuctionUsedAllocations
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
				return ErrInvalidLengthAuctionUsedAllocations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuctionUsedAllocations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NumAllocations.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuctionUsedAllocations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuctionUsedAllocations
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
func skipAuctionUsedAllocations(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuctionUsedAllocations
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
					return 0, ErrIntOverflowAuctionUsedAllocations
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
					return 0, ErrIntOverflowAuctionUsedAllocations
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
				return 0, ErrInvalidLengthAuctionUsedAllocations
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuctionUsedAllocations
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuctionUsedAllocations
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuctionUsedAllocations        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuctionUsedAllocations          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuctionUsedAllocations = fmt.Errorf("proto: unexpected end of group")
)
