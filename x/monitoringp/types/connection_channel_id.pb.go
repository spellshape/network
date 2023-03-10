// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: monitoringp/connection_channel_id.proto

package types

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

type ConnectionChannelID struct {
	ChannelID string `protobuf:"bytes,1,opt,name=channelID,proto3" json:"channelID,omitempty"`
}

func (m *ConnectionChannelID) Reset()         { *m = ConnectionChannelID{} }
func (m *ConnectionChannelID) String() string { return proto.CompactTextString(m) }
func (*ConnectionChannelID) ProtoMessage()    {}
func (*ConnectionChannelID) Descriptor() ([]byte, []int) {
	return fileDescriptor_554428a97e5aa49d, []int{0}
}
func (m *ConnectionChannelID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConnectionChannelID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConnectionChannelID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConnectionChannelID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionChannelID.Merge(m, src)
}
func (m *ConnectionChannelID) XXX_Size() int {
	return m.Size()
}
func (m *ConnectionChannelID) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionChannelID.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionChannelID proto.InternalMessageInfo

func (m *ConnectionChannelID) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func init() {
	proto.RegisterType((*ConnectionChannelID)(nil), "tendermint.spn.monitoringp.ConnectionChannelID")
}

func init() {
	proto.RegisterFile("monitoringp/connection_channel_id.proto", fileDescriptor_554428a97e5aa49d)
}

var fileDescriptor_554428a97e5aa49d = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcf, 0xcd, 0xcf, 0xcb,
	0x2c, 0xc9, 0x2f, 0xca, 0xcc, 0x4b, 0x2f, 0xd0, 0x4f, 0xce, 0xcf, 0xcb, 0x4b, 0x4d, 0x2e, 0xc9,
	0xcc, 0xcf, 0x8b, 0x4f, 0xce, 0x48, 0xcc, 0xcb, 0x4b, 0xcd, 0x89, 0xcf, 0x4c, 0xd1, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2a, 0x49, 0xcd, 0x4b, 0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1,
	0x2b, 0x2e, 0xc8, 0xd3, 0x43, 0xd2, 0xa7, 0x64, 0xcc, 0x25, 0xec, 0x0c, 0xd7, 0xea, 0x0c, 0xd1,
	0xe9, 0xe9, 0x22, 0x24, 0xc3, 0xc5, 0x99, 0x0c, 0xe3, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x21, 0x04, 0x9c, 0xbc, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39,
	0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x30,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0xb8, 0x20, 0x35, 0x27, 0xa7,
	0x38, 0x23, 0xb1, 0x20, 0x55, 0x3f, 0x2f, 0xb5, 0xa4, 0x3c, 0xbf, 0x28, 0x5b, 0xbf, 0x42, 0x1f,
	0xd9, 0xcd, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x47, 0x1a, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xcd, 0x7b, 0x8b, 0x38, 0xcf, 0x00, 0x00, 0x00,
}

func (m *ConnectionChannelID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConnectionChannelID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConnectionChannelID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChannelID) > 0 {
		i -= len(m.ChannelID)
		copy(dAtA[i:], m.ChannelID)
		i = encodeVarintConnectionChannelId(dAtA, i, uint64(len(m.ChannelID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConnectionChannelId(dAtA []byte, offset int, v uint64) int {
	offset -= sovConnectionChannelId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConnectionChannelID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovConnectionChannelId(uint64(l))
	}
	return n
}

func sovConnectionChannelId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConnectionChannelId(x uint64) (n int) {
	return sovConnectionChannelId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConnectionChannelID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConnectionChannelId
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
			return fmt.Errorf("proto: ConnectionChannelID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConnectionChannelID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConnectionChannelId
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
				return ErrInvalidLengthConnectionChannelId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConnectionChannelId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConnectionChannelId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConnectionChannelId
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
func skipConnectionChannelId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConnectionChannelId
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
					return 0, ErrIntOverflowConnectionChannelId
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
					return 0, ErrIntOverflowConnectionChannelId
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
				return 0, ErrInvalidLengthConnectionChannelId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConnectionChannelId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConnectionChannelId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConnectionChannelId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConnectionChannelId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConnectionChannelId = fmt.Errorf("proto: unexpected end of group")
)
