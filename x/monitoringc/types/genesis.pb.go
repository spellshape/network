// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: monitoringc/genesis.proto

package types

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

// GenesisState defines the monitoringc module's genesis state.
type GenesisState struct {
	PortId                        string                         `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	VerifiedClientIDs             []VerifiedClientID             `protobuf:"bytes,2,rep,name=verifiedClientIDs,proto3" json:"verifiedClientIDs"`
	ProviderClientIDs             []ProviderClientID             `protobuf:"bytes,3,rep,name=providerClientIDs,proto3" json:"providerClientIDs"`
	LaunchIDsFromVerifiedClientID []LaunchIDFromVerifiedClientID `protobuf:"bytes,4,rep,name=launchIDsFromVerifiedClientID,proto3" json:"launchIDsFromVerifiedClientID"`
	LaunchIDsFromChannelID        []LaunchIDFromChannelID        `protobuf:"bytes,5,rep,name=launchIDsFromChannelID,proto3" json:"launchIDsFromChannelID"`
	MonitoringHistories           []MonitoringHistory            `protobuf:"bytes,6,rep,name=monitoringHistories,proto3" json:"monitoringHistories"`
	Params                        Params                         `protobuf:"bytes,7,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_10c269cb89bd5f03, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetVerifiedClientIDs() []VerifiedClientID {
	if m != nil {
		return m.VerifiedClientIDs
	}
	return nil
}

func (m *GenesisState) GetProviderClientIDs() []ProviderClientID {
	if m != nil {
		return m.ProviderClientIDs
	}
	return nil
}

func (m *GenesisState) GetLaunchIDsFromVerifiedClientID() []LaunchIDFromVerifiedClientID {
	if m != nil {
		return m.LaunchIDsFromVerifiedClientID
	}
	return nil
}

func (m *GenesisState) GetLaunchIDsFromChannelID() []LaunchIDFromChannelID {
	if m != nil {
		return m.LaunchIDsFromChannelID
	}
	return nil
}

