// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: relayers/relayer.proto

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

type Relayer struct {
	Arena string   `protobuf:"bytes,1,opt,name=arena,proto3" json:"arena,omitempty"`
	Denom string   `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Addrs []string `protobuf:"bytes,3,rep,name=addrs,proto3" json:"addrs,omitempty"`
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

func (m *Relayer) GetArena() string {
	if m != nil {
		return m.Arena
	}
	return ""
}

func (m *Relayer) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Relayer) GetAddrs() []string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

type Threshold struct {
	Arena string `protobuf:"bytes,1,opt,name=arena,proto3" json:"arena,omitempty"`
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Value uint32 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
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

func (m *Threshold) GetArena() string {
	if m != nil {
		return m.Arena
	}
	return ""
}

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

type LastVoter struct {
	Arena string `protobuf:"bytes,1,opt,name=arena,proto3" json:"arena,omitempty"`
	Denom string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Voter string `protobuf:"bytes,3,opt,name=voter,proto3" json:"voter,omitempty"`
}

func (m *LastVoter) Reset()         { *m = LastVoter{} }
func (m *LastVoter) String() string { return proto.CompactTextString(m) }
func (*LastVoter) ProtoMessage()    {}
func (*LastVoter) Descriptor() ([]byte, []int) {
	return fileDescriptor_40adf04c18425525, []int{2}
}
func (m *LastVoter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LastVoter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LastVoter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LastVoter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LastVoter.Merge(m, src)
}
func (m *LastVoter) XXX_Size() int {
	return m.Size()
}
func (m *LastVoter) XXX_DiscardUnknown() {
	xxx_messageInfo_LastVoter.DiscardUnknown(m)
}

var xxx_messageInfo_LastVoter proto.InternalMessageInfo

func (m *LastVoter) GetArena() string {
	if m != nil {
		return m.Arena
	}
	return ""
}

func (m *LastVoter) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *LastVoter) GetVoter() string {
	if m != nil {
		return m.Voter
	}
	return ""
}

func init() {
	proto.RegisterType((*Relayer)(nil), "stafihub.stafihub.relayers.Relayer")
	proto.RegisterType((*Threshold)(nil), "stafihub.stafihub.relayers.Threshold")
	proto.RegisterType((*LastVoter)(nil), "stafihub.stafihub.relayers.LastVoter")
}

func init() { proto.RegisterFile("relayers/relayer.proto", fileDescriptor_40adf04c18425525) }

var fileDescriptor_40adf04c18425525 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4a, 0xcd, 0x49,
	0xac, 0x4c, 0x2d, 0x2a, 0xd6, 0x87, 0x32, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xa4, 0x8a,
	0x4b, 0x12, 0xd3, 0x32, 0x33, 0x4a, 0x93, 0xf4, 0xe0, 0x0c, 0x98, 0x4a, 0x25, 0x6f, 0x2e, 0xf6,
	0x20, 0x08, 0x5b, 0x48, 0x84, 0x8b, 0x35, 0xb1, 0x28, 0x35, 0x2f, 0x51, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x33, 0x08, 0xc2, 0x01, 0x89, 0xa6, 0xa4, 0xe6, 0xe5, 0xe7, 0x4a, 0x30, 0x41, 0x44, 0xc1,
	0x1c, 0xb0, 0xda, 0x94, 0x94, 0xa2, 0x62, 0x09, 0x66, 0x05, 0x66, 0xb0, 0x5a, 0x10, 0x47, 0xc9,
	0x97, 0x8b, 0x33, 0x24, 0xa3, 0x28, 0xb5, 0x38, 0x23, 0x3f, 0x27, 0x85, 0x54, 0xe3, 0xca, 0x12,
	0x73, 0x4a, 0x53, 0x25, 0x98, 0x15, 0x18, 0x35, 0x78, 0x83, 0x20, 0x1c, 0x90, 0x71, 0x3e, 0x89,
	0xc5, 0x25, 0x61, 0xf9, 0x25, 0xa4, 0xbb, 0xae, 0x0c, 0xa4, 0x09, 0x6c, 0x1c, 0x67, 0x10, 0x84,
	0xe3, 0xe4, 0x7e, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e,
	0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xba, 0xe9, 0x99,
	0x25, 0xa0, 0x50, 0x49, 0xce, 0xcf, 0xd5, 0x87, 0x05, 0x11, 0x82, 0x51, 0xa1, 0x0f, 0x0f, 0xd7,
	0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0xb0, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xb9, 0x2b, 0xe2, 0x4b, 0x70, 0x01, 0x00, 0x00,
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
	if len(m.Addrs) > 0 {
		for iNdEx := len(m.Addrs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addrs[iNdEx])
			copy(dAtA[i:], m.Addrs[iNdEx])
			i = encodeVarintRelayer(dAtA, i, uint64(len(m.Addrs[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Arena) > 0 {
		i -= len(m.Arena)
		copy(dAtA[i:], m.Arena)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Arena)))
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
		dAtA[i] = 0x18
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Arena) > 0 {
		i -= len(m.Arena)
		copy(dAtA[i:], m.Arena)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Arena)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LastVoter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LastVoter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LastVoter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Arena) > 0 {
		i -= len(m.Arena)
		copy(dAtA[i:], m.Arena)
		i = encodeVarintRelayer(dAtA, i, uint64(len(m.Arena)))
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
	l = len(m.Arena)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	if len(m.Addrs) > 0 {
		for _, s := range m.Addrs {
			l = len(s)
			n += 1 + l + sovRelayer(uint64(l))
		}
	}
	return n
}

func (m *Threshold) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Arena)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovRelayer(uint64(m.Value))
	}
	return n
}

func (m *LastVoter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Arena)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovRelayer(uint64(l))
	}
	l = len(m.Voter)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Arena", wireType)
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
			m.Arena = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addrs", wireType)
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
			m.Addrs = append(m.Addrs, string(dAtA[iNdEx:postIndex]))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Arena", wireType)
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
			m.Arena = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
		case 3:
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
func (m *LastVoter) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LastVoter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LastVoter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Arena", wireType)
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
			m.Arena = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
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
			m.Voter = string(dAtA[iNdEx:postIndex])
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
