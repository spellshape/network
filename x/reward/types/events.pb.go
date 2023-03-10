// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reward/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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

type EventRewardPoolCreated struct {
	LaunchID uint64 `protobuf:"varint,1,opt,name=launchID,proto3" json:"launchID,omitempty"`
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (m *EventRewardPoolCreated) Reset()         { *m = EventRewardPoolCreated{} }
func (m *EventRewardPoolCreated) String() string { return proto.CompactTextString(m) }
func (*EventRewardPoolCreated) ProtoMessage()    {}
func (*EventRewardPoolCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aa0ffbf8147a08e, []int{0}
}
func (m *EventRewardPoolCreated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRewardPoolCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRewardPoolCreated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRewardPoolCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRewardPoolCreated.Merge(m, src)
}
func (m *EventRewardPoolCreated) XXX_Size() int {
	return m.Size()
}
func (m *EventRewardPoolCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRewardPoolCreated.DiscardUnknown(m)
}

var xxx_messageInfo_EventRewardPoolCreated proto.InternalMessageInfo

func (m *EventRewardPoolCreated) GetLaunchID() uint64 {
	if m != nil {
		return m.LaunchID
	}
	return 0
}

func (m *EventRewardPoolCreated) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

type EventRewardPoolRemoved struct {
	LaunchID uint64 `protobuf:"varint,1,opt,name=launchID,proto3" json:"launchID,omitempty"`
}

func (m *EventRewardPoolRemoved) Reset()         { *m = EventRewardPoolRemoved{} }
func (m *EventRewardPoolRemoved) String() string { return proto.CompactTextString(m) }
func (*EventRewardPoolRemoved) ProtoMessage()    {}
func (*EventRewardPoolRemoved) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aa0ffbf8147a08e, []int{1}
}
func (m *EventRewardPoolRemoved) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRewardPoolRemoved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRewardPoolRemoved.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRewardPoolRemoved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRewardPoolRemoved.Merge(m, src)
}
func (m *EventRewardPoolRemoved) XXX_Size() int {
	return m.Size()
}
func (m *EventRewardPoolRemoved) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRewardPoolRemoved.DiscardUnknown(m)
}

var xxx_messageInfo_EventRewardPoolRemoved proto.InternalMessageInfo

func (m *EventRewardPoolRemoved) GetLaunchID() uint64 {
	if m != nil {
		return m.LaunchID
	}
	return 0
}

type EventRewardsDistributed struct {
	LaunchID uint64                                   `protobuf:"varint,1,opt,name=launchID,proto3" json:"launchID,omitempty"`
	Receiver string                                   `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Rewards  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=rewards,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"rewards"`
}

func (m *EventRewardsDistributed) Reset()         { *m = EventRewardsDistributed{} }
func (m *EventRewardsDistributed) String() string { return proto.CompactTextString(m) }
func (*EventRewardsDistributed) ProtoMessage()    {}
func (*EventRewardsDistributed) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aa0ffbf8147a08e, []int{2}
}
func (m *EventRewardsDistributed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRewardsDistributed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRewardsDistributed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRewardsDistributed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRewardsDistributed.Merge(m, src)
}
func (m *EventRewardsDistributed) XXX_Size() int {
	return m.Size()
}
func (m *EventRewardsDistributed) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRewardsDistributed.DiscardUnknown(m)
}

var xxx_messageInfo_EventRewardsDistributed proto.InternalMessageInfo

func (m *EventRewardsDistributed) GetLaunchID() uint64 {
	if m != nil {
		return m.LaunchID
	}
	return 0
}

func (m *EventRewardsDistributed) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *EventRewardsDistributed) GetRewards() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func init() {
	proto.RegisterType((*EventRewardPoolCreated)(nil), "tendermint.spn.reward.EventRewardPoolCreated")
	proto.RegisterType((*EventRewardPoolRemoved)(nil), "tendermint.spn.reward.EventRewardPoolRemoved")
	proto.RegisterType((*EventRewardsDistributed)(nil), "tendermint.spn.reward.EventRewardsDistributed")
}

func init() { proto.RegisterFile("reward/events.proto", fileDescriptor_3aa0ffbf8147a08e) }

