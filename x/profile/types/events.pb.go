// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: profile/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type EventCoordinatorCreated struct {
	CoordinatorID uint64 `protobuf:"varint,1,opt,name=coordinatorID,proto3" json:"coordinatorID,omitempty"`
	Address       string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *EventCoordinatorCreated) Reset()         { *m = EventCoordinatorCreated{} }
func (m *EventCoordinatorCreated) String() string { return proto.CompactTextString(m) }
func (*EventCoordinatorCreated) ProtoMessage()    {}
func (*EventCoordinatorCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f195f0d2c25dc7b, []int{0}
}
func (m *EventCoordinatorCreated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCoordinatorCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCoordinatorCreated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCoordinatorCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCoordinatorCreated.Merge(m, src)
}
func (m *EventCoordinatorCreated) XXX_Size() int {
	return m.Size()
}
func (m *EventCoordinatorCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCoordinatorCreated.DiscardUnknown(m)
}

var xxx_messageInfo_EventCoordinatorCreated proto.InternalMessageInfo

func (m *EventCoordinatorCreated) GetCoordinatorID() uint64 {
	if m != nil {
		return m.CoordinatorID
	}
	return 0
}

func (m *EventCoordinatorCreated) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type EventCoordinatorAddressUpdated struct {
	CoordinatorID uint64 `protobuf:"varint,1,opt,name=coordinatorID,proto3" json:"coordinatorID,omitempty"`
	NewAddress    string `protobuf:"bytes,2,opt,name=newAddress,proto3" json:"newAddress,omitempty"`
}

func (m *EventCoordinatorAddressUpdated) Reset()         { *m = EventCoordinatorAddressUpdated{} }
func (m *EventCoordinatorAddressUpdated) String() string { return proto.CompactTextString(m) }
func (*EventCoordinatorAddressUpdated) ProtoMessage()    {}
func (*EventCoordinatorAddressUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f195f0d2c25dc7b, []int{1}
}
func (m *EventCoordinatorAddressUpdated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCoordinatorAddressUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCoordinatorAddressUpdated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCoordinatorAddressUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCoordinatorAddressUpdated.Merge(m, src)
}
func (m *EventCoordinatorAddressUpdated) XXX_Size() int {
	return m.Size()
}
func (m *EventCoordinatorAddressUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCoordinatorAddressUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_EventCoordinatorAddressUpdated proto.InternalMessageInfo

func (m *EventCoordinatorAddressUpdated) GetCoordinatorID() uint64 {
	if m != nil {
		return m.CoordinatorID
	}
	return 0
}

func (m *EventCoordinatorAddressUpdated) GetNewAddress() string {
	if m != nil {
		return m.NewAddress
	}
	return ""
}

type EventCoordinatorDisabled struct {
	CoordinatorID uint64 `protobuf:"varint,1,opt,name=coordinatorID,proto3" json:"coordinatorID,omitempty"`
	Address       string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *EventCoordinatorDisabled) Reset()         { *m = EventCoordinatorDisabled{} }
func (m *EventCoordinatorDisabled) String() string { return proto.CompactTextString(m) }
func (*EventCoordinatorDisabled) ProtoMessage()    {}
func (*EventCoordinatorDisabled) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f195f0d2c25dc7b, []int{2}
}
func (m *EventCoordinatorDisabled) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCoordinatorDisabled) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCoordinatorDisabled.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCoordinatorDisabled) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCoordinatorDisabled.Merge(m, src)
}
func (m *EventCoordinatorDisabled) XXX_Size() int {
	return m.Size()
}
func (m *EventCoordinatorDisabled) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCoordinatorDisabled.DiscardUnknown(m)
}

var xxx_messageInfo_EventCoordinatorDisabled proto.InternalMessageInfo

func (m *EventCoordinatorDisabled) GetCoordinatorID() uint64 {
	if m != nil {
		return m.CoordinatorID
	}
	return 0
}

func (m *EventCoordinatorDisabled) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type EventValidatorCreated struct {
	Address           string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	OperatorAddresses []string `protobuf:"bytes,2,rep,name=operatorAddresses,proto3" json:"operatorAddresses,omitempty"`
}

func (m *EventValidatorCreated) Reset()         { *m = EventValidatorCreated{} }
func (m *EventValidatorCreated) String() string { return proto.CompactTextString(m) }
func (*EventValidatorCreated) ProtoMessage()    {}
func (*EventValidatorCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f195f0d2c25dc7b, []int{3}
}
func (m *EventValidatorCreated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventValidatorCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventValidatorCreated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventValidatorCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventValidatorCreated.Merge(m, src)
}
func (m *EventValidatorCreated) XXX_Size() int {
	return m.Size()
}
func (m *EventValidatorCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventValidatorCreated.DiscardUnknown(m)
}

var xxx_messageInfo_EventValidatorCreated proto.InternalMessageInfo

func (m *EventValidatorCreated) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventValidatorCreated) GetOperatorAddresses() []string {
	if m != nil {
		return m.OperatorAddresses
	}
	return nil
}

