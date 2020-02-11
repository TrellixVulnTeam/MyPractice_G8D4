// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CalculateRequest struct {
	FirstNum             int32    `protobuf:"varint,1,opt,name=firstNum,proto3" json:"firstNum,omitempty"`
	SecondNum            int32    `protobuf:"varint,2,opt,name=secondNum,proto3" json:"secondNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculateRequest) Reset()         { *m = CalculateRequest{} }
func (m *CalculateRequest) String() string { return proto.CompactTextString(m) }
func (*CalculateRequest) ProtoMessage()    {}
func (*CalculateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *CalculateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculateRequest.Unmarshal(m, b)
}
func (m *CalculateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculateRequest.Marshal(b, m, deterministic)
}
func (m *CalculateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculateRequest.Merge(m, src)
}
func (m *CalculateRequest) XXX_Size() int {
	return xxx_messageInfo_CalculateRequest.Size(m)
}
func (m *CalculateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculateRequest proto.InternalMessageInfo

func (m *CalculateRequest) GetFirstNum() int32 {
	if m != nil {
		return m.FirstNum
	}
	return 0
}

func (m *CalculateRequest) GetSecondNum() int32 {
	if m != nil {
		return m.SecondNum
	}
	return 0
}

type CalculateResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculateResponse) Reset()         { *m = CalculateResponse{} }
func (m *CalculateResponse) String() string { return proto.CompactTextString(m) }
func (*CalculateResponse) ProtoMessage()    {}
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *CalculateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculateResponse.Unmarshal(m, b)
}
func (m *CalculateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculateResponse.Marshal(b, m, deterministic)
}
func (m *CalculateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculateResponse.Merge(m, src)
}
func (m *CalculateResponse) XXX_Size() int {
	return xxx_messageInfo_CalculateResponse.Size(m)
}
func (m *CalculateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculateResponse proto.InternalMessageInfo

func (m *CalculateResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type ManyCalculateRequests struct {
	FirstNum             int32    `protobuf:"varint,1,opt,name=firstNum,proto3" json:"firstNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManyCalculateRequests) Reset()         { *m = ManyCalculateRequests{} }
func (m *ManyCalculateRequests) String() string { return proto.CompactTextString(m) }
func (*ManyCalculateRequests) ProtoMessage()    {}
func (*ManyCalculateRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *ManyCalculateRequests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManyCalculateRequests.Unmarshal(m, b)
}
func (m *ManyCalculateRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManyCalculateRequests.Marshal(b, m, deterministic)
}
func (m *ManyCalculateRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManyCalculateRequests.Merge(m, src)
}
func (m *ManyCalculateRequests) XXX_Size() int {
	return xxx_messageInfo_ManyCalculateRequests.Size(m)
}
func (m *ManyCalculateRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_ManyCalculateRequests.DiscardUnknown(m)
}

var xxx_messageInfo_ManyCalculateRequests proto.InternalMessageInfo

func (m *ManyCalculateRequests) GetFirstNum() int32 {
	if m != nil {
		return m.FirstNum
	}
	return 0
}

type ManyCalculateResponses struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManyCalculateResponses) Reset()         { *m = ManyCalculateResponses{} }
func (m *ManyCalculateResponses) String() string { return proto.CompactTextString(m) }
func (*ManyCalculateResponses) ProtoMessage()    {}
func (*ManyCalculateResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *ManyCalculateResponses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManyCalculateResponses.Unmarshal(m, b)
}
func (m *ManyCalculateResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManyCalculateResponses.Marshal(b, m, deterministic)
}
func (m *ManyCalculateResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManyCalculateResponses.Merge(m, src)
}
func (m *ManyCalculateResponses) XXX_Size() int {
	return xxx_messageInfo_ManyCalculateResponses.Size(m)
}
func (m *ManyCalculateResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_ManyCalculateResponses.DiscardUnknown(m)
}

var xxx_messageInfo_ManyCalculateResponses proto.InternalMessageInfo

