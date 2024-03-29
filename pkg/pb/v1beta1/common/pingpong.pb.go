// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v1beta1/common/pingpong.proto

package common

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TestEnum int32

const (
	TestEnumUnknown TestEnum = 0
	TestEnumFirst   TestEnum = 1
)

var TestEnum_name = map[int32]string{
	0: "UNKNOWN_VALUE",
	1: "FIRST_VALUE",
}

var TestEnum_value = map[string]int32{
	"UNKNOWN_VALUE": 0,
	"FIRST_VALUE":   1,
}

func (x TestEnum) String() string {
	return proto.EnumName(TestEnum_name, int32(x))
}

func (TestEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f32312b25b9a9350, []int{0}
}

// Ping
type Ping struct {
	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32312b25b9a9350, []int{0}
}
func (m *Ping) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return m.Size()
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

// Pong
type Pong struct {
	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32312b25b9a9350, []int{1}
}
func (m *Pong) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return m.Size()
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

type TestMessage struct {
	TestID          string            `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty" bson:"_id,omitempty"`
	Email           string            `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name            string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Age             int32             `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	Tag             map[string]string `protobuf:"bytes,5,rep,name=tag,proto3" json:"tag,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Created         string            `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	Value           TestEnum          `protobuf:"varint,7,opt,name=value,proto3,enum=common.TestEnum" json:"value,omitempty"`
	TestCompoundKey string            `protobuf:"bytes,8,opt,name=test_compound_key,json=testCompoundKey,proto3" json:"test_compound_key,omitempty" bson:"test_compound_key"`
}

func (m *TestMessage) Reset()         { *m = TestMessage{} }
func (m *TestMessage) String() string { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()    {}
func (*TestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32312b25b9a9350, []int{2}
}
func (m *TestMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessage.Merge(m, src)
}
func (m *TestMessage) XXX_Size() int {
	return m.Size()
}
func (m *TestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessage proto.InternalMessageInfo

func (m *TestMessage) GetTestID() string {
	if m != nil {
		return m.TestID
	}
	return ""
}

func (m *TestMessage) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *TestMessage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestMessage) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *TestMessage) GetTag() map[string]string {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *TestMessage) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *TestMessage) GetValue() TestEnum {
	if m != nil {
		return m.Value
	}
	return TestEnumUnknown
}

func (m *TestMessage) GetTestCompoundKey() string {
	if m != nil {
		return m.TestCompoundKey
	}
	return ""
}

func init() {
	proto.RegisterEnum("common.TestEnum", TestEnum_name, TestEnum_value)
	proto.RegisterType((*Ping)(nil), "common.Ping")
	proto.RegisterType((*Pong)(nil), "common.Pong")
	proto.RegisterType((*TestMessage)(nil), "common.TestMessage")
	proto.RegisterMapType((map[string]string)(nil), "common.TestMessage.TagEntry")
}

func init() { proto.RegisterFile("v1beta1/common/pingpong.proto", fileDescriptor_f32312b25b9a9350) }

