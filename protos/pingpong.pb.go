// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/pingpong.proto

package pingpong

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type PingRequest struct {
	Ping                 string   `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_523185c595eeb0c7, []int{0}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetPing() string {
	if m != nil {
		return m.Ping
	}
	return ""
}

type PingResponse struct {
	Pong                 string   `protobuf:"bytes,2,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_523185c595eeb0c7, []int{1}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "PingRequest")
	proto.RegisterType((*PingResponse)(nil), "PingResponse")
}

func init() { proto.RegisterFile("protos/pingpong.proto", fileDescriptor_523185c595eeb0c7) }

var fileDescriptor_523185c595eeb0c7 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0xc8, 0xcc, 0x4b, 0x2f, 0xc8, 0xcf, 0x4b, 0xd7, 0x03, 0xf3, 0xa5, 0x64,
	0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b,
	0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xb2, 0x4a, 0x8a, 0x5c, 0xdc, 0x01, 0x99, 0x79, 0xe9,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x20, 0xed, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x12, 0x17, 0x0f, 0x44, 0x49, 0x71, 0x41, 0x7e, 0x5e,
	0x71, 0x2a, 0x58, 0x4d, 0x7e, 0x5e, 0xba, 0x04, 0x13, 0x54, 0x4d, 0x7e, 0x5e, 0xba, 0x91, 0x23,
	0x17, 0x07, 0x48, 0x4d, 0x40, 0x7e, 0x5e, 0xba, 0x90, 0x29, 0x17, 0x0b, 0x88, 0x2d, 0xc4, 0xa3,
	0x87, 0x64, 0xb2, 0x14, 0xaf, 0x1e, 0xb2, 0x21, 0x4a, 0x02, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0xe2,
	0x52, 0x62, 0x05, 0x3b, 0xd7, 0x8a, 0x51, 0x2b, 0x89, 0x0d, 0xec, 0x20, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x4c, 0x11, 0x1b, 0x64, 0xc7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingPongClient is the client API for PingPong service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingPongClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type pingPongClient struct {
	cc *grpc.ClientConn
}

func NewPingPongClient(cc *grpc.ClientConn) PingPongClient {
	return &pingPongClient{cc}
}

func (c *pingPongClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/PingPong/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingPongServer is the server API for PingPong service.
type PingPongServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

// UnimplementedPingPongServer can be embedded to have forward compatible implementations.
type UnimplementedPingPongServer struct {
}

func (*UnimplementedPingPongServer) Ping(ctx context.Context, req *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterPingPongServer(s *grpc.Server, srv PingPongServer) {
	s.RegisterService(&_PingPong_serviceDesc, srv)
}

func _PingPong_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PingPong/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingPong_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PingPong",
	HandlerType: (*PingPongServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PingPong_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/pingpong.proto",
}