func (m *ManyCalculateResponses) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type AvgCalculateRequests struct {
	Num                  float64  `protobuf:"fixed64,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvgCalculateRequests) Reset()         { *m = AvgCalculateRequests{} }
func (m *AvgCalculateRequests) String() string { return proto.CompactTextString(m) }
func (*AvgCalculateRequests) ProtoMessage()    {}
func (*AvgCalculateRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *AvgCalculateRequests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvgCalculateRequests.Unmarshal(m, b)
}
func (m *AvgCalculateRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvgCalculateRequests.Marshal(b, m, deterministic)
}
func (m *AvgCalculateRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvgCalculateRequests.Merge(m, src)
}
func (m *AvgCalculateRequests) XXX_Size() int {
	return xxx_messageInfo_AvgCalculateRequests.Size(m)
}
func (m *AvgCalculateRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_AvgCalculateRequests.DiscardUnknown(m)
}

var xxx_messageInfo_AvgCalculateRequests proto.InternalMessageInfo

func (m *AvgCalculateRequests) GetNum() float64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type AvgCalculateResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvgCalculateResponse) Reset()         { *m = AvgCalculateResponse{} }
func (m *AvgCalculateResponse) String() string { return proto.CompactTextString(m) }
func (*AvgCalculateResponse) ProtoMessage()    {}
func (*AvgCalculateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *AvgCalculateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvgCalculateResponse.Unmarshal(m, b)
}
func (m *AvgCalculateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvgCalculateResponse.Marshal(b, m, deterministic)
}
func (m *AvgCalculateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvgCalculateResponse.Merge(m, src)
}
func (m *AvgCalculateResponse) XXX_Size() int {
	return xxx_messageInfo_AvgCalculateResponse.Size(m)
}
func (m *AvgCalculateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AvgCalculateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AvgCalculateResponse proto.InternalMessageInfo

func (m *AvgCalculateResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type FindMaxRequests struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxRequests) Reset()         { *m = FindMaxRequests{} }
func (m *FindMaxRequests) String() string { return proto.CompactTextString(m) }
func (*FindMaxRequests) ProtoMessage()    {}
func (*FindMaxRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{6}
}

func (m *FindMaxRequests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxRequests.Unmarshal(m, b)
}
func (m *FindMaxRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxRequests.Marshal(b, m, deterministic)
}
func (m *FindMaxRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxRequests.Merge(m, src)
}
func (m *FindMaxRequests) XXX_Size() int {
	return xxx_messageInfo_FindMaxRequests.Size(m)
}
func (m *FindMaxRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxRequests.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxRequests proto.InternalMessageInfo

func (m *FindMaxRequests) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type FindMaxResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxResponse) Reset()         { *m = FindMaxResponse{} }
func (m *FindMaxResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaxResponse) ProtoMessage()    {}
func (*FindMaxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{7}
}

func (m *FindMaxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxResponse.Unmarshal(m, b)
}
func (m *FindMaxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxResponse.Marshal(b, m, deterministic)
}
func (m *FindMaxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxResponse.Merge(m, src)
}
func (m *FindMaxResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaxResponse.Size(m)
}
func (m *FindMaxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxResponse proto.InternalMessageInfo

func (m *FindMaxResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type SquareRootRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootRequest) Reset()         { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()    {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{8}
}

func (m *SquareRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootRequest.Unmarshal(m, b)
}
func (m *SquareRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootRequest.Marshal(b, m, deterministic)
}
func (m *SquareRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootRequest.Merge(m, src)
}
func (m *SquareRootRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRootRequest.Size(m)
}
func (m *SquareRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootRequest proto.InternalMessageInfo

func (m *SquareRootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquareRootResponse struct {
	NumberRoot           float64  `protobuf:"fixed64,1,opt,name=number_root,json=numberRoot,proto3" json:"number_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootResponse) Reset()         { *m = SquareRootResponse{} }
func (m *SquareRootResponse) String() string { return proto.CompactTextString(m) }
func (*SquareRootResponse) ProtoMessage()    {}
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{9}
}

