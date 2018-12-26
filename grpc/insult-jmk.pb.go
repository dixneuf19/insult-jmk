// Code generated by protoc-gen-go. DO NOT EDIT.
// source: insult-jmk.proto

package insultJMK

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type InsultRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsultRequest) Reset()         { *m = InsultRequest{} }
func (m *InsultRequest) String() string { return proto.CompactTextString(m) }
func (*InsultRequest) ProtoMessage()    {}
func (*InsultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7175cb32aa7d3e7a, []int{0}
}

func (m *InsultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsultRequest.Unmarshal(m, b)
}
func (m *InsultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsultRequest.Marshal(b, m, deterministic)
}
func (m *InsultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsultRequest.Merge(m, src)
}
func (m *InsultRequest) XXX_Size() int {
	return xxx_messageInfo_InsultRequest.Size(m)
}
func (m *InsultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InsultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InsultRequest proto.InternalMessageInfo

func (m *InsultRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the insult, with user name
type InsultResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsultResponse) Reset()         { *m = InsultResponse{} }
func (m *InsultResponse) String() string { return proto.CompactTextString(m) }
func (*InsultResponse) ProtoMessage()    {}
func (*InsultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7175cb32aa7d3e7a, []int{1}
}

func (m *InsultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsultResponse.Unmarshal(m, b)
}
func (m *InsultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsultResponse.Marshal(b, m, deterministic)
}
func (m *InsultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsultResponse.Merge(m, src)
}
func (m *InsultResponse) XXX_Size() int {
	return xxx_messageInfo_InsultResponse.Size(m)
}
func (m *InsultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InsultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InsultResponse proto.InternalMessageInfo

func (m *InsultResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*InsultRequest)(nil), "insultJMK.InsultRequest")
	proto.RegisterType((*InsultResponse)(nil), "insultJMK.InsultResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InsulterClient is the client API for Insulter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InsulterClient interface {
	// Sends a greeting
	GetInsult(ctx context.Context, in *InsultRequest, opts ...grpc.CallOption) (*InsultResponse, error)
}

type insulterClient struct {
	cc *grpc.ClientConn
}

func NewInsulterClient(cc *grpc.ClientConn) InsulterClient {
	return &insulterClient{cc}
}

func (c *insulterClient) GetInsult(ctx context.Context, in *InsultRequest, opts ...grpc.CallOption) (*InsultResponse, error) {
	out := new(InsultResponse)
	err := c.cc.Invoke(ctx, "/insultJMK.Insulter/GetInsult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InsulterServer is the server API for Insulter service.
type InsulterServer interface {
	// Sends a greeting
	GetInsult(context.Context, *InsultRequest) (*InsultResponse, error)
}

func RegisterInsulterServer(s *grpc.Server, srv InsulterServer) {
	s.RegisterService(&_Insulter_serviceDesc, srv)
}

func _Insulter_GetInsult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InsulterServer).GetInsult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/insultJMK.Insulter/GetInsult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InsulterServer).GetInsult(ctx, req.(*InsultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Insulter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "insultJMK.Insulter",
	HandlerType: (*InsulterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInsult",
			Handler:    _Insulter_GetInsult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "insult-jmk.proto",
}

func init() { proto.RegisterFile("insult-jmk.proto", fileDescriptor_7175cb32aa7d3e7a) }

var fileDescriptor_7175cb32aa7d3e7a = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0xcc, 0x2b, 0x2e,
	0xcd, 0x29, 0xd1, 0xcd, 0xca, 0xcd, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x88,
	0x78, 0xf9, 0x7a, 0x2b, 0x29, 0x73, 0xf1, 0x7a, 0x82, 0x39, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x60, 0xb6, 0x92, 0x16, 0x17, 0x1f, 0x4c, 0x51, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x04,
	0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x3a, 0x4c, 0x21, 0x8c, 0x6b, 0xe4, 0xc7, 0xc5, 0x01,
	0x51, 0x9b, 0x5a, 0x24, 0xe4, 0xc4, 0xc5, 0xe9, 0x9e, 0x5a, 0x02, 0xe1, 0x0a, 0x49, 0xe8, 0xc1,
	0x6d, 0xd5, 0x43, 0xb1, 0x52, 0x4a, 0x12, 0x8b, 0x0c, 0xc4, 0x1e, 0x25, 0x86, 0x24, 0x36, 0xb0,
	0x93, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x4b, 0x40, 0x63, 0xc6, 0x00, 0x00, 0x00,
}