// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/claim/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the claim module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// list of claim records with the corresponding airdrop recipient
	ClaimRecords []ClaimRecordAddress `protobuf:"bytes,2,rep,name=claim_records,json=claimRecords,proto3" json:"claim_records"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_45c5d451b1a38941, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetClaimRecords() []ClaimRecordAddress {
	if m != nil {
		return m.ClaimRecords
	}
	return nil
}

// Params defines the claim module's parameters.
type Params struct {
	// enable claiming process
	EnableClaim bool `protobuf:"varint,1,opt,name=enable_claim,json=enableClaim,proto3" json:"enable_claim,omitempty"`
	// timestamp of the airdrop start
	AirdropStartTime time.Time `protobuf:"bytes,2,opt,name=airdrop_start_time,json=airdropStartTime,proto3,stdtime" json:"airdrop_start_time"`
	// duration until decay of claimable tokens begin
	DurationUntilDecay time.Duration `protobuf:"bytes,3,opt,name=duration_until_decay,json=durationUntilDecay,proto3,stdduration" json:"duration_until_decay"`
	// duration of the token claim decay period
	DurationOfDecay time.Duration `protobuf:"bytes,4,opt,name=duration_of_decay,json=durationOfDecay,proto3,stdduration" json:"duration_of_decay"`
	// denom of claimable coin
	ClaimDenom string `protobuf:"bytes,5,opt,name=claim_denom,json=claimDenom,proto3" json:"claim_denom,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_45c5d451b1a38941, []int{1}
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

func (m *Params) GetEnableClaim() bool {
	if m != nil {
		return m.EnableClaim
	}
	return false
}

func (m *Params) GetAirdropStartTime() time.Time {
	if m != nil {
		return m.AirdropStartTime
	}
	return time.Time{}
}

func (m *Params) GetDurationUntilDecay() time.Duration {
	if m != nil {
		return m.DurationUntilDecay
	}
	return 0
}

func (m *Params) GetDurationOfDecay() time.Duration {
	if m != nil {
		return m.DurationOfDecay
	}
	return 0
}

