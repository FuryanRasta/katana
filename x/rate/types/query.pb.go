// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rate/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QueryRateRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryRateRequest) Reset()         { *m = QueryRateRequest{} }
func (m *QueryRateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRateRequest) ProtoMessage()    {}
func (*QueryRateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d2a6803f5955cc, []int{0}
}
func (m *QueryRateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRateRequest.Merge(m, src)
}
func (m *QueryRateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRateRequest proto.InternalMessageInfo

func (m *QueryRateRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type QueryRateResponse struct {
	Ratio *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=ratio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"ratio,omitempty"`
}

func (m *QueryRateResponse) Reset()         { *m = QueryRateResponse{} }
func (m *QueryRateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRateResponse) ProtoMessage()    {}
func (*QueryRateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d2a6803f5955cc, []int{1}
}
func (m *QueryRateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRateResponse.Merge(m, src)
}
func (m *QueryRateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRateResponse proto.InternalMessageInfo

type QueryEraRateRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Era   string `protobuf:"bytes,2,opt,name=era,proto3" json:"era,omitempty"`
}

func (m *QueryEraRateRequest) Reset()         { *m = QueryEraRateRequest{} }
func (m *QueryEraRateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEraRateRequest) ProtoMessage()    {}
func (*QueryEraRateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d2a6803f5955cc, []int{2}
}
func (m *QueryEraRateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEraRateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEraRateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEraRateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEraRateRequest.Merge(m, src)
}
func (m *QueryEraRateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEraRateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEraRateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEraRateRequest proto.InternalMessageInfo

func (m *QueryEraRateRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *QueryEraRateRequest) GetEra() string {
	if m != nil {
		return m.Era
	}
	return ""
}

type QueryEraRateResponse struct {
	Ratio *github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=ratio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"ratio,omitempty"`
}

func (m *QueryEraRateResponse) Reset()         { *m = QueryEraRateResponse{} }
func (m *QueryEraRateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEraRateResponse) ProtoMessage()    {}
func (*QueryEraRateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4d2a6803f5955cc, []int{3}
}
func (m *QueryEraRateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEraRateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEraRateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEraRateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEraRateResponse.Merge(m, src)
}
func (m *QueryEraRateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEraRateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEraRateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEraRateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryRateRequest)(nil), "stafiprotocol.stafihub.rate.QueryRateRequest")
	proto.RegisterType((*QueryRateResponse)(nil), "stafiprotocol.stafihub.rate.QueryRateResponse")
	proto.RegisterType((*QueryEraRateRequest)(nil), "stafiprotocol.stafihub.rate.QueryEraRateRequest")
	proto.RegisterType((*QueryEraRateResponse)(nil), "stafiprotocol.stafihub.rate.QueryEraRateResponse")
}

func init() { proto.RegisterFile("rate/query.proto", fileDescriptor_a4d2a6803f5955cc) }

