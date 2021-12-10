// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relayers/relayer.proto

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

// ProposalStatus enumerates the valid statuses of a proposal.
type ProposalStatus int32

const (
	StatusInitiated ProposalStatus = 0
	StatusApproved  ProposalStatus = 1
	StatusRejected  ProposalStatus = 2
	StatusExpired   ProposalStatus = 3
)

var ProposalStatus_name = map[int32]string{
	0: "PROPOSAL_STATUS_INITIATED",
	1: "PROPOSAL_STATUS_APPROVED",
	2: "PROPOSAL_STATUS_REJECTED",
	3: "PROPOSAL_STATUS_EXPIRED",
}

var ProposalStatus_value = map[string]int32{
	"PROPOSAL_STATUS_INITIATED": 0,
	"PROPOSAL_STATUS_APPROVED":  1,
	"PROPOSAL_STATUS_REJECTED":  2,
	"PROPOSAL_STATUS_EXPIRED":   3,
}

func (x ProposalStatus) String() string {
	return proto.EnumName(ProposalStatus_name, int32(x))
}

func (ProposalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_40adf04c18425525, []int{0}
}

type Relayer struct {
	Denom   string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *Relayer) Reset()         { *m = Relayer{} }
func (m *Relayer) String() string { return proto.CompactTextString(m) }
func (*Relayer) ProtoMessage()    {}
func (*Relayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_40adf04c18425525, []int{0}
}
func (m *Relayer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Relayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Relayer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Relayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relayer.Merge(m, src)
}
func (m *Relayer) XXX_Size() int {
	return m.Size()
}
func (m *Relayer) XXX_DiscardUnknown() {
	xxx_messageInfo_Relayer.DiscardUnknown(m)
}

var xxx_messageInfo_Relayer proto.InternalMessageInfo

func (m *Relayer) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Relayer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Threshold struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Value uint32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Threshold) Reset()         { *m = Threshold{} }
func (m *Threshold) String() string { return proto.CompactTextString(m) }
func (*Threshold) ProtoMessage()    {}
func (*Threshold) Descriptor() ([]byte, []int) {
	return fileDescriptor_40adf04c18425525, []int{1}
}
func (m *Threshold) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Threshold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Threshold.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Threshold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Threshold.Merge(m, src)
}
func (m *Threshold) XXX_Size() int {
	return m.Size()
}
func (m *Threshold) XXX_DiscardUnknown() {
	xxx_messageInfo_Threshold.DiscardUnknown(m)
}

var xxx_messageInfo_Threshold proto.InternalMessageInfo

func (m *Threshold) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Threshold) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Proposal struct {
	ProposalId   []byte           `protobuf:"bytes,1,opt,name=proposalId,proto3" json:"proposalId,omitempty"`
	Content      *ProposalContent `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Status       ProposalStatus   `protobuf:"varint,3,opt,name=status,proto3,enum=stafiprotocol.stafihub.relayers.ProposalStatus" json:"status,omitempty"`
	VotesFor     []string         `protobuf:"bytes,4,rep,name=votesFor,proto3" json:"votesFor,omitempty"`
	VotesAgainst []string         `protobuf:"bytes,5,rep,name=votesAgainst,proto3" json:"votesAgainst,omitempty"`
	StartBlock   int64            `protobuf:"varint,6,opt,name=startBlock,proto3" json:"startBlock,omitempty"`
	ExpireBlock  int64            `protobuf:"varint,7,opt,name=expireBlock,proto3" json:"expireBlock,omitempty"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_40adf04c18425525, []int{2}
}
func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return m.Size()
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetProposalId() []byte {
	if m != nil {
		return m.ProposalId
	}
	return nil
}

func (m *Proposal) GetContent() *ProposalContent {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Proposal) GetStatus() ProposalStatus {
	if m != nil {
		return m.Status
	}
	return StatusInitiated
}

func (m *Proposal) GetVotesFor() []string {
	if m != nil {
		return m.VotesFor
	}
	return nil
}

func (m *Proposal) GetVotesAgainst() []string {
	if m != nil {
		return m.VotesAgainst
	}
	return nil
}

func (m *Proposal) GetStartBlock() int64 {
	if m != nil {
		return m.StartBlock
	}
	return 0
}

func (m *Proposal) GetExpireBlock() int64 {
	if m != nil {
		return m.ExpireBlock
	}
	return 0
}