func (m *Params) GetClaimDenom() string {
	if m != nil {
		return m.ClaimDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "evmos.claim.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "evmos.claim.v1.Params")
}

func init() { proto.RegisterFile("evmos/claim/v1/genesis.proto", fileDescriptor_45c5d451b1a38941) }

var fileDescriptor_45c5d451b1a38941 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0xeb, 0x6e, 0x54, 0xc3, 0x2d, 0xff, 0xac, 0x09, 0x85, 0x08, 0xb9, 0x5d, 0x0f, 0xa8,
	0x27, 0x5b, 0x1d, 0x7c, 0x00, 0x28, 0x95, 0x38, 0xa1, 0xa1, 0x8c, 0x5d, 0xb8, 0x44, 0x4e, 0xe2,
	0x66, 0x16, 0x4d, 0x1c, 0xd9, 0x6e, 0xc4, 0xbe, 0xc5, 0x24, 0x2e, 0x7c, 0xa4, 0x1e, 0x77, 0xe4,
	0x04, 0xa8, 0xfd, 0x22, 0xc8, 0x76, 0x52, 0xd4, 0x72, 0xd9, 0x2d, 0xf1, 0xf3, 0x3c, 0xbf, 0xf7,
	0xf5, 0x23, 0xc3, 0x97, 0xbc, 0x2e, 0xa4, 0xa6, 0xe9, 0x92, 0x89, 0x82, 0xd6, 0x53, 0x9a, 0xf3,
	0x92, 0x6b, 0xa1, 0x49, 0xa5, 0xa4, 0x91, 0xe8, 0xb1, 0x53, 0x89, 0x53, 0x49, 0x3d, 0x0d, 0x4f,
	0x73, 0x99, 0x4b, 0x27, 0x51, 0xfb, 0xe5, 0x5d, 0x21, 0x4e, 0xa5, 0xb6, 0x90, 0x84, 0x69, 0x4e,
	0xeb, 0x69, 0xc2, 0x0d, 0x9b, 0xd2, 0x54, 0x8a, 0xb2, 0xd1, 0xcf, 0x76, 0x7a, 0xf9, 0x75, 0xa7,
	0xef, 0x0d, 0x0a, 0x71, 0x2e, 0x65, 0xbe, 0xe4, 0xd4, 0xfd, 0x25, 0xab, 0x05, 0xcd, 0x56, 0x8a,
	0x19, 0x21, 0x5b, 0xc4, 0xf0, 0x50, 0x37, 0xa2, 0xe0, 0xda, 0xb0, 0xa2, 0x6a, 0x0c, 0xe1, 0xc1,
	0x3d, 0xfc, 0xca, 0x4e, 0x1b, 0x7f, 0x07, 0x70, 0xf0, 0xc1, 0x8f, 0xbb, 0x34, 0xcc, 0x70, 0xf4,
	0x06, 0xf6, 0x2a, 0xa6, 0x58, 0xa1, 0x03, 0x30, 0x02, 0x93, 0xfe, 0xf9, 0x73, 0xb2, 0x7f, 0x4f,
	0xf2, 0xc9, 0xa9, 0xb3, 0xe3, 0xf5, 0xaf, 0x61, 0x27, 0x6a, 0xbc, 0xe8, 0x23, 0x7c, 0xe4, 0x0c,
	0xb1, 0xe2, 0xa9, 0x54, 0x99, 0x0e, 0xba, 0xa3, 0xa3, 0x49, 0xff, 0x7c, 0x7c, 0x18, 0x7e, 0x6f,
	0x3f, 0x22, 0xe7, 0x79, 0x97, 0x65, 0x8a, 0xeb, 0x16, 0x34, 0x48, 0xff, 0x29, 0x7a, 0xbc, 0xee,
	0xc2, 0x9e, 0x9f, 0x83, 0xce, 0xe0, 0x80, 0x97, 0x2c, 0x59, 0xf2, 0xd8, 0x39, 0xdc, 0x56, 0x27,
	0x51, 0xdf, 0x9f, 0x39, 0x1c, 0x8a, 0x20, 0x62, 0x42, 0x65, 0x4a, 0x56, 0xb1, 0x36, 0x4c, 0x99,
	0xd8, 0x16, 0x10, 0x74, 0xdd, 0xfa, 0x21, 0xf1, 0xed, 0x90, 0xb6, 0x1d, 0xf2, 0xb9, 0x6d, 0x67,
	0x76, 0x62, 0x27, 0xdf, 0xfe, 0x1e, 0x82, 0xe8, 0x69, 0x93, 0xbf, 0xb4, 0x71, 0x6b, 0x40, 0x57,
	0xf0, 0xb4, 0xad, 0x39, 0x5e, 0x95, 0x46, 0x2c, 0xe3, 0x8c, 0xa7, 0xec, 0x26, 0x38, 0x72, 0xd4,
	0x17, 0xff, 0x51, 0xe7, 0x8d, 0xd9, 0x43, 0x7f, 0x58, 0x28, 0x6a, 0x01, 0x57, 0x36, 0x3f, 0xb7,
	0x71, 0x74, 0x01, 0x9f, 0xed, 0xb0, 0x72, 0xd1, 0x30, 0x8f, 0xef, 0xcf, 0x7c, 0xd2, 0xa6, 0x2f,
	0x16, 0x1e, 0x38, 0x84, 0x7d, 0x5f, 0x7c, 0xc6, 0x4b, 0x59, 0x04, 0x0f, 0x46, 0x60, 0xf2, 0x30,
	0x82, 0xee, 0x68, 0x6e, 0x4f, 0x66, 0x6f, 0xd7, 0x1b, 0x0c, 0xee, 0x36, 0x18, 0xfc, 0xd9, 0x60,
	0x70, 0xbb, 0xc5, 0x9d, 0xbb, 0x2d, 0xee, 0xfc, 0xdc, 0xe2, 0xce, 0x97, 0x57, 0xb9, 0x30, 0xd7,
	0xab, 0x84, 0xa4, 0xb2, 0xa0, 0xe6, 0x9a, 0x29, 0x2d, 0x34, 0xf5, 0x2f, 0xe5, 0x5b, 0xf3, 0x56,
	0xcc, 0x4d, 0xc5, 0x75, 0xd2, 0x73, 0x0b, 0xbd, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x55, 0x27,
	0x96, 0x36, 0x0f, 0x03, 0x00, 0x00,
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
	if len(m.ClaimRecords) > 0 {
		for iNdEx := len(m.ClaimRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ClaimRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
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
	if len(m.ClaimDenom) > 0 {
		i -= len(m.ClaimDenom)
		copy(dAtA[i:], m.ClaimDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ClaimDenom)))
		i--
		dAtA[i] = 0x2a
	}
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationOfDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationUntilDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGenesis(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.AirdropStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintGenesis(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x12
	if m.EnableClaim {
		i--
		if m.EnableClaim {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.ClaimRecords) > 0 {
		for _, e := range m.ClaimRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableClaim {
		n += 2
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStartTime)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay)
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.ClaimDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimRecords", wireType)
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
			m.ClaimRecords = append(m.ClaimRecords, ClaimRecordAddress{})
			if err := m.ClaimRecords[len(m.ClaimRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableClaim", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.EnableClaim = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropStartTime", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.AirdropStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationUntilDecay", wireType)
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationUntilDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationOfDecay", wireType)
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationOfDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimDenom", wireType)
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
			m.ClaimDenom = string(dAtA[iNdEx:postIndex])
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