func (m *SquareRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootResponse.Unmarshal(m, b)
}
func (m *SquareRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootResponse.Marshal(b, m, deterministic)
}
func (m *SquareRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootResponse.Merge(m, src)
}
func (m *SquareRootResponse) XXX_Size() int {
	return xxx_messageInfo_SquareRootResponse.Size(m)
}
func (m *SquareRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootResponse proto.InternalMessageInfo

func (m *SquareRootResponse) GetNumberRoot() float64 {
	if m != nil {
		return m.NumberRoot
	}
	return 0
}

func init() {
	proto.RegisterType((*CalculateRequest)(nil), "calculator.CalculateRequest")
	proto.RegisterType((*CalculateResponse)(nil), "calculator.CalculateResponse")
	proto.RegisterType((*ManyCalculateRequests)(nil), "calculator.ManyCalculateRequests")
	proto.RegisterType((*ManyCalculateResponses)(nil), "calculator.ManyCalculateResponses")
	proto.RegisterType((*AvgCalculateRequests)(nil), "calculator.AvgCalculateRequests")
	proto.RegisterType((*AvgCalculateResponse)(nil), "calculator.AvgCalculateResponse")
	proto.RegisterType((*FindMaxRequests)(nil), "calculator.FindMaxRequests")
	proto.RegisterType((*FindMaxResponse)(nil), "calculator.FindMaxResponse")
	proto.RegisterType((*SquareRootRequest)(nil), "calculator.SquareRootRequest")
	proto.RegisterType((*SquareRootResponse)(nil), "calculator.SquareRootResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5f, 0x4f, 0xfa, 0x30,
	0x14, 0xa5, 0x3f, 0x02, 0x3f, 0xb9, 0xfe, 0x83, 0x1b, 0x25, 0x66, 0xa2, 0x62, 0x7d, 0x99, 0x21,
	0x41, 0x22, 0xf1, 0x03, 0xa0, 0x89, 0x89, 0x46, 0x7c, 0x18, 0x6f, 0xbc, 0x98, 0x31, 0x2a, 0x21,
	0x81, 0x15, 0xda, 0x8d, 0xe8, 0xb7, 0xf2, 0x23, 0x1a, 0x58, 0x47, 0xcb, 0x18, 0xf3, 0xad, 0xf7,
	0x9e, 0x73, 0xcf, 0xe9, 0xee, 0x59, 0xc1, 0xf6, 0xdc, 0x89, 0x17, 0x4e, 0xdc, 0x80, 0x8b, 0x3b,
	0x7d, 0x9c, 0x0d, 0x8c, 0xa2, 0x39, 0x13, 0x3c, 0xe0, 0x08, 0xba, 0x43, 0xdf, 0xa0, 0xfc, 0xa4,
	0x2a, 0xe6, 0xb0, 0x79, 0xc8, 0x64, 0x80, 0x16, 0xec, 0x7d, 0x8e, 0x85, 0x0c, 0xde, 0xc3, 0xe9,
	0x19, 0xa9, 0x13, 0xbb, 0xe0, 0xac, 0x6b, 0xac, 0x41, 0x49, 0x32, 0x8f, 0xfb, 0xc3, 0x25, 0xf8,
	0x6f, 0x05, 0xea, 0x06, 0x6d, 0x40, 0xc5, 0x50, 0x93, 0x33, 0xee, 0x4b, 0x86, 0x55, 0x28, 0x0a,
	0x26, 0xc3, 0x49, 0xa0, 0xc4, 0x54, 0x45, 0xdb, 0x70, 0xda, 0x75, 0xfd, 0xef, 0xa4, 0xbd, 0xcc,
	0xf2, 0xa7, 0x2d, 0xa8, 0x26, 0x86, 0x22, 0x17, 0xb9, 0xd3, 0xc6, 0x86, 0x93, 0xce, 0x62, 0xb4,
	0xed, 0x52, 0x86, 0xbc, 0xaf, 0x0c, 0x88, 0xb3, 0x3c, 0xd2, 0x66, 0x92, 0x99, 0xfa, 0x01, 0x64,
	0xad, 0x7c, 0x03, 0xc7, 0xcf, 0x63, 0x7f, 0xd8, 0x75, 0xbf, 0xd2, 0x44, 0x0b, 0x91, 0xe8, 0xad,
	0x41, 0xfa, 0x63, 0x21, 0x0d, 0xa8, 0xf4, 0xe6, 0xa1, 0x2b, 0x98, 0xc3, 0x79, 0x10, 0x87, 0x51,
	0x85, 0xa2, 0x1f, 0x4e, 0x07, 0x4c, 0xc4, 0xe4, 0xa8, 0xa2, 0x0f, 0x80, 0x26, 0x59, 0x49, 0x5f,
	0xc1, 0x7e, 0x84, 0x7f, 0x08, 0xce, 0xe3, 0xfb, 0x42, 0xd4, 0x5a, 0x12, 0xef, 0x7f, 0xf2, 0x46,
	0xe0, 0x3d, 0x26, 0x16, 0x63, 0x8f, 0xe1, 0x2b, 0x94, 0xd6, 0x3d, 0xac, 0x35, 0x8d, 0x1f, 0x26,
	0xb9, 0x36, 0xeb, 0x62, 0x07, 0x1a, 0xf9, 0xd3, 0x1c, 0xf6, 0xe1, 0x70, 0x23, 0x20, 0xbc, 0x36,
	0x27, 0x52, 0x03, 0xb7, 0x68, 0x06, 0x45, 0xc5, 0x4b, 0x73, 0x2d, 0x82, 0x7d, 0x28, 0x77, 0x16,
	0x4c, 0xb8, 0x23, 0xa6, 0xe5, 0xeb, 0xe6, 0x6c, 0x5a, 0xd0, 0x56, 0x06, 0x23, 0xbe, 0xb5, 0x4d,
	0xf0, 0x05, 0xfe, 0xab, 0x9c, 0xf0, 0xdc, 0x1c, 0x48, 0x24, 0x6c, 0xa5, 0x83, 0x5a, 0xa8, 0x45,
	0xb0, 0x0b, 0xa0, 0xa3, 0xc1, 0x8d, 0x8d, 0x6d, 0xe5, 0x6b, 0x5d, 0xee, 0x82, 0x63, 0xc9, 0xc7,
	0xa3, 0xfe, 0x81, 0xf9, 0x9e, 0x07, 0xc5, 0xd5, 0x2b, 0x6e, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xb7, 0x3c, 0x03, 0xb5, 0xf1, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculateServiceClient is the client API for CalculateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculateServiceClient interface {
	Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error)
	ManyCalculate(ctx context.Context, in *ManyCalculateRequests, opts ...grpc.CallOption) (CalculateService_ManyCalculateClient, error)
	AverageCalculate(ctx context.Context, opts ...grpc.CallOption) (CalculateService_AverageCalculateClient, error)
	FindMax(ctx context.Context, opts ...grpc.CallOption) (CalculateService_FindMaxClient, error)
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calculateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculateServiceClient(cc grpc.ClientConnInterface) CalculateServiceClient {
	return &calculateServiceClient{cc}
}

func (c *calculateServiceClient) Calculate(ctx context.Context, in *CalculateRequest, opts ...grpc.CallOption) (*CalculateResponse, error) {
	out := new(CalculateResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculateService/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculateServiceClient) ManyCalculate(ctx context.Context, in *ManyCalculateRequests, opts ...grpc.CallOption) (CalculateService_ManyCalculateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculateService_serviceDesc.Streams[0], "/calculator.CalculateService/ManyCalculate", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceManyCalculateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculateService_ManyCalculateClient interface {
	Recv() (*ManyCalculateResponses, error)
	grpc.ClientStream
}

type calculateServiceManyCalculateClient struct {
	grpc.ClientStream
}

func (x *calculateServiceManyCalculateClient) Recv() (*ManyCalculateResponses, error) {
	m := new(ManyCalculateResponses)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) AverageCalculate(ctx context.Context, opts ...grpc.CallOption) (CalculateService_AverageCalculateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculateService_serviceDesc.Streams[1], "/calculator.CalculateService/AverageCalculate", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceAverageCalculateClient{stream}
	return x, nil
}

type CalculateService_AverageCalculateClient interface {
	Send(*AvgCalculateRequests) error
	CloseAndRecv() (*AvgCalculateResponse, error)
	grpc.ClientStream
}

type calculateServiceAverageCalculateClient struct {
	grpc.ClientStream
}

func (x *calculateServiceAverageCalculateClient) Send(m *AvgCalculateRequests) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceAverageCalculateClient) CloseAndRecv() (*AvgCalculateResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AvgCalculateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) FindMax(ctx context.Context, opts ...grpc.CallOption) (CalculateService_FindMaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculateService_serviceDesc.Streams[2], "/calculator.CalculateService/FindMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculateServiceFindMaxClient{stream}
	return x, nil
}

type CalculateService_FindMaxClient interface {
	Send(*FindMaxRequests) error
	Recv() (*FindMaxResponse, error)
	grpc.ClientStream
}

type calculateServiceFindMaxClient struct {
	grpc.ClientStream
}

func (x *calculateServiceFindMaxClient) Send(m *FindMaxRequests) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculateServiceFindMaxClient) Recv() (*FindMaxResponse, error) {
	m := new(FindMaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculateServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculateService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculateServiceServer is the server API for CalculateService service.
type CalculateServiceServer interface {
	Calculate(context.Context, *CalculateRequest) (*CalculateResponse, error)
	ManyCalculate(*ManyCalculateRequests, CalculateService_ManyCalculateServer) error
	AverageCalculate(CalculateService_AverageCalculateServer) error
	FindMax(CalculateService_FindMaxServer) error
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

// UnimplementedCalculateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculateServiceServer struct {
}

func (*UnimplementedCalculateServiceServer) Calculate(ctx context.Context, req *CalculateRequest) (*CalculateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (*UnimplementedCalculateServiceServer) ManyCalculate(req *ManyCalculateRequests, srv CalculateService_ManyCalculateServer) error {
	return status.Errorf(codes.Unimplemented, "method ManyCalculate not implemented")
}
func (*UnimplementedCalculateServiceServer) AverageCalculate(srv CalculateService_AverageCalculateServer) error {
	return status.Errorf(codes.Unimplemented, "method AverageCalculate not implemented")
}
func (*UnimplementedCalculateServiceServer) FindMax(srv CalculateService_FindMaxServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMax not implemented")
}
func (*UnimplementedCalculateServiceServer) SquareRoot(ctx context.Context, req *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}

func RegisterCalculateServiceServer(s *grpc.Server, srv CalculateServiceServer) {
	s.RegisterService(&_CalculateService_serviceDesc, srv)
}

func _CalculateService_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculateService/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).Calculate(ctx, req.(*CalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculateService_ManyCalculate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ManyCalculateRequests)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculateServiceServer).ManyCalculate(m, &calculateServiceManyCalculateServer{stream})
}

type CalculateService_ManyCalculateServer interface {
	Send(*ManyCalculateResponses) error
	grpc.ServerStream
}

type calculateServiceManyCalculateServer struct {
	grpc.ServerStream
}

func (x *calculateServiceManyCalculateServer) Send(m *ManyCalculateResponses) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculateService_AverageCalculate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).AverageCalculate(&calculateServiceAverageCalculateServer{stream})
}

type CalculateService_AverageCalculateServer interface {
	SendAndClose(*AvgCalculateResponse) error
	Recv() (*AvgCalculateRequests, error)
	grpc.ServerStream
}

type calculateServiceAverageCalculateServer struct {
	grpc.ServerStream
}

func (x *calculateServiceAverageCalculateServer) SendAndClose(m *AvgCalculateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceAverageCalculateServer) Recv() (*AvgCalculateRequests, error) {
	m := new(AvgCalculateRequests)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateService_FindMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculateServiceServer).FindMax(&calculateServiceFindMaxServer{stream})
}

type CalculateService_FindMaxServer interface {
	Send(*FindMaxResponse) error
	Recv() (*FindMaxRequests, error)
	grpc.ServerStream
}

type calculateServiceFindMaxServer struct {
	grpc.ServerStream
}

func (x *calculateServiceFindMaxServer) Send(m *FindMaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculateServiceFindMaxServer) Recv() (*FindMaxRequests, error) {
	m := new(FindMaxRequests)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculateService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculateServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculateService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculateServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculateService",
	HandlerType: (*CalculateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _CalculateService_Calculate_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalculateService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ManyCalculate",
			Handler:       _CalculateService_ManyCalculate_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AverageCalculate",
			Handler:       _CalculateService_AverageCalculate_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMax",
			Handler:       _CalculateService_FindMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