var fileDescriptor_f32312b25b9a9350 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0xce, 0xe6, 0x57, 0xe3, 0xc4, 0xda, 0x74, 0xec, 0x61, 0x58, 0xd2, 0xcd, 0xb2, 0x42, 0x09,
	0x45, 0x77, 0x69, 0x04, 0x91, 0x78, 0xb2, 0x9a, 0x62, 0x88, 0x46, 0x59, 0x13, 0x0b, 0x22, 0x84,
	0x49, 0x32, 0x8e, 0x43, 0xb3, 0x33, 0x4b, 0x76, 0x52, 0xc9, 0xd5, 0x93, 0xf4, 0x24, 0x78, 0xee,
	0xc9, 0x7f, 0xc6, 0x63, 0xc1, 0x8b, 0xa7, 0x20, 0x89, 0x7f, 0x41, 0x8f, 0x9e, 0x64, 0x66, 0xb2,
	0x18, 0xf0, 0xf6, 0x7d, 0xef, 0x7d, 0x33, 0xef, 0xbd, 0x8f, 0x0f, 0xec, 0x9f, 0x1f, 0x0d, 0x89,
	0xc4, 0x47, 0xc1, 0x48, 0x44, 0x91, 0xe0, 0x41, 0xcc, 0x38, 0x8d, 0x05, 0xa7, 0x7e, 0x3c, 0x15,
	0x52, 0xc0, 0xa2, 0x29, 0xdb, 0x55, 0x2a, 0x04, 0x9d, 0x90, 0x00, 0xc7, 0x2c, 0xc0, 0x9c, 0x0b,
	0x89, 0x25, 0x13, 0x3c, 0x31, 0x2a, 0xfb, 0x1e, 0x65, 0xf2, 0xc3, 0x6c, 0xe8, 0x8f, 0x44, 0x14,
	0x50, 0x41, 0x45, 0xa0, 0xcb, 0xc3, 0xd9, 0x7b, 0xcd, 0x34, 0xd1, 0xc8, 0xc8, 0x3d, 0x1b, 0xe4,
	0x5f, 0x31, 0x4e, 0x21, 0x04, 0x79, 0x35, 0x0e, 0x59, 0xae, 0x55, 0xbf, 0x11, 0x6a, 0xac, 0x7b,
	0x62, 0xdd, 0x13, 0x1b, 0x3d, 0xc1, 0xa9, 0xf7, 0x27, 0x0b, 0xca, 0x3d, 0x92, 0xc8, 0x17, 0x24,
	0x49, 0x30, 0x25, 0xf0, 0x11, 0xd8, 0x92, 0x24, 0x91, 0x03, 0x36, 0x36, 0xb2, 0x63, 0x6f, 0xb9,
	0xa8, 0x15, 0x95, 0xa2, 0xfd, 0xf4, 0x7a, 0x51, 0xdb, 0x1b, 0x26, 0x82, 0x37, 0xbd, 0x01, 0x1b,
	0xdf, 0x15, 0x11, 0x93, 0x24, 0x8a, 0xe5, 0xdc, 0x0b, 0x8b, 0xea, 0x49, 0x7b, 0x0c, 0xf7, 0x40,
	0x81, 0x44, 0x98, 0x4d, 0x50, 0x56, 0x4f, 0x30, 0x44, 0x8d, 0xe5, 0x38, 0x22, 0x28, 0x67, 0xc6,
	0x2a, 0x0c, 0x2b, 0x20, 0x87, 0x29, 0x41, 0x79, 0xd7, 0xaa, 0x17, 0x42, 0x05, 0xa1, 0x0f, 0x72,
	0x12, 0x53, 0x54, 0x70, 0x73, 0xf5, 0x72, 0xa3, 0xea, 0x1b, 0x8f, 0xfc, 0x8d, 0xd5, 0xfc, 0x1e,
	0xa6, 0x2d, 0x2e, 0xa7, 0xf3, 0x50, 0x09, 0x21, 0x02, 0x5b, 0xa3, 0x29, 0xc1, 0x92, 0x8c, 0x51,
	0x51, 0x7f, 0x9c, 0x52, 0x78, 0x00, 0x0a, 0xe7, 0x78, 0x32, 0x23, 0x68, 0xcb, 0xb5, 0xea, 0xb7,
	0x1a, 0x95, 0xcd, 0xbf, 0x5a, 0x7c, 0x16, 0x85, 0xa6, 0x0d, 0x9f, 0x81, 0x5d, 0x7d, 0xea, 0x48,
	0x44, 0xb1, 0x98, 0xf1, 0xf1, 0xe0, 0x8c, 0xcc, 0x51, 0x49, 0x1f, 0x5d, 0xbd, 0x5e, 0xd4, 0x90,
	0x39, 0xf5, 0x3f, 0x89, 0x17, 0xee, 0xa8, 0xda, 0x93, 0x75, 0xa9, 0x43, 0xe6, 0xf6, 0x03, 0x50,
	0x4a, 0x97, 0x53, 0x97, 0xa9, 0x7f, 0x8c, 0xc7, 0x0a, 0x2a, 0x57, 0xcc, 0x3e, 0x6b, 0x57, 0x34,
	0x69, 0x66, 0x1f, 0x5a, 0x87, 0xef, 0x40, 0x29, 0x5d, 0x0a, 0x1e, 0x80, 0xed, 0x7e, 0xb7, 0xd3,
	0x7d, 0x79, 0xda, 0x1d, 0xbc, 0x79, 0xfc, 0xbc, 0xdf, 0xaa, 0x64, 0xec, 0xdb, 0x17, 0x97, 0xee,
	0x4e, 0x2a, 0xe8, 0xf3, 0x33, 0x2e, 0x3e, 0x72, 0xe8, 0x81, 0xf2, 0x49, 0x3b, 0x7c, 0xdd, 0x5b,
	0xab, 0x2c, 0x7b, 0xf7, 0xe2, 0xd2, 0xdd, 0x4e, 0x55, 0x27, 0x6c, 0x9a, 0x48, 0x3b, 0xff, 0xf9,
	0x9b, 0x93, 0x69, 0x9c, 0x82, 0x92, 0x8a, 0x84, 0x6a, 0xc1, 0x8e, 0xc1, 0x3a, 0x06, 0x37, 0x53,
	0x43, 0x54, 0xc5, 0xfe, 0xc7, 0x54, 0x1c, 0xee, 0x7c, 0xfa, 0xf1, 0xfb, 0x6b, 0x76, 0xdf, 0x43,
	0x69, 0x76, 0xd3, 0x28, 0xa7, 0x19, 0x6e, 0x5a, 0x87, 0xc7, 0xee, 0xf7, 0xa5, 0x63, 0x5d, 0x2d,
	0x1d, 0xeb, 0xd7, 0xd2, 0xb1, 0xbe, 0xac, 0x9c, 0xcc, 0xd5, 0xca, 0xc9, 0xfc, 0x5c, 0x39, 0x99,
	0xb7, 0xeb, 0x68, 0x0f, 0x8b, 0x3a, 0x94, 0xf7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x84, 0x5e,
	0x42, 0x82, 0x0a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingTestClient is the client API for PingTest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingTestClient interface {
	// Send a ping and receive a pong
	PingPong(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type pingTestClient struct {
	cc *grpc.ClientConn
}

func NewPingTestClient(cc *grpc.ClientConn) PingTestClient {
	return &pingTestClient{cc}
}

func (c *pingTestClient) PingPong(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/common.PingTest/PingPong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingTestServer is the server API for PingTest service.
type PingTestServer interface {
	// Send a ping and receive a pong
	PingPong(context.Context, *Ping) (*Pong, error)
}

// UnimplementedPingTestServer can be embedded to have forward compatible implementations.
type UnimplementedPingTestServer struct {
}

func (*UnimplementedPingTestServer) PingPong(ctx context.Context, req *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}

func RegisterPingTestServer(s *grpc.Server, srv PingTestServer) {
	s.RegisterService(&_PingTest_serviceDesc, srv)
}

func _PingTest_PingPong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingTestServer).PingPong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.PingTest/PingPong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingTestServer).PingPong(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingTest_serviceDesc = grpc.ServiceDesc{
	ServiceName: "common.PingTest",
	HandlerType: (*PingTestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingPong",
			Handler:    _PingTest_PingPong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1beta1/common/pingpong.proto",
}

func (m *Ping) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ping) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Ping) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPingpong(dAtA, i, uint64(len(m.Ping)))
		i += copy(dAtA[i:], m.Ping)
	}
	return i, nil
}

