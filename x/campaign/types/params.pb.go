// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: campaign/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// Params defines the set of params for the campaign module.
type Params struct {
	TotalSupplyRange    TotalSupplyRange                         `protobuf:"bytes,1,opt,name=totalSupplyRange,proto3" json:"totalSupplyRange"`
	CampaignCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=campaignCreationFee,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"campaignCreationFee"`
	MaxMetadataLength   uint64                                   `protobuf:"varint,3,opt,name=maxMetadataLength,proto3" json:"maxMetadataLength,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f21b288c6be0f59, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetTotalSupplyRange() TotalSupplyRange {
	if m != nil {
		return m.TotalSupplyRange
	}
	return TotalSupplyRange{}
}

func (m *Params) GetCampaignCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.CampaignCreationFee
	}
	return nil
}

func (m *Params) GetMaxMetadataLength() uint64 {
	if m != nil {
		return m.MaxMetadataLength
	}
	return 0
}

// TotalSupplyRange defines the range of allowed values for total supply
type TotalSupplyRange struct {
	MinTotalSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=minTotalSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"minTotalSupply"`
	MaxTotalSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=maxTotalSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"maxTotalSupply"`
}

func (m *TotalSupplyRange) Reset()         { *m = TotalSupplyRange{} }
func (m *TotalSupplyRange) String() string { return proto.CompactTextString(m) }
func (*TotalSupplyRange) ProtoMessage()    {}
func (*TotalSupplyRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f21b288c6be0f59, []int{1}
}
func (m *TotalSupplyRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TotalSupplyRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TotalSupplyRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TotalSupplyRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TotalSupplyRange.Merge(m, src)
}
func (m *TotalSupplyRange) XXX_Size() int {
	return m.Size()
}
func (m *TotalSupplyRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TotalSupplyRange.DiscardUnknown(m)
}

var xxx_messageInfo_TotalSupplyRange proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "tendermint.spn.campaign.Params")
	proto.RegisterType((*TotalSupplyRange)(nil), "tendermint.spn.campaign.TotalSupplyRange")
}

func init() { proto.RegisterFile("campaign/params.proto", fileDescriptor_6f21b288c6be0f59) }

