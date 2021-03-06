// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cblocks.proto

package cbservice

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

type LieAlgebra_LieType int32

const (
	LieAlgebra_A LieAlgebra_LieType = 0
	LieAlgebra_B LieAlgebra_LieType = 1
	LieAlgebra_C LieAlgebra_LieType = 2
	LieAlgebra_D LieAlgebra_LieType = 3
)

var LieAlgebra_LieType_name = map[int32]string{
	0: "A",
	1: "B",
	2: "C",
	3: "D",
}

var LieAlgebra_LieType_value = map[string]int32{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
}

func (x LieAlgebra_LieType) String() string {
	return proto.EnumName(LieAlgebra_LieType_name, int32(x))
}

func (LieAlgebra_LieType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c67941f876b8a60b, []int{2, 0}
}

// A conformal blocks request
type ConformalBlocksRequest struct {
	Algebra              *LieAlgebra `protobuf:"bytes,1,opt,name=algebra,proto3" json:"algebra,omitempty"`
	Weights              []*Weight   `protobuf:"bytes,2,rep,name=weights,proto3" json:"weights,omitempty"`
	Level                int64       `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ConformalBlocksRequest) Reset()         { *m = ConformalBlocksRequest{} }
func (m *ConformalBlocksRequest) String() string { return proto.CompactTextString(m) }
func (*ConformalBlocksRequest) ProtoMessage()    {}
func (*ConformalBlocksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67941f876b8a60b, []int{0}
}

func (m *ConformalBlocksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConformalBlocksRequest.Unmarshal(m, b)
}
func (m *ConformalBlocksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConformalBlocksRequest.Marshal(b, m, deterministic)
}
func (m *ConformalBlocksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConformalBlocksRequest.Merge(m, src)
}
func (m *ConformalBlocksRequest) XXX_Size() int {
	return xxx_messageInfo_ConformalBlocksRequest.Size(m)
}
func (m *ConformalBlocksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConformalBlocksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConformalBlocksRequest proto.InternalMessageInfo

func (m *ConformalBlocksRequest) GetAlgebra() *LieAlgebra {
	if m != nil {
		return m.Algebra
	}
	return nil
}

func (m *ConformalBlocksRequest) GetWeights() []*Weight {
	if m != nil {
		return m.Weights
	}
	return nil
}

func (m *ConformalBlocksRequest) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

// A symmetric conformal blocks request
type SymConformalBlocksRequest struct {
	Algebra              *LieAlgebra `protobuf:"bytes,1,opt,name=algebra,proto3" json:"algebra,omitempty"`
	Weight               *Weight     `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
	NumPoints            int64       `protobuf:"varint,3,opt,name=num_points,json=numPoints,proto3" json:"num_points,omitempty"`
	Level                int64       `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SymConformalBlocksRequest) Reset()         { *m = SymConformalBlocksRequest{} }
func (m *SymConformalBlocksRequest) String() string { return proto.CompactTextString(m) }
func (*SymConformalBlocksRequest) ProtoMessage()    {}
func (*SymConformalBlocksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67941f876b8a60b, []int{1}
}

func (m *SymConformalBlocksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SymConformalBlocksRequest.Unmarshal(m, b)
}
func (m *SymConformalBlocksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SymConformalBlocksRequest.Marshal(b, m, deterministic)
}
func (m *SymConformalBlocksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SymConformalBlocksRequest.Merge(m, src)
}
func (m *SymConformalBlocksRequest) XXX_Size() int {
	return xxx_messageInfo_SymConformalBlocksRequest.Size(m)
}
func (m *SymConformalBlocksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SymConformalBlocksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SymConformalBlocksRequest proto.InternalMessageInfo

func (m *SymConformalBlocksRequest) GetAlgebra() *LieAlgebra {
	if m != nil {
		return m.Algebra
	}
	return nil
}

func (m *SymConformalBlocksRequest) GetWeight() *Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *SymConformalBlocksRequest) GetNumPoints() int64 {
	if m != nil {
		return m.NumPoints
	}
	return 0
}

func (m *SymConformalBlocksRequest) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

// Represents a Lie algebra
type LieAlgebra struct {
	Type                 LieAlgebra_LieType `protobuf:"varint,1,opt,name=type,proto3,enum=cbservice.LieAlgebra_LieType" json:"type,omitempty"`
	Rank                 int64              `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LieAlgebra) Reset()         { *m = LieAlgebra{} }
