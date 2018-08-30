// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/jinbanglin/examples/grpc/greeter/srv/proto/hello/hello.proto

/*
Package go_micro_srv_greeter is a generated protocol buffer package.

It is generated from these files:
	github.com/jinbanglin/examples/grpc/greeter/srv/proto/hello/hello.proto

It has these top-level messages:
	Request
	Response
*/
package go_micro_srv_greeter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.greeter.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.greeter.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Say service

type SayClient interface {
	Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type sayClient struct {
	cc *grpc.ClientConn
}

func NewSayClient(cc *grpc.ClientConn) SayClient {
	return &sayClient{cc}
}

func (c *sayClient) Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/go.micro.srv.greeter.Say/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Say service

type SayServer interface {
	Hello(context.Context, *Request) (*Response, error)
}

func RegisterSayServer(s *grpc.Server, srv SayServer) {
	s.RegisterService(&_Say_serviceDesc, srv)
}

func _Say_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.greeter.Say/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayServer).Hello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Say_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.greeter.Say",
	HandlerType: (*SayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Say_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/jinbanglin/examples/grpc/greeter/srv/proto/hello/hello.proto",
}

func init() {
	proto.RegisterFile("github.com/jinbanglin/examples/grpc/greeter/srv/proto/hello/hello.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0xcd, 0xaa, 0xc2, 0x30,
	0x10, 0x85, 0x6f, 0xe9, 0xfd, 0x33, 0x2b, 0x09, 0x2e, 0x44, 0xac, 0x48, 0x57, 0xae, 0x32, 0xa0,
	0x6f, 0xe0, 0xaa, 0x3b, 0xa1, 0x3e, 0x41, 0x5b, 0x86, 0xb4, 0xd0, 0x34, 0x71, 0x26, 0x2d, 0xfa,
	0xf6, 0xd2, 0x98, 0xa5, 0x6e, 0x86, 0xc3, 0x1c, 0x3e, 0xbe, 0x23, 0xce, 0xba, 0xf3, 0xed, 0x58,
	0xab, 0xc6, 0x1a, 0x30, 0x5d, 0x43, 0x16, 0xf0, 0x5e, 0x19, 0xd7, 0x23, 0x83, 0x26, 0xd7, 0x80,
	0x26, 0x44, 0x8f, 0x04, 0x4c, 0x13, 0x38, 0xb2, 0xde, 0x42, 0x8b, 0x7d, 0x1f, 0xaf, 0x0a, 0x1f,
	0xb9, 0xd2, 0x56, 0x05, 0x56, 0x31, 0x4d, 0x2a, 0x12, 0x79, 0x26, 0xfe, 0x4a, 0xbc, 0x8d, 0xc8,
	0x5e, 0x4a, 0xf1, 0x3d, 0x54, 0x06, 0xd7, 0xc9, 0x3e, 0x39, 0x2c, 0xca, 0x90, 0xf3, 0xad, 0xf8,
	0x2f, 0x91, 0x9d, 0x1d, 0x18, 0xe5, 0x52, 0xa4, 0x86, 0x75, 0xac, 0xe7, 0x78, 0xbc, 0x88, 0xf4,
	0x5a, 0x3d, 0x64, 0x21, 0x7e, 0x8a, 0x59, 0x24, 0x33, 0xf5, 0xce, 0xa1, 0xa2, 0x60, 0xb3, 0xfb,
	0x54, 0xbf, 0x04, 0xf9, 0x57, 0xfd, 0x1b, 0xa6, 0x9e, 0x9e, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc1,
	0xe5, 0xed, 0x58, 0xf0, 0x00, 0x00, 0x00,
}