var fileDescriptor_6f21b288c6be0f59 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x8b, 0xd3, 0x40,
	0x18, 0xce, 0xb4, 0xa5, 0xe0, 0x14, 0xa4, 0x46, 0xc5, 0xb6, 0x87, 0xa4, 0xf4, 0xa0, 0x11, 0xec,
	0x0c, 0xad, 0x37, 0xf1, 0x94, 0x82, 0x58, 0x50, 0x90, 0xe8, 0xc9, 0x1e, 0x64, 0x92, 0x0c, 0x69,
	0x68, 0x32, 0x33, 0x64, 0xa6, 0xda, 0xfe, 0x0b, 0x8f, 0x1e, 0x3d, 0x8a, 0x67, 0x7f, 0x82, 0x87,
	0x1e, 0x8b, 0x27, 0xf1, 0x10, 0x97, 0xf6, 0x5f, 0xec, 0x69, 0xc9, 0x47, 0x97, 0x6e, 0xbb, 0x0b,
	0x3d, 0xec, 0x29, 0x1f, 0xef, 0xf3, 0xf9, 0xf2, 0xc2, 0x87, 0x1e, 0x89, 0x05, 0x09, 0x03, 0x86,
	0x05, 0x49, 0x48, 0x2c, 0x91, 0x48, 0xb8, 0xe2, 0xfa, 0x23, 0x45, 0x99, 0x4f, 0x93, 0x38, 0x64,
	0x0a, 0x49, 0xc1, 0xd0, 0x0e, 0xd5, 0x79, 0x10, 0xf0, 0x80, 0xe7, 0x18, 0x9c, 0xbd, 0x15, 0xf0,
	0x8e, 0xe1, 0x71, 0x19, 0x73, 0x89, 0x5d, 0x22, 0x29, 0xfe, 0x3c, 0x70, 0xa9, 0x22, 0x03, 0xec,
	0xf1, 0x90, 0x95, 0xf3, 0x76, 0x31, 0xff, 0x54, 0x10, 0x8b, 0x8f, 0x62, 0xd4, 0xfb, 0x5d, 0x81,
	0xf5, 0x77, 0xb9, 0xb5, 0x3e, 0x81, 0x4d, 0xc5, 0x15, 0x89, 0xde, 0xcf, 0x85, 0x88, 0x96, 0x0e,
	0x61, 0x01, 0x6d, 0x81, 0x2e, 0xb0, 0x1a, 0xc3, 0xa7, 0xe8, 0x86, 0x3c, 0xe8, 0xc3, 0x01, 0xc1,
	0xae, 0xad, 0x52, 0x53, 0x73, 0x8e, 0x84, 0xf4, 0x1f, 0x00, 0xde, 0xdf, 0xb1, 0x46, 0x09, 0x25,
	0x2a, 0xe4, 0xec, 0x15, 0xa5, 0xad, 0x4a, 0xb7, 0x6a, 0x35, 0x86, 0x6d, 0x54, 0x86, 0xca, 0x1a,
	0xa0, 0xb2, 0x01, 0x1a, 0xf1, 0x90, 0xd9, 0x93, 0x4c, 0xf0, 0x3c, 0x35, 0x9f, 0x04, 0xa1, 0x9a,
	0xce, 0x5d, 0xe4, 0xf1, 0xb8, 0x6c, 0x50, 0x3e, 0xfa, 0xd2, 0x9f, 0x61, 0xb5, 0x14, 0x54, 0xe6,
	0x84, 0x9f, 0xff, 0x4d, 0xeb, 0x44, 0xa8, 0x74, 0xae, 0x8b, 0xa4, 0x3f, 0x83, 0xf7, 0x62, 0xb2,
	0x78, 0x4b, 0x15, 0xf1, 0x89, 0x22, 0x6f, 0x28, 0x0b, 0xd4, 0xb4, 0x55, 0xed, 0x02, 0xab, 0xe6,
	0x1c, 0x0f, 0x5e, 0xd4, 0xbe, 0x7d, 0x37, 0xb5, 0x5e, 0x0a, 0x60, 0xf3, 0x70, 0x17, 0xba, 0x0f,
	0xef, 0xc6, 0x21, 0xdb, 0xfb, 0x9d, 0xaf, 0xf3, 0x8e, 0xfd, 0x32, 0xab, 0xf4, 0x2f, 0x35, 0x1f,
	0x9f, 0x90, 0x73, 0xcc, 0xd4, 0x9f, 0x5f, 0x7d, 0x58, 0xae, 0x67, 0xcc, 0x94, 0x73, 0xa0, 0x99,
	0xbb, 0x90, 0xc5, 0xbe, 0x4b, 0xe5, 0x56, 0x5c, 0xae, 0x68, 0xda, 0xaf, 0x57, 0x1b, 0x03, 0xac,
	0x37, 0x06, 0x38, 0xdb, 0x18, 0xe0, 0xeb, 0xd6, 0xd0, 0xd6, 0x5b, 0x43, 0xfb, 0xbb, 0x35, 0xb4,
	0x8f, 0x68, 0x4f, 0x5f, 0x0a, 0x1a, 0x45, 0x72, 0x4a, 0x04, 0xc5, 0x8c, 0xaa, 0x2f, 0x3c, 0x99,
	0xe1, 0x05, 0xbe, 0x3c, 0xf1, 0xdc, 0xcb, 0xad, 0xe7, 0x87, 0xf7, 0xfc, 0x22, 0x00, 0x00, 0xff,
	0xff, 0x82, 0xa6, 0x58, 0xb5, 0xfb, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxMetadataLength != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxMetadataLength))
		i--
		dAtA[i] = 0x18
	}
	if len(m.CampaignCreationFee) > 0 {
		for iNdEx := len(m.CampaignCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CampaignCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.TotalSupplyRange.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *TotalSupplyRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TotalSupplyRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TotalSupplyRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxTotalSupply.Size()
		i -= size
		if _, err := m.MaxTotalSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinTotalSupply.Size()
		i -= size
		if _, err := m.MinTotalSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TotalSupplyRange.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.CampaignCreationFee) > 0 {
		for _, e := range m.CampaignCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.MaxMetadataLength != 0 {
		n += 1 + sovParams(uint64(m.MaxMetadataLength))
	}
	return n
}

func (m *TotalSupplyRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinTotalSupply.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxTotalSupply.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupplyRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalSupplyRange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CampaignCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CampaignCreationFee = append(m.CampaignCreationFee, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.CampaignCreationFee[len(m.CampaignCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMetadataLength", wireType)
			}
			m.MaxMetadataLength = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxMetadataLength |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *TotalSupplyRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: TotalSupplyRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TotalSupplyRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTotalSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinTotalSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTotalSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxTotalSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