func (m *GenesisState) GetMonitoringHistories() []MonitoringHistory {
	if m != nil {
		return m.MonitoringHistories
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tendermint.spn.monitoringc.GenesisState")
}

func init() { proto.RegisterFile("monitoringc/genesis.proto", fileDescriptor_10c269cb89bd5f03) }

var fileDescriptor_10c269cb89bd5f03 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x8f, 0x93, 0x40,
	0x18, 0x86, 0xc1, 0xdd, 0x65, 0xe3, 0xac, 0x17, 0xd1, 0x28, 0x92, 0x88, 0xcd, 0xc6, 0x43, 0x13,
	0x15, 0xb2, 0xbb, 0x17, 0x8f, 0xa6, 0x6d, 0x54, 0xa2, 0x26, 0xa6, 0x26, 0x1e, 0xbc, 0x20, 0x85,
	0xaf, 0x30, 0x11, 0x66, 0x26, 0x33, 0xd3, 0x6a, 0xef, 0x1e, 0x3d, 0xf8, 0xb3, 0x7a, 0xec, 0xd1,
	0x93, 0x31, 0xed, 0x1f, 0x31, 0x0c, 0xd4, 0x22, 0x2d, 0xe8, 0x6d, 0x9a, 0xbe, 0xef, 0xf3, 0x7c,
	0xf9, 0x66, 0x40, 0xf7, 0x72, 0x4a, 0xb0, 0xa4, 0x1c, 0x93, 0x24, 0xf2, 0x12, 0x20, 0x20, 0xb0,
	0x70, 0x19, 0xa7, 0x92, 0x9a, 0xb6, 0x04, 0x12, 0x03, 0xcf, 0x31, 0x91, 0xae, 0x60, 0xc4, 0xad,
	0x25, 0xed, 0xdb, 0x09, 0x4d, 0xa8, 0x8a, 0x79, 0xc5, 0xa9, 0x6c, 0xd8, 0x56, 0x1d, 0xc6, 0x42,
	0x1e, 0xe6, 0x15, 0xcb, 0x7e, 0x58, 0xff, 0x67, 0x0e, 0x1c, 0x4f, 0x31, 0xc4, 0x41, 0x94, 0x61,
	0x20, 0x32, 0xc0, 0xf1, 0xa1, 0x14, 0xe3, 0x74, 0x8e, 0x63, 0xe0, 0x7b, 0xa9, 0xab, 0x7a, 0x2a,
	0x0b, 0x67, 0x24, 0x4a, 0x03, 0x1c, 0x07, 0x53, 0x4e, 0xf3, 0xa0, 0x15, 0xfd, 0xa8, 0xa3, 0x14,
	0xa5, 0x21, 0x21, 0x90, 0xb5, 0xcc, 0xb1, 0x3b, 0x07, 0x29, 0x16, 0x92, 0xf2, 0x45, 0x99, 0x3a,
	0xff, 0x76, 0x82, 0x6e, 0xbc, 0x28, 0x37, 0xf6, 0x4e, 0x86, 0x12, 0xcc, 0xbb, 0xe8, 0x94, 0x51,
	0x5e, 0x48, 0x2d, 0xbd, 0xa7, 0xf7, 0xaf, 0x8f, 0x8d, 0xe2, 0xa7, 0x1f, 0x9b, 0x1f, 0xd1, 0xcd,
	0xed, 0x60, 0x43, 0x35, 0x97, 0x3f, 0x12, 0xd6, 0xb5, 0xde, 0x51, 0xff, 0xec, 0xf2, 0xb1, 0xdb,
	0xbe, 0x65, 0xf7, 0x7d, 0xa3, 0x34, 0x38, 0x5e, 0xfe, 0x7c, 0xa0, 0x8d, 0xf7, 0x61, 0x85, 0x61,
	0xbb, 0xaf, 0x9d, 0xe1, 0xe8, 0xdf, 0x86, 0xb7, 0x8d, 0xd2, 0xd6, 0xb0, 0x07, 0x33, 0xbf, 0xea,
	0xe8, 0x7e, 0xb9, 0x37, 0x7f, 0x24, 0x9e, 0x73, 0x9a, 0x37, 0x87, 0xb3, 0x8e, 0x95, 0xee, 0x69,
	0x97, 0xee, 0x75, 0x05, 0x38, 0xd4, 0xaf, 0xd4, 0xdd, 0x12, 0x93, 0xa2, 0x3b, 0x7f, 0x05, 0x86,
	0xe5, 0xdd, 0xf9, 0x23, 0xeb, 0x44, 0xe9, 0x2f, 0xfe, 0x57, 0xff, 0xa7, 0x58, 0x79, 0x5b, 0xb0,
	0x26, 0xa0, 0x5b, 0x3b, 0xc4, 0x4b, 0xf5, 0x00, 0x30, 0x08, 0xcb, 0x50, 0xb6, 0x27, 0x5d, 0xb6,
	0x37, 0x8d, 0xda, 0xa2, 0x32, 0x1d, 0xe2, 0x99, 0xcf, 0x90, 0x51, 0x7e, 0x30, 0xd6, 0x69, 0x4f,
	0xef, 0x9f, 0x5d, 0x9e, 0x77, 0xde, 0x9a, 0x4a, 0x56, 0xb8, 0xaa, 0x37, 0x78, 0xb5, 0x5c, 0x3b,
	0xfa, 0x6a, 0xed, 0xe8, 0xbf, 0xd6, 0x8e, 0xfe, 0x7d, 0xe3, 0x68, 0xab, 0x8d, 0xa3, 0xfd, 0xd8,
	0x38, 0xda, 0x87, 0x8b, 0x04, 0xcb, 0x74, 0x36, 0x71, 0x23, 0x9a, 0x7b, 0x82, 0x41, 0x96, 0x89,
	0x34, 0x64, 0xe0, 0x11, 0x90, 0x9f, 0x29, 0xff, 0xe4, 0x7d, 0xf1, 0xea, 0xcf, 0x5d, 0x2e, 0x18,
	0x88, 0x89, 0xa1, 0x9e, 0xf8, 0xd5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x3a, 0xe0, 0xfa,
	0x1f, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.MonitoringHistories) > 0 {
		for iNdEx := len(m.MonitoringHistories) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MonitoringHistories[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.LaunchIDsFromChannelID) > 0 {
		for iNdEx := len(m.LaunchIDsFromChannelID) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LaunchIDsFromChannelID[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.LaunchIDsFromVerifiedClientID) > 0 {
		for iNdEx := len(m.LaunchIDsFromVerifiedClientID) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LaunchIDsFromVerifiedClientID[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ProviderClientIDs) > 0 {
		for iNdEx := len(m.ProviderClientIDs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProviderClientIDs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.VerifiedClientIDs) > 0 {
		for iNdEx := len(m.VerifiedClientIDs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VerifiedClientIDs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.VerifiedClientIDs) > 0 {
		for _, e := range m.VerifiedClientIDs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProviderClientIDs) > 0 {
		for _, e := range m.ProviderClientIDs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LaunchIDsFromVerifiedClientID) > 0 {
		for _, e := range m.LaunchIDsFromVerifiedClientID {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LaunchIDsFromChannelID) > 0 {
		for _, e := range m.LaunchIDsFromChannelID {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MonitoringHistories) > 0 {
		for _, e := range m.MonitoringHistories {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifiedClientIDs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerifiedClientIDs = append(m.VerifiedClientIDs, VerifiedClientID{})
			if err := m.VerifiedClientIDs[len(m.VerifiedClientIDs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProviderClientIDs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProviderClientIDs = append(m.ProviderClientIDs, ProviderClientID{})
			if err := m.ProviderClientIDs[len(m.ProviderClientIDs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchIDsFromVerifiedClientID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LaunchIDsFromVerifiedClientID = append(m.LaunchIDsFromVerifiedClientID, LaunchIDFromVerifiedClientID{})
			if err := m.LaunchIDsFromVerifiedClientID[len(m.LaunchIDsFromVerifiedClientID)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchIDsFromChannelID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LaunchIDsFromChannelID = append(m.LaunchIDsFromChannelID, LaunchIDFromChannelID{})
			if err := m.LaunchIDsFromChannelID[len(m.LaunchIDsFromChannelID)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MonitoringHistories", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MonitoringHistories = append(m.MonitoringHistories, MonitoringHistory{})
			if err := m.MonitoringHistories[len(m.MonitoringHistories)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