func (m *LieAlgebra) String() string { return proto.CompactTextString(m) }
func (*LieAlgebra) ProtoMessage()    {}
func (*LieAlgebra) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67941f876b8a60b, []int{2}
}

func (m *LieAlgebra) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LieAlgebra.Unmarshal(m, b)
}
func (m *LieAlgebra) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LieAlgebra.Marshal(b, m, deterministic)
}
func (m *LieAlgebra) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LieAlgebra.Merge(m, src)
}
func (m *LieAlgebra) XXX_Size() int {
	return xxx_messageInfo_LieAlgebra.Size(m)
}
func (m *LieAlgebra) XXX_DiscardUnknown() {
	xxx_messageInfo_LieAlgebra.DiscardUnknown(m)
}

var xxx_messageInfo_LieAlgebra proto.InternalMessageInfo

func (m *LieAlgebra) GetType() LieAlgebra_LieType {
	if m != nil {
		return m.Type
	}
	return LieAlgebra_A
}

func (m *LieAlgebra) GetRank() int64 {
	if m != nil {
		return m.Rank
	}
	return 0
}

// An integral weight
type Weight struct {
	Coords               []int64  `protobuf:"varint,1,rep,packed,name=coords,proto3" json:"coords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Weight) Reset()         { *m = Weight{} }
func (m *Weight) String() string { return proto.CompactTextString(m) }
func (*Weight) ProtoMessage()    {}
func (*Weight) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67941f876b8a60b, []int{3}
}

func (m *Weight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Weight.Unmarshal(m, b)
}
func (m *Weight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Weight.Marshal(b, m, deterministic)
}
func (m *Weight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Weight.Merge(m, src)
}
func (m *Weight) XXX_Size() int {
	return xxx_messageInfo_Weight.Size(m)
}
func (m *Weight) XXX_DiscardUnknown() {
	xxx_messageInfo_Weight.DiscardUnknown(m)
}

var xxx_messageInfo_Weight proto.InternalMessageInfo

func (m *Weight) GetCoords() []int64 {
	if m != nil {
		return m.Coords
	}
	return nil
}

// A vector reply message
type VectorReply struct {
	Coords               []*RatReply `protobuf:"bytes,1,rep,name=coords,proto3" json:"coords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *VectorReply) Reset()         { *m = VectorReply{} }
func (m *VectorReply) String() string { return proto.CompactTextString(m) }
func (*VectorReply) ProtoMessage()    {}
func (*VectorReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67941f876b8a60b, []int{4}
}

func (m *VectorReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VectorReply.Unmarshal(m, b)
}
func (m *VectorReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VectorReply.Marshal(b, m, deterministic)
}
func (m *VectorReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VectorReply.Merge(m, src)
}
func (m *VectorReply) XXX_Size() int {
	return xxx_messageInfo_VectorReply.Size(m)
}
func (m *VectorReply) XXX_DiscardUnknown() {
	xxx_messageInfo_VectorReply.DiscardUnknown(m)
}

var xxx_messageInfo_VectorReply proto.InternalMessageInfo

func (m *VectorReply) GetCoords() []*RatReply {
	if m != nil {
		return m.Coords
	}
	return nil
}

// A rational reply message
type RatReply struct {
	Numerator            *IntReply `protobuf:"bytes,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	Denominator          *IntReply `protobuf:"bytes,2,opt,name=denominator,proto3" json:"denominator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RatReply) Reset()         { *m = RatReply{} }
func (m *RatReply) String() string { return proto.CompactTextString(m) }
func (*RatReply) ProtoMessage()    {}
func (*RatReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67941f876b8a60b, []int{5}
}

func (m *RatReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RatReply.Unmarshal(m, b)
}
func (m *RatReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RatReply.Marshal(b, m, deterministic)
}
func (m *RatReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RatReply.Merge(m, src)
}
func (m *RatReply) XXX_Size() int {
	return xxx_messageInfo_RatReply.Size(m)
}
func (m *RatReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RatReply.DiscardUnknown(m)
}

var xxx_messageInfo_RatReply proto.InternalMessageInfo

func (m *RatReply) GetNumerator() *IntReply {
	if m != nil {
		return m.Numerator
	}
	return nil
}

func (m *RatReply) GetDenominator() *IntReply {
	if m != nil {
		return m.Denominator
	}
	return nil
}

// An integer reply message
type IntReply struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	BigResult            string   `protobuf:"bytes,2,opt,name=big_result,json=bigResult,proto3" json:"big_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntReply) Reset()         { *m = IntReply{} }
func (m *IntReply) String() string { return proto.CompactTextString(m) }
func (*IntReply) ProtoMessage()    {}
func (*IntReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c67941f876b8a60b, []int{6}
}

func (m *IntReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntReply.Unmarshal(m, b)
}
func (m *IntReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntReply.Marshal(b, m, deterministic)
}
func (m *IntReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntReply.Merge(m, src)
}
func (m *IntReply) XXX_Size() int {
	return xxx_messageInfo_IntReply.Size(m)
}
func (m *IntReply) XXX_DiscardUnknown() {
	xxx_messageInfo_IntReply.DiscardUnknown(m)
}

var xxx_messageInfo_IntReply proto.InternalMessageInfo

func (m *IntReply) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *IntReply) GetBigResult() string {
	if m != nil {
		return m.BigResult
	}
	return ""
}

func init() {
	proto.RegisterEnum("cbservice.LieAlgebra_LieType", LieAlgebra_LieType_name, LieAlgebra_LieType_value)
	proto.RegisterType((*ConformalBlocksRequest)(nil), "cbservice.ConformalBlocksRequest")
	proto.RegisterType((*SymConformalBlocksRequest)(nil), "cbservice.SymConformalBlocksRequest")
	proto.RegisterType((*LieAlgebra)(nil), "cbservice.LieAlgebra")
	proto.RegisterType((*Weight)(nil), "cbservice.Weight")
	proto.RegisterType((*VectorReply)(nil), "cbservice.VectorReply")
	proto.RegisterType((*RatReply)(nil), "cbservice.RatReply")
	proto.RegisterType((*IntReply)(nil), "cbservice.IntReply")
}

func init() { proto.RegisterFile("cblocks.proto", fileDescriptor_c67941f876b8a60b) }

var fileDescriptor_c67941f876b8a60b = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x51, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0xbb, 0x4d, 0xaf, 0xbd, 0x4e, 0xf0, 0xe8, 0xad, 0x5a, 0xa2, 0x70, 0x10, 0x83, 0x42,
	0xe5, 0xa0, 0x72, 0x11, 0x5f, 0x7c, 0xeb, 0xf5, 0x40, 0x04, 0x05, 0xd9, 0x13, 0x7d, 0x3c, 0x92,
	0xdc, 0x58, 0x97, 0x26, 0xbb, 0x71, 0xb3, 0xa9, 0xd4, 0xcf, 0xe0, 0x47, 0xf1, 0xf3, 0xf9, 0x2c,
	0xd9, 0x6c, 0x2f, 0x2b, 0xb4, 0xe0, 0x83, 0x4f, 0x9b, 0x9d, 0xf9, 0xfd, 0x67, 0xfe, 0xbb, 0x3b,
	0x81, 0x7b, 0x59, 0x9a, 0xcb, 0x6c, 0x5d, 0xcd, 0x4b, 0x25, 0xb5, 0xa4, 0xe3, 0x2c, 0xad, 0x50,
	0x6d, 0x78, 0x86, 0xd1, 0x4f, 0x02, 0xd3, 0xa5, 0x14, 0x5f, 0xa4, 0x2a, 0x92, 0xfc, 0xd2, 0x40,
	0x0c, 0xbf, 0xd5, 0x58, 0x69, 0xfa, 0x02, 0x46, 0x49, 0xbe, 0xc2, 0x54, 0x25, 0x01, 0x09, 0xc9,
	0xcc, 0x8f, 0x1f, 0xce, 0xef, 0x74, 0xf3, 0x77, 0x1c, 0x17, 0x6d, 0x92, 0xed, 0x28, 0x7a, 0x0e,
	0xa3, 0xef, 0xc8, 0x57, 0x5f, 0x75, 0x15, 0xf4, 0x43, 0x6f, 0xe6, 0xc7, 0xa7, 0x8e, 0xe0, 0xb3,
	0xc9, 0xb0, 0x1d, 0x41, 0x1f, 0xc0, 0x51, 0x8e, 0x1b, 0xcc, 0x03, 0x2f, 0x24, 0x33, 0x8f, 0xb5,
	0x9b, 0xe8, 0x17, 0x81, 0x47, 0xd7, 0xdb, 0xe2, 0x7f, 0x39, 0x7a, 0x0e, 0xc3, 0xb6, 0x5f, 0xd0,
	0x37, 0xfc, 0x1e, 0x43, 0x16, 0xa0, 0x67, 0x00, 0xa2, 0x2e, 0x6e, 0x4a, 0xc9, 0x85, 0xae, 0xac,
	0xa9, 0xb1, 0xa8, 0x8b, 0x0f, 0x26, 0xd0, 0xd9, 0x1d, 0xb8, 0x76, 0x7f, 0x00, 0x74, 0x6d, 0xe9,
	0x05, 0x0c, 0xf4, 0xb6, 0x44, 0xe3, 0xed, 0x24, 0x3e, 0xdb, 0xeb, 0xad, 0xf9, 0xfc, 0xb8, 0x2d,
	0x91, 0x19, 0x94, 0x52, 0x18, 0xa8, 0x44, 0xac, 0x8d, 0x3d, 0x8f, 0x99, 0xef, 0xe8, 0x19, 0x8c,
	0x2c, 0x44, 0x8f, 0x80, 0x2c, 0x26, 0xbd, 0x66, 0xb9, 0x9c, 0x90, 0x66, 0x59, 0x4e, 0xfa, 0xcd,
	0x72, 0x35, 0xf1, 0xa2, 0x10, 0x86, 0xed, 0x11, 0xe8, 0x14, 0x86, 0x99, 0x94, 0xea, 0xb6, 0x0a,
	0x48, 0xe8, 0xcd, 0x3c, 0x66, 0x77, 0xd1, 0x6b, 0xf0, 0x3f, 0x61, 0xa6, 0xa5, 0x62, 0x58, 0xe6,
	0x5b, 0x7a, 0xfe, 0x17, 0xe6, 0xc7, 0xf7, 0x1d, 0x83, 0x2c, 0xd1, 0x06, 0xba, 0xd3, 0x6a, 0x38,
	0xde, 0xc5, 0xe8, 0x05, 0x34, 0x17, 0x81, 0x2a, 0xd1, 0x52, 0xd9, 0x8b, 0x77, 0xb5, 0x6f, 0x85,
	0xd5, 0x76, 0x14, 0x7d, 0x05, 0xfe, 0x2d, 0x0a, 0x59, 0x70, 0x61, 0x44, 0xfd, 0xc3, 0x22, 0x97,
	0x8b, 0x16, 0x70, 0xbc, 0x4b, 0x34, 0xa7, 0x52, 0x58, 0xd5, 0xb9, 0x36, 0x2d, 0x3d, 0x66, 0x77,
	0xcd, 0x43, 0xa5, 0x7c, 0x75, 0x63, 0x73, 0x4d, 0xe5, 0x31, 0x1b, 0xa7, 0x7c, 0xc5, 0x4c, 0x20,
	0xfe, 0x4d, 0x60, 0xb4, 0x6c, 0xc7, 0x86, 0xbe, 0x01, 0x7f, 0x29, 0x8b, 0xb2, 0xd6, 0xc8, 0x12,
	0xb1, 0xa6, 0x4f, 0x9c, 0xfe, 0xfb, 0x27, 0xec, 0xf1, 0x3e, 0x8b, 0x51, 0x8f, 0xbe, 0x87, 0x13,
	0x33, 0x95, 0x5d, 0xad, 0xa7, 0x0e, 0x78, 0x70, 0x60, 0x0f, 0x95, 0xbb, 0x86, 0xd3, 0xae, 0xdc,
	0x15, 0xdf, 0xf0, 0x4a, 0xaa, 0x7f, 0xac, 0x38, 0x75, 0x28, 0xe7, 0x71, 0xa3, 0x5e, 0x3a, 0x34,
	0xff, 0xf6, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x58, 0x8a, 0x1b, 0x45, 0xec, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CBlocksClient is the client API for CBlocks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CBlocksClient interface {
	// Computes the rank of a symmetric conformal blocks bundle
	ComputeRank(ctx context.Context, in *ConformalBlocksRequest, opts ...grpc.CallOption) (*IntReply, error)
	// Computes the rank of a symmetric conformal blocks bundle
	SymComputeRank(ctx context.Context, in *SymConformalBlocksRequest, opts ...grpc.CallOption) (*IntReply, error)
	// Computes the divisor of the symmetric conformal blocks bundle
	SymComputeDivisor(ctx context.Context, in *SymConformalBlocksRequest, opts ...grpc.CallOption) (*VectorReply, error)
}

type cBlocksClient struct {
	cc *grpc.ClientConn
}

func NewCBlocksClient(cc *grpc.ClientConn) CBlocksClient {
	return &cBlocksClient{cc}
}

func (c *cBlocksClient) ComputeRank(ctx context.Context, in *ConformalBlocksRequest, opts ...grpc.CallOption) (*IntReply, error) {
	out := new(IntReply)
	err := c.cc.Invoke(ctx, "/cbservice.CBlocks/ComputeRank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cBlocksClient) SymComputeRank(ctx context.Context, in *SymConformalBlocksRequest, opts ...grpc.CallOption) (*IntReply, error) {
	out := new(IntReply)
	err := c.cc.Invoke(ctx, "/cbservice.CBlocks/SymComputeRank", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cBlocksClient) SymComputeDivisor(ctx context.Context, in *SymConformalBlocksRequest, opts ...grpc.CallOption) (*VectorReply, error) {
	out := new(VectorReply)
	err := c.cc.Invoke(ctx, "/cbservice.CBlocks/SymComputeDivisor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CBlocksServer is the server API for CBlocks service.
type CBlocksServer interface {
	// Computes the rank of a symmetric conformal blocks bundle
	ComputeRank(context.Context, *ConformalBlocksRequest) (*IntReply, error)
	// Computes the rank of a symmetric conformal blocks bundle
	SymComputeRank(context.Context, *SymConformalBlocksRequest) (*IntReply, error)
	// Computes the divisor of the symmetric conformal blocks bundle
	SymComputeDivisor(context.Context, *SymConformalBlocksRequest) (*VectorReply, error)
}

// UnimplementedCBlocksServer can be embedded to have forward compatible implementations.
type UnimplementedCBlocksServer struct {
}

func (*UnimplementedCBlocksServer) ComputeRank(ctx context.Context, req *ConformalBlocksRequest) (*IntReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeRank not implemented")
}
func (*UnimplementedCBlocksServer) SymComputeRank(ctx context.Context, req *SymConformalBlocksRequest) (*IntReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SymComputeRank not implemented")
}
func (*UnimplementedCBlocksServer) SymComputeDivisor(ctx context.Context, req *SymConformalBlocksRequest) (*VectorReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SymComputeDivisor not implemented")
}

func RegisterCBlocksServer(s *grpc.Server, srv CBlocksServer) {
	s.RegisterService(&_CBlocks_serviceDesc, srv)
}

func _CBlocks_ComputeRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConformalBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CBlocksServer).ComputeRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbservice.CBlocks/ComputeRank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CBlocksServer).ComputeRank(ctx, req.(*ConformalBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CBlocks_SymComputeRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymConformalBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CBlocksServer).SymComputeRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbservice.CBlocks/SymComputeRank",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CBlocksServer).SymComputeRank(ctx, req.(*SymConformalBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CBlocks_SymComputeDivisor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymConformalBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CBlocksServer).SymComputeDivisor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbservice.CBlocks/SymComputeDivisor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CBlocksServer).SymComputeDivisor(ctx, req.(*SymConformalBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CBlocks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cbservice.CBlocks",
	HandlerType: (*CBlocksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeRank",
			Handler:    _CBlocks_ComputeRank_Handler,
		},
		{
			MethodName: "SymComputeRank",
			Handler:    _CBlocks_SymComputeRank_Handler,
		},
		{
			MethodName: "SymComputeDivisor",
			Handler:    _CBlocks_SymComputeDivisor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cblocks.proto",
}