var fileDescriptor_3aa0ffbf8147a08e = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcd, 0x4e, 0xea, 0x40,
	0x14, 0xee, 0x5c, 0x6e, 0xee, 0x4f, 0xef, 0x8e, 0x7b, 0xaf, 0x56, 0x16, 0x85, 0xb0, 0xb1, 0x0b,
	0x9d, 0x09, 0xea, 0x13, 0x00, 0x9a, 0xb8, 0x23, 0x5d, 0xea, 0xc2, 0xf4, 0xe7, 0x04, 0x26, 0xb4,
	0x73, 0x9a, 0x99, 0xa1, 0xe8, 0x5b, 0xf0, 0x1c, 0x3e, 0x09, 0x4b, 0x96, 0xae, 0xd0, 0x80, 0x4f,
	0xe1, 0xca, 0xb4, 0x53, 0x90, 0xc4, 0xc4, 0xb0, 0x3a, 0x73, 0xf2, 0xfd, 0xcc, 0x97, 0xef, 0xd8,
	0x7f, 0x25, 0x4c, 0x03, 0x19, 0x33, 0xc8, 0x41, 0x68, 0x45, 0x33, 0x89, 0x1a, 0xeb, 0xff, 0x35,
	0x88, 0x18, 0x64, 0xca, 0x85, 0xa6, 0x2a, 0x13, 0xd4, 0x70, 0x1a, 0xff, 0x86, 0x38, 0xc4, 0x92,
	0xc1, 0x8a, 0x97, 0x21, 0x37, 0x9c, 0xca, 0xc1, 0x8c, 0xbb, 0x0c, 0x31, 0xa9, 0x10, 0x37, 0x42,
	0x95, 0xa2, 0x62, 0x61, 0xa0, 0x80, 0xe5, 0x9d, 0x10, 0x74, 0xd0, 0x61, 0x11, 0x72, 0x61, 0xf0,
	0xf6, 0xc0, 0x3e, 0xb8, 0x2c, 0xbe, 0xf5, 0x4b, 0xe5, 0x00, 0x31, 0xe9, 0x49, 0x08, 0x34, 0xc4,
	0xf5, 0x86, 0xfd, 0x2b, 0x09, 0x26, 0x22, 0x1a, 0x5d, 0xf7, 0x1d, 0xd2, 0x22, 0xde, 0x77, 0x7f,
	0xbb, 0x17, 0x58, 0x26, 0x31, 0xe7, 0x31, 0x48, 0xe7, 0x5b, 0x8b, 0x78, 0xbf, 0xfd, 0xed, 0xde,
	0xbe, 0xf8, 0xe4, 0xe8, 0x43, 0x8a, 0xf9, 0xd7, 0x8e, 0xed, 0x57, 0x62, 0x1f, 0xee, 0xc8, 0x54,
	0x9f, 0x2b, 0x2d, 0x79, 0x38, 0xd9, 0x23, 0x89, 0x84, 0x08, 0x78, 0xfe, 0x91, 0x64, 0xb3, 0xd7,
	0x67, 0xc4, 0xfe, 0x69, 0x1a, 0x51, 0x4e, 0xad, 0x55, 0xf3, 0xfe, 0x9c, 0x1d, 0x51, 0x53, 0x07,
	0x2d, 0xea, 0xa0, 0x55, 0x1d, 0xb4, 0x87, 0x5c, 0x74, 0x6f, 0xe7, 0xcb, 0xa6, 0xf5, 0xb6, 0x6c,
	0x1e, 0x0f, 0xb9, 0x1e, 0x4d, 0x42, 0x1a, 0x61, 0xca, 0xaa, 0xee, 0xcc, 0x38, 0x55, 0xf1, 0x98,
	0xe9, 0x87, 0x0c, 0x54, 0x29, 0x78, 0x7c, 0x6e, 0x7a, 0x7b, 0x52, 0x95, 0xbf, 0x89, 0xd1, 0xbd,
	0x9a, 0xaf, 0x5c, 0xb2, 0x58, 0xb9, 0xe4, 0x65, 0xe5, 0x92, 0xd9, 0xda, 0xb5, 0x16, 0x6b, 0xd7,
	0x7a, 0x5a, 0xbb, 0xd6, 0xcd, 0xc9, 0x8e, 0x99, 0xca, 0x20, 0x49, 0xd4, 0x28, 0xc8, 0x80, 0x09,
	0xd0, 0x53, 0x94, 0x63, 0x76, 0x5f, 0xdd, 0xd6, 0xd8, 0x86, 0x3f, 0xca, 0xeb, 0x9d, 0xbf, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xc9, 0x53, 0x6f, 0x1a, 0x3b, 0x02, 0x00, 0x00,
}

func (m *EventRewardPoolCreated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRewardPoolCreated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRewardPoolCreated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0x12
	}
	if m.LaunchID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.LaunchID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventRewardPoolRemoved) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRewardPoolRemoved) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRewardPoolRemoved) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LaunchID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.LaunchID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventRewardsDistributed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRewardsDistributed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRewardsDistributed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	if m.LaunchID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.LaunchID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventRewardPoolCreated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LaunchID != 0 {
		n += 1 + sovEvents(uint64(m.LaunchID))
	}
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventRewardPoolRemoved) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LaunchID != 0 {
		n += 1 + sovEvents(uint64(m.LaunchID))
	}
	return n
}

func (m *EventRewardsDistributed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LaunchID != 0 {
		n += 1 + sovEvents(uint64(m.LaunchID))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventRewardPoolCreated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRewardPoolCreated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRewardPoolCreated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchID", wireType)
			}
			m.LaunchID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LaunchID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventRewardPoolRemoved) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRewardPoolRemoved: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRewardPoolRemoved: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchID", wireType)
			}
			m.LaunchID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LaunchID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventRewardsDistributed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRewardsDistributed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRewardsDistributed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchID", wireType)
			}
			m.LaunchID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LaunchID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