func (m *Pong) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pong) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Pong) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPingpong(dAtA, i, uint64(len(m.Pong)))
		i += copy(dAtA[i:], m.Pong)
	}
	return i, nil
}

func (m *TestMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TestID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPingpong(dAtA, i, uint64(len(m.TestID)))
		i += copy(dAtA[i:], m.TestID)
	}
	if len(m.Email) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPingpong(dAtA, i, uint64(len(m.Email)))
		i += copy(dAtA[i:], m.Email)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPingpong(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Age != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintPingpong(dAtA, i, uint64(m.Age))
	}
	if len(m.Tag) > 0 {
		for k, _ := range m.Tag {
			dAtA[i] = 0x2a
			i++
			v := m.Tag[k]
			mapSize := 1 + len(k) + sovPingpong(uint64(len(k))) + 1 + len(v) + sovPingpong(uint64(len(v)))
			i = encodeVarintPingpong(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintPingpong(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintPingpong(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.Created) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintPingpong(dAtA, i, uint64(len(m.Created)))
		i += copy(dAtA[i:], m.Created)
	}
	if m.Value != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintPingpong(dAtA, i, uint64(m.Value))
	}
	if len(m.TestCompoundKey) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintPingpong(dAtA, i, uint64(len(m.TestCompoundKey)))
		i += copy(dAtA[i:], m.TestCompoundKey)
	}
	return i, nil
}