type EventValidatorOperatorAddressesUpdated struct {
	Address           string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	OperatorAddresses []string `protobuf:"bytes,2,rep,name=operatorAddresses,proto3" json:"operatorAddresses,omitempty"`
}

func (m *EventValidatorOperatorAddressesUpdated) Reset() {
	*m = EventValidatorOperatorAddressesUpdated{}
}
func (m *EventValidatorOperatorAddressesUpdated) String() string { return proto.CompactTextString(m) }
func (*EventValidatorOperatorAddressesUpdated) ProtoMessage()    {}
func (*EventValidatorOperatorAddressesUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f195f0d2c25dc7b, []int{4}
}
func (m *EventValidatorOperatorAddressesUpdated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventValidatorOperatorAddressesUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventValidatorOperatorAddressesUpdated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventValidatorOperatorAddressesUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventValidatorOperatorAddressesUpdated.Merge(m, src)
}
func (m *EventValidatorOperatorAddressesUpdated) XXX_Size() int {
	return m.Size()
}
func (m *EventValidatorOperatorAddressesUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventValidatorOperatorAddressesUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_EventValidatorOperatorAddressesUpdated proto.InternalMessageInfo

func (m *EventValidatorOperatorAddressesUpdated) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventValidatorOperatorAddressesUpdated) GetOperatorAddresses() []string {
	if m != nil {
		return m.OperatorAddresses
	}
	return nil
}

func init() {
	proto.RegisterType((*EventCoordinatorCreated)(nil), "tendermint.spn.profile.EventCoordinatorCreated")
	proto.RegisterType((*EventCoordinatorAddressUpdated)(nil), "tendermint.spn.profile.EventCoordinatorAddressUpdated")
	proto.RegisterType((*EventCoordinatorDisabled)(nil), "tendermint.spn.profile.EventCoordinatorDisabled")
	proto.RegisterType((*EventValidatorCreated)(nil), "tendermint.spn.profile.EventValidatorCreated")
	proto.RegisterType((*EventValidatorOperatorAddressesUpdated)(nil), "tendermint.spn.profile.EventValidatorOperatorAddressesUpdated")
}

func init() { proto.RegisterFile("profile/events.proto", fileDescriptor_2f195f0d2c25dc7b) }

var fileDescriptor_2f195f0d2c25dc7b = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x3f, 0x4f, 0xc2, 0x40,
	0x18, 0xc6, 0x39, 0x35, 0x1a, 0x2e, 0x71, 0xb0, 0x41, 0xa9, 0x0c, 0x17, 0xd2, 0x18, 0xc3, 0x20,
	0x6d, 0xa2, 0x8b, 0x2b, 0x7f, 0x8c, 0x71, 0x32, 0xc1, 0xe8, 0xe0, 0x62, 0x0a, 0xf7, 0x5a, 0x2e,
	0x96, 0xbb, 0xcb, 0xdd, 0x09, 0xb2, 0x39, 0x3b, 0xf9, 0x61, 0xfc, 0x10, 0x8e, 0xc4, 0xc9, 0xd1,
	0xc0, 0x17, 0x31, 0xd0, 0x5e, 0x28, 0xb8, 0xe0, 0xc0, 0xd6, 0xf6, 0xf9, 0xf5, 0x79, 0x9e, 0x7b,
	0xef, 0xc5, 0x05, 0xa9, 0xc4, 0x23, 0x8b, 0x21, 0x80, 0x3e, 0x70, 0xa3, 0x7d, 0xa9, 0x84, 0x11,
	0xce, 0x81, 0x01, 0x4e, 0x41, 0xf5, 0x18, 0x37, 0xbe, 0x96, 0xdc, 0x4f, 0xa1, 0x52, 0x21, 0x12,
	0x91, 0x98, 0x21, 0xc1, 0xf4, 0x29, 0xa1, 0x4b, 0x87, 0x1d, 0xa1, 0x7b, 0x42, 0x3f, 0x24, 0x42,
	0xf2, 0x62, 0x25, 0x6b, 0xdf, 0x11, 0x42, 0x51, 0xc6, 0x43, 0x23, 0x54, 0x2a, 0x15, 0xad, 0xd4,
	0x0f, 0x63, 0x46, 0xe7, 0x82, 0xa7, 0x71, 0xf1, 0x62, 0x5a, 0xa6, 0x31, 0xff, 0xa5, 0xa1, 0x20,
	0x34, 0x40, 0x9d, 0x23, 0xbc, 0x9b, 0x31, 0xba, 0x6a, 0xba, 0xa8, 0x8c, 0x2a, 0x5b, 0xad, 0xc5,
	0x8f, 0xce, 0x29, 0xde, 0x09, 0x29, 0x55, 0xa0, 0xb5, 0xbb, 0x51, 0x46, 0x95, 0x7c, 0xdd, 0xfd,
	0xfa, 0xa8, 0x16, 0xd2, 0x5e, 0xb5, 0x44, 0xb9, 0x31, 0x8a, 0xf1, 0xa8, 0x65, 0x41, 0xef, 0x15,
	0x61, 0xb2, 0x9c, 0x9a, 0xa2, 0xb7, 0x92, 0xfe, 0x23, 0xfc, 0x1c, 0x63, 0x0e, 0x83, 0xda, 0x8a,
	0xf9, 0x19, 0xd6, 0x33, 0xd8, 0x5d, 0x6e, 0xd0, 0x64, 0x3a, 0x6c, 0xc7, 0x6b, 0x3d, 0xf8, 0x10,
	0xef, 0xcf, 0x52, 0xef, 0xec, 0x2d, 0xd8, 0x59, 0x67, 0xcc, 0xd0, 0x8a, 0x66, 0xce, 0x09, 0xde,
	0x13, 0x12, 0x54, 0x66, 0x78, 0x30, 0xad, 0xb2, 0x59, 0xc9, 0xb7, 0xfe, 0x0a, 0xde, 0x1b, 0xc2,
	0xc7, 0x8b, 0xd9, 0xd7, 0xcb, 0x8c, 0x9d, 0xfd, 0xda, 0xcb, 0xd4, 0x2f, 0x3f, 0xc7, 0x04, 0x8d,
	0xc6, 0x04, 0xfd, 0x8c, 0x09, 0x7a, 0x9f, 0x90, 0xdc, 0x68, 0x42, 0x72, 0xdf, 0x13, 0x92, 0xbb,
	0xaf, 0x46, 0xcc, 0x74, 0x9f, 0xdb, 0x7e, 0x47, 0xf4, 0x02, 0x2d, 0x21, 0x8e, 0x75, 0x37, 0x94,
	0x10, 0x70, 0x30, 0x03, 0xa1, 0x9e, 0x82, 0x97, 0xc0, 0x2e, 0xb2, 0x19, 0x4a, 0xd0, 0xed, 0xed,
	0xd9, 0x16, 0x9f, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x29, 0x27, 0xec, 0xef, 0x5a, 0x03, 0x00,
	0x00,
}

