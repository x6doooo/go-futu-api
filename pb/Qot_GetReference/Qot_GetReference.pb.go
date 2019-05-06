// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetReference.proto

package Qot_GetReference

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

type ReferenceType int32

const (
	ReferenceType_ReferenceType_Unknow  ReferenceType = 0
	ReferenceType_ReferenceType_Warrant ReferenceType = 1
)

var ReferenceType_name = map[int32]string{
	0: "ReferenceType_Unknow",
	1: "ReferenceType_Warrant",
}

var ReferenceType_value = map[string]int32{
	"ReferenceType_Unknow":  0,
	"ReferenceType_Warrant": 1,
}

func (x ReferenceType) Enum() *ReferenceType {
	p := new(ReferenceType)
	*p = x
	return p
}

func (x ReferenceType) String() string {
	return proto.EnumName(ReferenceType_name, int32(x))
}

func (x *ReferenceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ReferenceType_value, data, "ReferenceType")
	if err != nil {
		return err
	}
	*x = ReferenceType(value)
	return nil
}

func (ReferenceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ae9b6d763a42feb1, []int{0}
}

type C2S struct {
	Security             *Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	ReferenceType        *int32    `protobuf:"varint,2,req,name=referenceType" json:"referenceType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b6d763a42feb1, []int{0}
}

func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (m *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(m, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *C2S) GetReferenceType() int32 {
	if m != nil && m.ReferenceType != nil {
		return *m.ReferenceType
	}
	return 0
}

type S2C struct {
	StaticInfoList       []*SecurityStaticInfo `protobuf:"bytes,2,rep,name=staticInfoList" json:"staticInfoList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b6d763a42feb1, []int{1}
}

func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (m *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(m, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetStaticInfoList() []*SecurityStaticInfo {
	if m != nil {
		return m.StaticInfoList
	}
	return nil
}

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b6d763a42feb1, []int{2}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae9b6d763a42feb1, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterEnum("Qot_GetReference.ReferenceType", ReferenceType_name, ReferenceType_value)
	proto.RegisterType((*C2S)(nil), "Qot_GetReference.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetReference.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetReference.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetReference.Response")
}

func init() { proto.RegisterFile("Qot_GetReference.proto", fileDescriptor_ae9b6d763a42feb1) }

var fileDescriptor_ae9b6d763a42feb1 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x51, 0x4b, 0x32, 0x41,
	0x14, 0x86, 0xbf, 0xdd, 0xd5, 0x4f, 0x3b, 0x66, 0xc8, 0xa0, 0x32, 0x79, 0x21, 0xc3, 0x12, 0xb4,
	0x04, 0x89, 0x0c, 0x5d, 0x75, 0x3b, 0x51, 0x04, 0x79, 0xd1, 0x99, 0xa2, 0xab, 0x10, 0xd9, 0x8e,
	0x21, 0xe1, 0x8c, 0xcd, 0x8c, 0x84, 0x3f, 0xa0, 0xff, 0x1d, 0x6b, 0x2a, 0x6a, 0x5e, 0x9e, 0xe7,
	0x7d, 0x0e, 0xe7, 0xe5, 0x40, 0xfb, 0xd1, 0x86, 0xe1, 0x1d, 0x05, 0xa4, 0x31, 0x39, 0x32, 0x39,
	0xf5, 0x66, 0xce, 0x06, 0xcb, 0x1a, 0xfb, 0xbc, 0x73, 0xac, 0xec, 0x74, 0x6a, 0xcd, 0x6f, 0xde,
	0x59, 0xe6, 0xdb, 0x24, 0x7d, 0x85, 0x44, 0x49, 0xcd, 0xfa, 0x50, 0xf5, 0x94, 0xcf, 0xdd, 0x24,
	0x2c, 0x78, 0x24, 0xe2, 0xac, 0x26, 0x9b, 0xbd, 0x2d, 0x57, 0xaf, 0x32, 0xdc, 0x58, 0xec, 0x0c,
	0xea, 0x6e, 0x7d, 0xe5, 0x69, 0x31, 0x23, 0x1e, 0x8b, 0x38, 0x2b, 0xe3, 0x2e, 0x4c, 0x07, 0x90,
	0x68, 0xa9, 0xd8, 0x2d, 0x9c, 0xf8, 0x30, 0x0a, 0x93, 0xfc, 0xde, 0x8c, 0xed, 0xc3, 0xc4, 0x07,
	0x1e, 0x8b, 0x24, 0xab, 0xc9, 0xee, 0xa1, 0x23, 0x7a, 0x63, 0xe2, 0xde, 0x56, 0x2a, 0xa1, 0x82,
	0xf4, 0x39, 0x27, 0x1f, 0xd8, 0x39, 0x24, 0xb9, 0xf4, 0xab, 0xb2, 0xad, 0xde, 0x9f, 0x87, 0x28,
	0xa9, 0xb1, 0x30, 0xd2, 0xef, 0x08, 0xaa, 0x48, 0x7e, 0x66, 0x8d, 0x27, 0xd6, 0x85, 0x8a, 0xa3,
	0xb0, 0xec, 0x5b, 0x6c, 0x96, 0xaf, 0x4b, 0x97, 0x57, 0xfd, 0x3e, 0xae, 0x21, 0x6b, 0xc3, 0x7f,
	0x47, 0x61, 0xe0, 0xdf, 0x79, 0x2c, 0xa2, 0xec, 0x08, 0x57, 0x13, 0xe3, 0x50, 0x21, 0xe7, 0x94,
	0x7d, 0x23, 0x9e, 0x88, 0x28, 0x2b, 0xe3, 0x7a, 0x2c, 0x7a, 0x78, 0x99, 0xf3, 0x92, 0x88, 0x0e,
	0xf7, 0xd0, 0x52, 0x61, 0x61, 0x5c, 0xdc, 0x40, 0x1d, 0xb7, 0x7f, 0xc3, 0x38, 0x34, 0x77, 0xc0,
	0xf0, 0xd9, 0x7c, 0x18, 0xfb, 0xd5, 0xf8, 0xc7, 0x4e, 0xa1, 0xb5, 0x9b, 0xbc, 0x8c, 0x9c, 0x1b,
	0x99, 0xd0, 0x88, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x8e, 0x1f, 0xbf, 0xfa, 0x01, 0x00,
	0x00,
}