type ProposalContent struct {
	Denom         string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	ProposalRoute string `protobuf:"bytes,2,opt,name=proposalRoute,proto3" json:"proposalRoute,omitempty"`
	Name          string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Params        []byte `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *ProposalContent) Reset()         { *m = ProposalContent{} }
func (m *ProposalContent) String() string { return proto.CompactTextString(m) }
func (*ProposalContent) ProtoMessage()    {}
func (*ProposalContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_40adf04c18425525, []int{3}
}
func (m *ProposalContent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalContent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalContent.Merge(m, src)
}
func (m *ProposalContent) XXX_Size() int {
	return m.Size()
}
func (m *ProposalContent) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalContent.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalContent proto.InternalMessageInfo

func (m *ProposalContent) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *ProposalContent) GetProposalRoute() string {
	if m != nil {
		return m.ProposalRoute
	}
	return ""
}

func (m *ProposalContent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProposalContent) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterEnum("stafiprotocol.stafihub.relayers.ProposalStatus", ProposalStatus_name, ProposalStatus_value)
	proto.RegisterType((*Relayer)(nil), "stafiprotocol.stafihub.relayers.Relayer")
	proto.RegisterType((*Threshold)(nil), "stafiprotocol.stafihub.relayers.Threshold")
	proto.RegisterType((*Proposal)(nil), "stafiprotocol.stafihub.relayers.Proposal")
	proto.RegisterType((*ProposalContent)(nil), "stafiprotocol.stafihub.relayers.ProposalContent")
}

func init() { proto.RegisterFile("relayers/relayer.proto", fileDescriptor_40adf04c18425525) }

var fileDescriptor_40adf04c18425525 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x8f, 0xd2, 0x40,
	0x14, 0xc7, 0xe9, 0xc2, 0xc2, 0xf2, 0x76, 0x61, 0x71, 0x24, 0x6b, 0xed, 0xa1, 0x36, 0xc4, 0x03,
	0xf1, 0xd0, 0x6e, 0xf0, 0x60, 0x3c, 0x76, 0x97, 0x6a, 0xba, 0xd9, 0x48, 0x33, 0xa0, 0x31, 0x5e,
	0x36, 0xb3, 0x74, 0x84, 0x6a, 0xe9, 0x34, 0x9d, 0x81, 0x2c, 0xdf, 0xc0, 0x70, 0xf2, 0x0b, 0x70,
	0xf2, 0x93, 0x78, 0xf3, 0xb8, 0x47, 0x4f, 0xc6, 0xc0, 0x17, 0x31, 0x4c, 0xe9, 0x06, 0x88, 0x1b,
	0xbd, 0xbd, 0xff, 0xbf, 0xff, 0xdf, 0x9b, 0xcc, 0x9b, 0x57, 0x38, 0x49, 0x68, 0x48, 0xa6, 0x34,
	0xe1, 0xd6, 0xba, 0x30, 0xe3, 0x84, 0x09, 0x86, 0x9e, 0x70, 0x41, 0x3e, 0x06, 0xb2, 0xee, 0xb3,
	0xd0, 0x94, 0x6a, 0x38, 0xbe, 0x36, 0xb3, 0xb8, 0x56, 0x1f, 0xb0, 0x01, 0x93, 0xdf, 0xad, 0x55,
	0x95, 0x62, 0x8d, 0x97, 0x50, 0xc2, 0x69, 0x02, 0xd5, 0x61, 0xdf, 0xa7, 0x11, 0x1b, 0xa9, 0x8a,
	0xa1, 0x34, 0xcb, 0x38, 0x15, 0x48, 0x85, 0x12, 0xf1, 0xfd, 0x84, 0x72, 0xae, 0xee, 0x49, 0x3f,
	0x93, 0x8d, 0x17, 0x50, 0xee, 0x0d, 0x13, 0xca, 0x87, 0x2c, 0xf4, 0xef, 0x81, 0xeb, 0xb0, 0x3f,
	0x21, 0xe1, 0x98, 0x4a, 0xb4, 0x82, 0x53, 0xd1, 0xf8, 0xbe, 0x07, 0x07, 0x5e, 0xc2, 0x62, 0xc6,
	0x49, 0x88, 0x74, 0x80, 0x78, 0x5d, 0xbb, 0xbe, 0xa4, 0x8f, 0xf0, 0x86, 0x83, 0x2e, 0xa0, 0xd4,
	0x67, 0x91, 0xa0, 0x91, 0x90, 0x4d, 0x0e, 0x5b, 0xa7, 0xe6, 0x3f, 0x6e, 0x6a, 0x66, 0xbd, 0xcf,
	0x53, 0x0e, 0x67, 0x0d, 0xd0, 0x6b, 0x28, 0x72, 0x41, 0xc4, 0x98, 0xab, 0x79, 0x43, 0x69, 0x56,
	0x5b, 0xd6, 0x7f, 0xb7, 0xea, 0x4a, 0x0c, 0xaf, 0x71, 0xa4, 0xc1, 0xc1, 0x84, 0x09, 0xca, 0x5f,
	0xb1, 0x44, 0x2d, 0x18, 0xf9, 0x66, 0x19, 0xdf, 0x69, 0xd4, 0x80, 0x23, 0x59, 0xdb, 0x03, 0x12,
	0x44, 0x5c, 0xa8, 0xfb, 0xf2, 0xfb, 0x96, 0xb7, 0xba, 0x34, 0x17, 0x24, 0x11, 0x67, 0x21, 0xeb,
	0x7f, 0x56, 0x8b, 0x86, 0xd2, 0xcc, 0xe3, 0x0d, 0x07, 0x19, 0x70, 0x48, 0x6f, 0xe2, 0x20, 0xa1,
	0x69, 0xa0, 0x24, 0x03, 0x9b, 0x56, 0x63, 0x0a, 0xc7, 0x3b, 0xd7, 0xbc, 0xe7, 0x09, 0x9e, 0x42,
	0x25, 0x9b, 0x26, 0x66, 0x63, 0x41, 0xd7, 0xaf, 0xb8, 0x6d, 0x22, 0x04, 0x85, 0x88, 0x8c, 0xa8,
	0x9c, 0x4b, 0x19, 0xcb, 0x1a, 0x9d, 0x40, 0x31, 0x26, 0x09, 0x19, 0x71, 0xb5, 0x20, 0x5f, 0x65,
	0xad, 0x9e, 0xfd, 0x52, 0xa0, 0xba, 0x3d, 0x17, 0xd4, 0x82, 0xc7, 0x1e, 0xee, 0x78, 0x9d, 0xae,
	0x7d, 0x79, 0xd5, 0xed, 0xd9, 0xbd, 0xb7, 0xdd, 0x2b, 0xf7, 0x8d, 0xdb, 0x73, 0xed, 0x9e, 0xd3,
	0xae, 0xe5, 0xb4, 0x87, 0xb3, 0xb9, 0x71, 0x9c, 0x46, 0xdd, 0x28, 0x10, 0x01, 0x11, 0xd4, 0x47,
	0xa7, 0xa0, 0xee, 0x32, 0xb6, 0xe7, 0xe1, 0xce, 0x3b, 0xa7, 0x5d, 0x53, 0x34, 0x34, 0x9b, 0x1b,
	0xd5, 0x14, 0xb1, 0xe3, 0x38, 0x61, 0x93, 0xbf, 0x13, 0xd8, 0xb9, 0x70, 0xce, 0x57, 0x87, 0xec,
	0x6d, 0x12, 0x98, 0x7e, 0xa2, 0xfd, 0xd5, 0x19, 0x26, 0x3c, 0xda, 0x25, 0x9c, 0xf7, 0x9e, 0x8b,
	0x9d, 0x76, 0x2d, 0xaf, 0x3d, 0x98, 0xcd, 0x8d, 0x4a, 0x0a, 0x38, 0x72, 0xb2, 0xbe, 0x56, 0xf8,
	0xf2, 0x4d, 0xcf, 0x9d, 0x5d, 0xfe, 0x58, 0xe8, 0xca, 0xed, 0x42, 0x57, 0x7e, 0x2f, 0x74, 0xe5,
	0xeb, 0x52, 0xcf, 0xdd, 0x2e, 0xf5, 0xdc, 0xcf, 0xa5, 0x9e, 0xfb, 0xd0, 0x1a, 0x04, 0x62, 0xb5,
	0x1f, 0x7d, 0x36, 0xb2, 0xb6, 0x56, 0xc7, 0xca, 0x56, 0xc7, 0xba, 0xb1, 0xee, 0x7e, 0x50, 0x31,
	0x8d, 0x29, 0xbf, 0x2e, 0xca, 0xcc, 0xf3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x80, 0x9c, 0xac,
	0x41, 0xb9, 0x03, 0x00, 0x00,
}

func (m *Relayer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Relayer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Relayer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Threshold) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Threshold) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Threshold) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		i = encodeVarintRelayer(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Proposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ExpireBlock != 0 {
		i = encodeVarintRelayer(dAtA, i, uint64(m.ExpireBlock))
		i--
		dAtA[i] = 0x38
	}
	if m.StartBlock != 0 {
		i = encodeVarintRelayer(dAtA, i, uint64(m.StartBlock))
		i--
		dAtA[i] = 0x30
	}
	if len(m.VotesAgainst) > 0 {
		for iNdEx := len(m.VotesAgainst) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.VotesAgainst[iNdEx])
			copy(dAtA[i:], m.VotesAgainst[iNdEx])
			i = encodeVarintRelayer(dAtA, i, uint64(len(m.VotesAgainst[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.VotesFor) > 0 {
		for iNdEx := len(m.VotesFor) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.VotesFor[iNdEx])
			copy(dAtA[i:], m.VotesFor[iNdEx])
			i = encodeVarintRelayer(dAtA, i, uint64(len(m.VotesFor[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Status != 0 {
		i = encodeVarintRelayer(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if m.Content != nil {
		{
			size, err := m.Content.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRelayer(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ProposalId) > 0 {
		i -= len(m.ProposalId)
		copy(dAtA[i:], m.ProposalId)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.ProposalId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ProposalContent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalContent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalContent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Params) > 0 {
		i -= len(m.Params)
		copy(dAtA[i:], m.Params)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Params)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ProposalRoute) > 0 {
		i -= len(m.ProposalRoute)
		copy(dAtA[i:], m.ProposalRoute)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.ProposalRoute)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRelayer(dAtA []byte, offset int, v uint64) int {
	offset -= sovRelayer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Relayer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	return n
}

func (m *Threshold) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovRelayer(uint64(m.Value))
	}
	return n
}

func (m *Proposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ProposalId)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	if m.Content != nil {
		l = m.Content.Size()
		n += 1 + l + sovRelayer(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovRelayer(uint64(m.Status))
	}
	if len(m.VotesFor) > 0 {
		for _, s := range m.VotesFor {
			l = len(s)
			n += 1 + l + sovRelayer(uint64(l))
		}
	}
	if len(m.VotesAgainst) > 0 {
		for _, s := range m.VotesAgainst {
			l = len(s)
			n += 1 + l + sovRelayer(uint64(l))
		}
	}
	if m.StartBlock != 0 {
		n += 1 + sovRelayer(uint64(m.StartBlock))
	}
	if m.ExpireBlock != 0 {
		n += 1 + sovRelayer(uint64(m.ExpireBlock))
	}
	return n
}

func (m *ProposalContent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	l = len(m.ProposalRoute)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	l = len(m.Params)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	return n
}

func sovRelayer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRelayer(x uint64) (n int) {
	return sovRelayer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Relayer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelayer
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
			return fmt.Errorf("proto: Relayer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Relayer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
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
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
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
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRelayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRelayer
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
func (m *Threshold) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelayer
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
			return fmt.Errorf("proto: Threshold: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Threshold: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
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
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRelayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRelayer
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
func (m *Proposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelayer
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
			return fmt.Errorf("proto: Proposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
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
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposalId = append(m.ProposalId[:0], dAtA[iNdEx:postIndex]...)
			if m.ProposalId == nil {
				m.ProposalId = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
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
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Content == nil {
				m.Content = &ProposalContent{}
			}
			if err := m.Content.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ProposalStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotesFor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
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
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VotesFor = append(m.VotesFor, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotesAgainst", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
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
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VotesAgainst = append(m.VotesAgainst, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartBlock", wireType)
			}
			m.StartBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireBlock", wireType)
			}
			m.ExpireBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpireBlock |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRelayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRelayer
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
func (m *ProposalContent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelayer
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
			return fmt.Errorf("proto: ProposalContent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalContent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
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
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalRoute", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
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
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposalRoute = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
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
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelayer
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
				return ErrInvalidLengthRelayer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRelayer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Params = append(m.Params[:0], dAtA[iNdEx:postIndex]...)
			if m.Params == nil {
				m.Params = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRelayer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRelayer
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
func skipRelayer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRelayer
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
					return 0, ErrIntOverflowRelayer
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
					return 0, ErrIntOverflowRelayer
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
				return 0, ErrInvalidLengthRelayer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRelayer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRelayer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRelayer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRelayer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRelayer = fmt.Errorf("proto: unexpected end of group")
)
