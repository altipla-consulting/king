// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test/example/example.proto

package king_example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/altipla-consulting/king/test/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FooRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FooRequest) Reset()         { *m = FooRequest{} }
func (m *FooRequest) String() string { return proto.CompactTextString(m) }
func (*FooRequest) ProtoMessage()    {}
func (*FooRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_bdb5a6d9061f4b14, []int{0}
}
func (m *FooRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FooRequest.Unmarshal(m, b)
}
func (m *FooRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FooRequest.Marshal(b, m, deterministic)
}
func (dst *FooRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FooRequest.Merge(dst, src)
}
func (m *FooRequest) XXX_Size() int {
	return xxx_messageInfo_FooRequest.Size(m)
}
func (m *FooRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FooRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FooRequest proto.InternalMessageInfo

func (m *FooRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FooRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type BarRequest struct {
	Increments           int32    `protobuf:"varint,1,opt,name=increments,proto3" json:"increments,omitempty"`
	Trying               bool     `protobuf:"varint,2,opt,name=trying,proto3" json:"trying,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BarRequest) Reset()         { *m = BarRequest{} }
func (m *BarRequest) String() string { return proto.CompactTextString(m) }
func (*BarRequest) ProtoMessage()    {}
func (*BarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_bdb5a6d9061f4b14, []int{1}
}
func (m *BarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarRequest.Unmarshal(m, b)
}
func (m *BarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarRequest.Marshal(b, m, deterministic)
}
func (dst *BarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarRequest.Merge(dst, src)
}
func (m *BarRequest) XXX_Size() int {
	return xxx_messageInfo_BarRequest.Size(m)
}
func (m *BarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BarRequest proto.InternalMessageInfo

func (m *BarRequest) GetIncrements() int32 {
	if m != nil {
		return m.Increments
	}
	return 0
}

func (m *BarRequest) GetTrying() bool {
	if m != nil {
		return m.Trying
	}
	return false
}

func init() {
	proto.RegisterType((*FooRequest)(nil), "king.example.FooRequest")
	proto.RegisterType((*BarRequest)(nil), "king.example.BarRequest")
}

func init() { proto.RegisterFile("test/example/example.proto", fileDescriptor_example_bdb5a6d9061f4b14) }

var fileDescriptor_example_bdb5a6d9061f4b14 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4a, 0x04, 0x31,
	0x10, 0x45, 0x6d, 0x75, 0x06, 0x2d, 0x5c, 0x15, 0x32, 0x84, 0x5e, 0x88, 0xf4, 0xca, 0x55, 0x04,
	0x45, 0x0f, 0x30, 0xea, 0xec, 0xdc, 0xc4, 0x13, 0xc4, 0x50, 0x34, 0xc1, 0x49, 0xaa, 0x4d, 0x4a,
	0x71, 0x2e, 0xe0, 0xb9, 0xc5, 0x74, 0x86, 0x6e, 0xc4, 0x55, 0xa5, 0x7e, 0xfe, 0x0b, 0xff, 0x07,
	0x5a, 0xa1, 0x2c, 0xd7, 0xf4, 0x65, 0xc3, 0xb0, 0xa5, 0xfd, 0xd4, 0x43, 0x62, 0x61, 0x3c, 0x7b,
	0xf3, 0xb1, 0xd7, 0x55, 0x6b, 0x55, 0x71, 0x3a, 0x0e, 0x81, 0x63, 0x1d, 0xa3, 0xaf, 0xbb, 0x07,
	0xd8, 0x30, 0x1b, 0x7a, 0xff, 0xa0, 0x2c, 0x88, 0x70, 0x1c, 0x6d, 0x20, 0xd5, 0x5c, 0x36, 0x57,
	0xa7, 0xa6, 0x9c, 0xf1, 0x1c, 0x16, 0x14, 0xac, 0xdf, 0xaa, 0xc3, 0x22, 0x8e, 0x4b, 0xf7, 0x08,
	0xb0, 0xb6, 0x69, 0xcf, 0x5d, 0x00, 0xf8, 0xe8, 0x12, 0x05, 0x8a, 0x92, 0x0b, 0xbd, 0x30, 0x33,
	0x05, 0x57, 0xb0, 0x94, 0xb4, 0xf3, 0xb1, 0x2f, 0x8f, 0x9c, 0x98, 0xba, 0xdd, 0x7c, 0x37, 0xb0,
	0x7a, 0xe0, 0x28, 0xd6, 0xc9, 0x33, 0xe5, 0x6c, 0x7b, 0xca, 0x2f, 0x94, 0x3e, 0xbd, 0x23, 0xbc,
	0x83, 0xa3, 0x0d, 0x33, 0x2a, 0x3d, 0x2f, 0xa2, 0xa7, 0xac, 0x2d, 0x8e, 0x37, 0xb5, 0xcd, 0x53,
	0x18, 0x64, 0xd7, 0x1d, 0xfc, 0x62, 0x6b, 0x9b, 0xfe, 0x62, 0x53, 0xd4, 0xff, 0xb1, 0xd7, 0x65,
	0xf9, 0x8d, 0xdb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xf3, 0xbf, 0xd6, 0x53, 0x01, 0x00,
	0x00,
}
