// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spotigraph/person/v1/person.proto

package personv1

import (
	fmt "fmt"

	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Person represents an user identity.
type Person struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Urn                  string   `protobuf:"bytes,2,opt,name=urn,proto3" json:"urn,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_d40f8a6aef2eb26d, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}

func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}

func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}

func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}

func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Person) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "spotigraph.person.v1.Person")
}

func init() { proto.RegisterFile("spotigraph/person/v1/person.proto", fileDescriptor_d40f8a6aef2eb26d) }

var fileDescriptor_d40f8a6aef2eb26d = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x2e, 0xc8, 0x2f,
	0xc9, 0x4c, 0x2f, 0x4a, 0x2c, 0xc8, 0xd0, 0x2f, 0x48, 0x2d, 0x2a, 0xce, 0xcf, 0xd3, 0x2f, 0x33,
	0x84, 0xb2, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0x10, 0x4a, 0xf4, 0xa0, 0x12, 0x65,
	0x86, 0x4a, 0x69, 0x5c, 0x6c, 0x01, 0x60, 0x8e, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x00, 0x17, 0x73, 0x69, 0x51, 0x9e, 0x04,
	0x13, 0x58, 0x00, 0xc4, 0x14, 0x92, 0xe5, 0xe2, 0x4a, 0xcb, 0x2c, 0x2a, 0x2e, 0x89, 0xcf, 0x4b,
	0xcc, 0x4d, 0x95, 0x60, 0x06, 0x4b, 0x70, 0x82, 0x45, 0xfc, 0x12, 0x73, 0x53, 0x85, 0xa4, 0xb9,
	0x38, 0x73, 0x12, 0x61, 0xb2, 0x2c, 0x60, 0x59, 0x0e, 0x90, 0x00, 0x48, 0xd2, 0xa9, 0x80, 0x4b,
	0x31, 0xbf, 0x28, 0x5d, 0xaf, 0x2a, 0x35, 0x2f, 0xb3, 0x24, 0x23, 0xb1, 0x48, 0x0f, 0x9b, 0x63,
	0x9c, 0xb8, 0x21, 0x4e, 0x09, 0x00, 0xb9, 0x37, 0x80, 0x31, 0x8a, 0x03, 0x22, 0x53, 0x66, 0xb8,
	0x88, 0x89, 0x39, 0x38, 0x20, 0x62, 0x15, 0x93, 0x48, 0x30, 0x42, 0x17, 0x44, 0xa9, 0x5e, 0x98,
	0xe1, 0x29, 0x64, 0xe1, 0x18, 0x88, 0x70, 0x4c, 0x98, 0x61, 0x12, 0x1b, 0xd8, 0xdb, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x4a, 0x0e, 0x94, 0x1b, 0x01, 0x00, 0x00,
}
