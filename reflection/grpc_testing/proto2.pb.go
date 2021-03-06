// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reflection/grpc_testing/proto2.proto

package grpc_testing

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ToBeExtended struct {
	Foo                          *int32   `protobuf:"varint,1,req,name=foo" json:"foo,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *ToBeExtended) Reset()         { *m = ToBeExtended{} }
func (m *ToBeExtended) String() string { return proto.CompactTextString(m) }
func (*ToBeExtended) ProtoMessage()    {}
func (*ToBeExtended) Descriptor() ([]byte, []int) {
	return fileDescriptor_dddbb2c1ebdcf2b8, []int{0}
}

var extRange_ToBeExtended = []proto.ExtensionRange{
	{Start: 10, End: 30},
}

func (*ToBeExtended) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_ToBeExtended
}

func (m *ToBeExtended) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToBeExtended.Unmarshal(m, b)
}
func (m *ToBeExtended) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToBeExtended.Marshal(b, m, deterministic)
}
func (m *ToBeExtended) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToBeExtended.Merge(m, src)
}
func (m *ToBeExtended) XXX_Size() int {
	return xxx_messageInfo_ToBeExtended.Size(m)
}
func (m *ToBeExtended) XXX_DiscardUnknown() {
	xxx_messageInfo_ToBeExtended.DiscardUnknown(m)
}

var xxx_messageInfo_ToBeExtended proto.InternalMessageInfo

func (m *ToBeExtended) GetFoo() int32 {
	if m != nil && m.Foo != nil {
		return *m.Foo
	}
	return 0
}

func init() {
	proto.RegisterType((*ToBeExtended)(nil), "grpc.testing.ToBeExtended")
}

func init() {
	proto.RegisterFile("reflection/grpc_testing/proto2.proto", fileDescriptor_dddbb2c1ebdcf2b8)
}

var fileDescriptor_dddbb2c1ebdcf2b8 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x4a, 0x4d, 0xcb,
	0x49, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x4f, 0x2f, 0x2a, 0x48, 0x8e, 0x2f, 0x49, 0x2d, 0x2e,
	0xc9, 0xcc, 0x4b, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x37, 0xd2, 0x03, 0x53, 0x42, 0x3c, 0x20,
	0x29, 0x3d, 0xa8, 0x94, 0x92, 0x1a, 0x17, 0x4f, 0x48, 0xbe, 0x53, 0xaa, 0x6b, 0x45, 0x49, 0x6a,
	0x5e, 0x4a, 0x6a, 0x8a, 0x90, 0x00, 0x17, 0x73, 0x5a, 0x7e, 0xbe, 0x04, 0xa3, 0x02, 0x93, 0x06,
	0x6b, 0x10, 0x88, 0xa9, 0xc5, 0xc2, 0xc1, 0x25, 0x20, 0xef, 0x64, 0x10, 0xa5, 0x97, 0x9e, 0x9f,
	0x9f, 0x9e, 0x93, 0xaa, 0x97, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f, 0x94, 0x0e, 0xb6,
	0x44, 0x1f, 0x87, 0xa5, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x05, 0x50, 0x64, 0x8e, 0x00,
	0x00, 0x00,
}