var fileDescriptor_a4d2a6803f5955cc = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x4a, 0x2c, 0x49,
	0xd5, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2e, 0x2e,
	0x49, 0x4c, 0xcb, 0x04, 0xb3, 0x93, 0xf3, 0x73, 0xf4, 0xc0, 0xbc, 0x8c, 0xd2, 0x24, 0x3d, 0x90,
	0x42, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0x9c, 0x3e, 0x88, 0x05, 0xd1, 0x22, 0x25, 0x93,
	0x9e, 0x9f, 0x9f, 0x9e, 0x93, 0xaa, 0x9f, 0x58, 0x90, 0xa9, 0x9f, 0x98, 0x97, 0x97, 0x5f, 0x92,
	0x58, 0x92, 0x99, 0x9f, 0x57, 0x0c, 0x91, 0x55, 0xd2, 0xe0, 0x12, 0x08, 0x04, 0x99, 0x1f, 0x94,
	0x58, 0x92, 0x1a, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc2, 0xc5, 0x9a, 0x92, 0x9a,
	0x97, 0x9f, 0x2b, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x28, 0x85, 0x72, 0x09, 0x22,
	0xa9, 0x2c, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x72, 0xe0, 0x62, 0x2d, 0x02, 0x99, 0x07, 0x51,
	0xea, 0xa4, 0x75, 0xeb, 0x9e, 0xbc, 0x5a, 0x7a, 0x66, 0x09, 0xc8, 0x49, 0xc9, 0xf9, 0xb9, 0xfa,
	0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0x50, 0x4a, 0xb7, 0x38, 0x25, 0x5b, 0xbf, 0xa4, 0xb2, 0x20,
	0xb5, 0x58, 0xcf, 0x25, 0x35, 0x39, 0x08, 0xa2, 0x51, 0xc9, 0x96, 0x4b, 0x18, 0x6c, 0xac, 0x6b,
	0x51, 0x22, 0x41, 0x37, 0x08, 0x09, 0x70, 0x31, 0xa7, 0x16, 0x25, 0x4a, 0x30, 0x81, 0xc5, 0x40,
	0x4c, 0xa5, 0x08, 0x2e, 0x11, 0x54, 0xed, 0xd4, 0x72, 0x98, 0xd1, 0x29, 0x26, 0x2e, 0x56, 0xb0,
	0xd1, 0x42, 0x93, 0x19, 0xb9, 0x58, 0x40, 0x86, 0x0b, 0xe9, 0xea, 0xe1, 0x09, 0x7e, 0x3d, 0xf4,
	0x70, 0x94, 0xd2, 0x23, 0x56, 0x39, 0xc4, 0xcd, 0x4a, 0x3a, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x52,
	0x13, 0x52, 0xd1, 0x47, 0xd1, 0xa7, 0x0f, 0xd3, 0xa7, 0x0f, 0x4e, 0x0e, 0xd5, 0xe0, 0xa0, 0xa8,
	0x15, 0x5a, 0xc5, 0xc8, 0xc5, 0x0e, 0xf5, 0xb5, 0x90, 0x01, 0x61, 0x9b, 0x50, 0xc3, 0x57, 0xca,
	0x90, 0x04, 0x1d, 0x50, 0xe7, 0x59, 0x81, 0x9d, 0x67, 0x22, 0x64, 0x84, 0xd7, 0x79, 0xa9, 0x10,
	0x5d, 0x30, 0x67, 0xea, 0x57, 0xa7, 0x16, 0x25, 0xd6, 0x3a, 0x79, 0x9c, 0x78, 0x24, 0xc7, 0x78,
	0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7,
	0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x1e, 0x52, 0xac, 0xe0, 0x30, 0xb7, 0x02, 0x62, 0x32, 0x38,
	0x86, 0x92, 0xd8, 0xc0, 0xf2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x8b, 0xf6, 0xad,
	0x1c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a list of rate items.
	Rate(ctx context.Context, in *QueryRateRequest, opts ...grpc.CallOption) (*QueryRateResponse, error)
	// Queries a list of eraRate items.
	EraRate(ctx context.Context, in *QueryEraRateRequest, opts ...grpc.CallOption) (*QueryEraRateResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Rate(ctx context.Context, in *QueryRateRequest, opts ...grpc.CallOption) (*QueryRateResponse, error) {
	out := new(QueryRateResponse)
	err := c.cc.Invoke(ctx, "/stafiprotocol.stafihub.rate.Query/Rate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EraRate(ctx context.Context, in *QueryEraRateRequest, opts ...grpc.CallOption) (*QueryEraRateResponse, error) {
	out := new(QueryEraRateResponse)
	err := c.cc.Invoke(ctx, "/stafiprotocol.stafihub.rate.Query/EraRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a list of rate items.
	Rate(context.Context, *QueryRateRequest) (*QueryRateResponse, error)
	// Queries a list of eraRate items.
	EraRate(context.Context, *QueryEraRateRequest) (*QueryEraRateResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Rate(ctx context.Context, req *QueryRateRequest) (*QueryRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rate not implemented")
}
func (*UnimplementedQueryServer) EraRate(ctx context.Context, req *QueryEraRateRequest) (*QueryEraRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EraRate not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Rate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Rate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stafiprotocol.stafihub.rate.Query/Rate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Rate(ctx, req.(*QueryRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EraRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEraRateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EraRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stafiprotocol.stafihub.rate.Query/EraRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EraRate(ctx, req.(*QueryEraRateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stafiprotocol.stafihub.rate.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Rate",
			Handler:    _Query_Rate_Handler,
		},
		{
			MethodName: "EraRate",
			Handler:    _Query_EraRate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rate/query.proto",
}

func (m *QueryRateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ratio != nil {
		{
			size := m.Ratio.Size()
			i -= size
			if _, err := m.Ratio.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryEraRateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEraRateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEraRateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Era) > 0 {
		i -= len(m.Era)
		copy(dAtA[i:], m.Era)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Era)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryEraRateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEraRateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEraRateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ratio != nil {
		{
			size := m.Ratio.Size()
			i -= size
			if _, err := m.Ratio.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ratio != nil {
		l = m.Ratio.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryEraRateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Era)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryEraRateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ratio != nil {
		l = m.Ratio.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryRateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ratio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.Ratio = &v
			if err := m.Ratio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEraRateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEraRateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEraRateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Era", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Era = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEraRateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEraRateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEraRateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ratio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Dec
			m.Ratio = &v
			if err := m.Ratio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
