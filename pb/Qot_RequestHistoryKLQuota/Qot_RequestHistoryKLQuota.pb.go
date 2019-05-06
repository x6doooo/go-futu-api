// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_RequestHistoryKLQuota.proto

package Qot_RequestHistoryKLQuota

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

type DetailItem struct {
	Security             *Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	RequestTime          *string   `protobuf:"bytes,2,req,name=requestTime" json:"requestTime,omitempty"`
	RequestTimeStamp     *int64    `protobuf:"varint,3,opt,name=requestTimeStamp" json:"requestTimeStamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DetailItem) Reset()         { *m = DetailItem{} }
func (m *DetailItem) String() string { return proto.CompactTextString(m) }
func (*DetailItem) ProtoMessage()    {}
func (*DetailItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e7bfd1756986128, []int{0}
}

func (m *DetailItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailItem.Unmarshal(m, b)
}
func (m *DetailItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailItem.Marshal(b, m, deterministic)
}
func (m *DetailItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailItem.Merge(m, src)
}
func (m *DetailItem) XXX_Size() int {
	return xxx_messageInfo_DetailItem.Size(m)
}
func (m *DetailItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailItem.DiscardUnknown(m)
}

var xxx_messageInfo_DetailItem proto.InternalMessageInfo

func (m *DetailItem) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *DetailItem) GetRequestTime() string {
	if m != nil && m.RequestTime != nil {
		return *m.RequestTime
	}
	return ""
}

func (m *DetailItem) GetRequestTimeStamp() int64 {
	if m != nil && m.RequestTimeStamp != nil {
		return *m.RequestTimeStamp
	}
	return 0
}

type C2S struct {
	BGetDetail           *bool    `protobuf:"varint,2,opt,name=bGetDetail" json:"bGetDetail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e7bfd1756986128, []int{1}
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

func (m *C2S) GetBGetDetail() bool {
	if m != nil && m.BGetDetail != nil {
		return *m.BGetDetail
	}
	return false
}

type S2C struct {
	UsedQuota            *int32        `protobuf:"varint,1,req,name=usedQuota" json:"usedQuota,omitempty"`
	RemainQuota          *int32        `protobuf:"varint,2,req,name=remainQuota" json:"remainQuota,omitempty"`
	DetailList           []*DetailItem `protobuf:"bytes,3,rep,name=detailList" json:"detailList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e7bfd1756986128, []int{2}
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

func (m *S2C) GetUsedQuota() int32 {
	if m != nil && m.UsedQuota != nil {
		return *m.UsedQuota
	}
	return 0
}

func (m *S2C) GetRemainQuota() int32 {
	if m != nil && m.RemainQuota != nil {
		return *m.RemainQuota
	}
	return 0
}

func (m *S2C) GetDetailList() []*DetailItem {
	if m != nil {
		return m.DetailList
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
	return fileDescriptor_2e7bfd1756986128, []int{3}
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
	return fileDescriptor_2e7bfd1756986128, []int{4}
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
	proto.RegisterType((*DetailItem)(nil), "Qot_RequestHistoryKLQuota.DetailItem")
	proto.RegisterType((*C2S)(nil), "Qot_RequestHistoryKLQuota.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_RequestHistoryKLQuota.S2C")
	proto.RegisterType((*Request)(nil), "Qot_RequestHistoryKLQuota.Request")
	proto.RegisterType((*Response)(nil), "Qot_RequestHistoryKLQuota.Response")
}

func init() { proto.RegisterFile("Qot_RequestHistoryKLQuota.proto", fileDescriptor_2e7bfd1756986128) }

var fileDescriptor_2e7bfd1756986128 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0x69, 0xb3, 0xbd, 0xdb, 0xce, 0xde, 0x8b, 0x11, 0x44, 0xe2, 0x90, 0x19, 0x0a, 0x83,
	0x22, 0x38, 0x4a, 0xf0, 0x4a, 0x2f, 0xab, 0xa8, 0x38, 0x2f, 0x96, 0xee, 0x5e, 0xea, 0x76, 0x90,
	0x82, 0x5d, 0x6a, 0x92, 0x5e, 0xec, 0x3b, 0x88, 0x9f, 0x59, 0xfa, 0x67, 0x2e, 0x20, 0xd3, 0xcb,
	0xf3, 0x3b, 0xbf, 0x24, 0x4f, 0x1e, 0x38, 0x5b, 0x28, 0xfb, 0x2c, 0xf1, 0xbd, 0x44, 0x63, 0xef,
	0x33, 0x63, 0x95, 0xde, 0x3e, 0xce, 0x17, 0xa5, 0xb2, 0xe9, 0xac, 0xd0, 0xca, 0x2a, 0x7a, 0x72,
	0x50, 0x18, 0xff, 0x8f, 0x55, 0x9e, 0xab, 0x4d, 0x23, 0x8e, 0x47, 0x95, 0xe8, 0x92, 0xe0, 0xc3,
	0x03, 0xb8, 0x41, 0x9b, 0x66, 0x6f, 0x0f, 0x16, 0x73, 0x1a, 0x41, 0xdf, 0xe0, 0xaa, 0xd4, 0x99,
	0xdd, 0x32, 0x8f, 0xfb, 0xe1, 0x50, 0x1c, 0xcd, 0x9c, 0x33, 0x49, 0xbb, 0x93, 0xdf, 0x16, 0xe5,
	0x30, 0xd4, 0xcd, 0xcb, 0xcb, 0x2c, 0x47, 0xe6, 0x73, 0x3f, 0x1c, 0x48, 0x17, 0xd1, 0x73, 0x18,
	0x39, 0x63, 0x62, 0xd3, 0xbc, 0x60, 0x84, 0x7b, 0x21, 0x91, 0x3f, 0x78, 0x30, 0x05, 0x12, 0x8b,
	0x84, 0x4e, 0x00, 0x5e, 0xee, 0xd0, 0x36, 0xc1, 0x98, 0xcf, 0xbd, 0xb0, 0x2f, 0x1d, 0x52, 0xa5,
	0x26, 0x89, 0x88, 0xe9, 0x29, 0x0c, 0x4a, 0x83, 0xeb, 0xfa, 0xab, 0x75, 0xde, 0xae, 0xdc, 0x83,
	0x26, 0x5a, 0x9e, 0x66, 0x9b, 0x66, 0xef, 0xd7, 0x7b, 0x17, 0xd1, 0x5b, 0x80, 0x75, 0x7d, 0xe3,
	0x3c, 0x33, 0x96, 0x11, 0x4e, 0xc2, 0xa1, 0x98, 0xce, 0x0e, 0xd7, 0xbd, 0x6f, 0x4a, 0x3a, 0x07,
	0x83, 0x6b, 0xe8, 0xb5, 0x3e, 0x8d, 0x80, 0xac, 0x84, 0x69, 0xbb, 0x9b, 0xfc, 0x72, 0x55, 0x2c,
	0x12, 0x59, 0xa9, 0xc1, 0xa7, 0x07, 0x7d, 0x89, 0xa6, 0x50, 0x1b, 0x83, 0x74, 0x02, 0x3d, 0x8d,
	0x76, 0xb9, 0x2d, 0xb0, 0xf9, 0xce, 0x55, 0xe7, 0xe2, 0x32, 0x8a, 0xe4, 0x0e, 0xd2, 0x63, 0xf8,
	0xa7, 0xd1, 0x3e, 0x99, 0xd7, 0xba, 0x94, 0x81, 0x6c, 0x27, 0xca, 0xa0, 0x87, 0x5a, 0xc7, 0x6a,
	0x8d, 0x75, 0xb5, 0x5d, 0xb9, 0x1b, 0xab, 0x40, 0x46, 0xac, 0x58, 0x87, 0x7b, 0x7f, 0x04, 0x4a,
	0x44, 0x2c, 0x2b, 0xf5, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x59, 0x5e, 0xac, 0x1a, 0x6f, 0x02, 0x00,
	0x00,
}