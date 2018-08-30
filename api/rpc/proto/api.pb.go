// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/jinbanglin/examples/api/rpc/proto/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	github.com/jinbanglin/examples/api/rpc/proto/api.proto

It has these top-level messages:
	CallRequest
	CallResponse
	EmptyRequest
	EmptyResponse
*/
package api

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

type CallRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CallRequest) Reset()                    { *m = CallRequest{} }
func (m *CallRequest) String() string            { return proto.CompactTextString(m) }
func (*CallRequest) ProtoMessage()               {}
func (*CallRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CallRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CallResponse struct {
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *CallResponse) Reset()                    { *m = CallResponse{} }
func (m *CallResponse) String() string            { return proto.CompactTextString(m) }
func (*CallResponse) ProtoMessage()               {}
func (*CallResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CallResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EmptyRequest struct {
}

func (m *EmptyRequest) Reset()                    { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string            { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()               {}
func (*EmptyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type EmptyResponse struct {
}

func (m *EmptyResponse) Reset()                    { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string            { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()               {}
func (*EmptyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*CallRequest)(nil), "CallRequest")
	proto.RegisterType((*CallResponse)(nil), "CallResponse")
	proto.RegisterType((*EmptyRequest)(nil), "EmptyRequest")
	proto.RegisterType((*EmptyResponse)(nil), "EmptyResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Example service

type ExampleClient interface {
	Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
}

type exampleClient struct {
	cc *grpc.ClientConn
}

func NewExampleClient(cc *grpc.ClientConn) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := grpc.Invoke(ctx, "/Example/Call", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Example service

type ExampleServer interface {
	Call(context.Context, *CallRequest) (*CallResponse, error)
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Example/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Call(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Example_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/jinbanglin/examples/api/rpc/proto/api.proto",
}

// Client API for Foo service

type FooClient interface {
	Bar(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type fooClient struct {
	cc *grpc.ClientConn
}

func NewFooClient(cc *grpc.ClientConn) FooClient {
	return &fooClient{cc}
}

func (c *fooClient) Bar(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := grpc.Invoke(ctx, "/Foo/Bar", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Foo service

type FooServer interface {
	Bar(context.Context, *EmptyRequest) (*EmptyResponse, error)
}

func RegisterFooServer(s *grpc.Server, srv FooServer) {
	s.RegisterService(&_Foo_serviceDesc, srv)
}

func _Foo_Bar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServer).Bar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Foo/Bar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServer).Bar(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Foo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Foo",
	HandlerType: (*FooServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bar",
			Handler:    _Foo_Bar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/jinbanglin/examples/api/rpc/proto/api.proto",
}

func init() { proto.RegisterFile("github.com/jinbanglin/examples/api/rpc/proto/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x51, 0x6a, 0x84, 0x30,
	0x14, 0x45, 0xb5, 0x4a, 0xa5, 0xaf, 0x6a, 0x21, 0x5f, 0xe2, 0x57, 0x1b, 0x68, 0xf1, 0xa7, 0x49,
	0x6b, 0x77, 0xd0, 0x62, 0x17, 0xe0, 0x0e, 0xa2, 0x3c, 0xac, 0x60, 0x4c, 0x9a, 0x44, 0x98, 0xd9,
	0xfd, 0x60, 0x74, 0xc0, 0xf9, 0xbb, 0x07, 0x6e, 0xce, 0xcd, 0x83, 0xcf, 0x61, 0x74, 0x7f, 0x4b,
	0xc7, 0x7a, 0x25, 0xb9, 0x1c, 0x7b, 0xa3, 0x38, 0x9e, 0x84, 0xd4, 0x13, 0x5a, 0x2e, 0xf4, 0xc8,
	0x8d, 0xee, 0xb9, 0x36, 0xca, 0xa9, 0x95, 0x98, 0x4f, 0xf4, 0x05, 0x1e, 0x7f, 0xc4, 0x34, 0xb5,
	0xf8, 0xbf, 0xa0, 0x75, 0x84, 0x40, 0x3c, 0x0b, 0x89, 0x45, 0xf8, 0x1c, 0x56, 0x0f, 0xad, 0xcf,
	0xb4, 0x82, 0x74, 0xab, 0x58, 0xad, 0x66, 0x8b, 0xa4, 0x80, 0x44, 0xa2, 0xb5, 0x62, 0xc0, 0xe2,
	0xce, 0xd7, 0xae, 0x48, 0x73, 0x48, 0x1b, 0xa9, 0xdd, 0x79, 0xb7, 0xd1, 0x27, 0xc8, 0x76, 0xde,
	0x9e, 0xd6, 0x1f, 0x90, 0x34, 0xdb, 0x8f, 0xc8, 0x2b, 0xc4, 0xab, 0x95, 0xa4, 0xec, 0xb0, 0x5f,
	0x66, 0xec, 0x38, 0x45, 0x83, 0xfa, 0x1d, 0xa2, 0x5f, 0xa5, 0xc8, 0x1b, 0x44, 0xdf, 0xc2, 0x90,
	0x8c, 0x1d, 0xfd, 0x65, 0xce, 0x6e, 0xf4, 0x34, 0xe8, 0xee, 0xfd, 0x55, 0x5f, 0x97, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x81, 0x2d, 0x74, 0x3d, 0x0a, 0x01, 0x00, 0x00,
}