func encodeVarintPingpong(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Ping) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Ping)
	if l > 0 {
		n += 1 + l + sovPingpong(uint64(l))
	}
	return n
}

func (m *Pong) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pong)
	if l > 0 {
		n += 1 + l + sovPingpong(uint64(l))
	}
	return n
}

func (m *TestMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TestID)
	if l > 0 {
		n += 1 + l + sovPingpong(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovPingpong(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovPingpong(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovPingpong(uint64(m.Age))
	}
	if len(m.Tag) > 0 {
		for k, v := range m.Tag {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovPingpong(uint64(len(k))) + 1 + len(v) + sovPingpong(uint64(len(v)))
			n += mapEntrySize + 1 + sovPingpong(uint64(mapEntrySize))
		}
	}
	l = len(m.Created)
	if l > 0 {
		n += 1 + l + sovPingpong(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovPingpong(uint64(m.Value))
	}
	l = len(m.TestCompoundKey)
	if l > 0 {
		n += 1 + l + sovPingpong(uint64(l))
	}
	return n
}

func sovPingpong(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPingpong(x uint64) (n int) {
	return sovPingpong(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ping) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPingpong
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
			return fmt.Errorf("proto: Ping: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ping: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ping", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPingpong
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
				return ErrInvalidLengthPingpong
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPingpong
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ping = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPingpong(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPingpong
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPingpong
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
func (m *Pong) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPingpong
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
			return fmt.Errorf("proto: Pong: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pong: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pong", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPingpong
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
				return ErrInvalidLengthPingpong
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPingpong
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pong = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPingpong(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPingpong
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPingpong
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
func (m *TestMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPingpong
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
			return fmt.Errorf("proto: TestMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPingpong
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
				return ErrInvalidLengthPingpong
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPingpong
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TestID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPingpong
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
				return ErrInvalidLengthPingpong
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPingpong
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPingpong
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
				return ErrInvalidLengthPingpong
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPingpong
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPingpong
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tag", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPingpong
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
				return ErrInvalidLengthPingpong
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPingpong
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tag == nil {
				m.Tag = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPingpong
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPingpong
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthPingpong
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthPingpong
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPingpong
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthPingpong
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthPingpong
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipPingpong(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthPingpong
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Tag[mapkey] = mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPingpong
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
				return ErrInvalidLengthPingpong
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPingpong
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Created = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPingpong
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= TestEnum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TestCompoundKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPingpong
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
				return ErrInvalidLengthPingpong
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPingpong
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TestCompoundKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPingpong(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPingpong
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPingpong
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
func skipPingpong(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPingpong
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
					return 0, ErrIntOverflowPingpong
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPingpong
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
				return 0, ErrInvalidLengthPingpong
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthPingpong
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPingpong
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPingpong(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthPingpong
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPingpong = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPingpong   = fmt.Errorf("proto: integer overflow")
)