func (m *EventCoordinatorCreated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCoordinatorCreated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCoordinatorCreated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.CoordinatorID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.CoordinatorID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventCoordinatorAddressUpdated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCoordinatorAddressUpdated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCoordinatorAddressUpdated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewAddress) > 0 {
		i -= len(m.NewAddress)
		copy(dAtA[i:], m.NewAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NewAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.CoordinatorID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.CoordinatorID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventCoordinatorDisabled) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCoordinatorDisabled) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCoordinatorDisabled) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.CoordinatorID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.CoordinatorID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventValidatorCreated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventValidatorCreated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventValidatorCreated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OperatorAddresses) > 0 {
		for iNdEx := len(m.OperatorAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.OperatorAddresses[iNdEx])
			copy(dAtA[i:], m.OperatorAddresses[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.OperatorAddresses[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventValidatorOperatorAddressesUpdated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventValidatorOperatorAddressesUpdated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventValidatorOperatorAddressesUpdated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OperatorAddresses) > 0 {
		for iNdEx := len(m.OperatorAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.OperatorAddresses[iNdEx])
			copy(dAtA[i:], m.OperatorAddresses[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.OperatorAddresses[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
func (m *EventCoordinatorCreated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CoordinatorID != 0 {
		n += 1 + sovEvents(uint64(m.CoordinatorID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCoordinatorAddressUpdated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CoordinatorID != 0 {
		n += 1 + sovEvents(uint64(m.CoordinatorID))
	}
	l = len(m.NewAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventCoordinatorDisabled) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CoordinatorID != 0 {
		n += 1 + sovEvents(uint64(m.CoordinatorID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventValidatorCreated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.OperatorAddresses) > 0 {
		for _, s := range m.OperatorAddresses {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventValidatorOperatorAddressesUpdated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.OperatorAddresses) > 0 {
		for _, s := range m.OperatorAddresses {
			l = len(s)
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
func (m *EventCoordinatorCreated) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventCoordinatorCreated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCoordinatorCreated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorID", wireType)
			}
			m.CoordinatorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoordinatorID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *EventCoordinatorAddressUpdated) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventCoordinatorAddressUpdated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCoordinatorAddressUpdated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorID", wireType)
			}
			m.CoordinatorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoordinatorID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewAddress", wireType)
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
			m.NewAddress = string(dAtA[iNdEx:postIndex])
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
func (m *EventCoordinatorDisabled) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventCoordinatorDisabled: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCoordinatorDisabled: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorID", wireType)
			}
			m.CoordinatorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoordinatorID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *EventValidatorCreated) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventValidatorCreated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventValidatorCreated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddresses", wireType)
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
			m.OperatorAddresses = append(m.OperatorAddresses, string(dAtA[iNdEx:postIndex]))
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
func (m *EventValidatorOperatorAddressesUpdated) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EventValidatorOperatorAddressesUpdated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventValidatorOperatorAddressesUpdated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddresses", wireType)
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
			m.OperatorAddresses = append(m.OperatorAddresses, string(dAtA[iNdEx:postIndex]))
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
