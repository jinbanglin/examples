// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/jinbanglin/examples/api/default/proto/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	github.com/jinbanglin/examples/api/default/proto/api.proto

It has these top-level messages:
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_api "github.com/jinbanglin/go-api/proto"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Example service

type ExampleClient interface {
	Call(ctx context.Context, in *go_api.Request, opts ...grpc.CallOption) (*go_api.Response, error)
}

type exampleClient struct {
	cc *grpc.ClientConn
}

func NewExampleClient(cc *grpc.ClientConn) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Call(ctx context.Context, in *go_api.Request, opts ...grpc.CallOption) (*go_api.Response, error) {
	out := new(go_api.Response)
	err := grpc.Invoke(ctx, "/Example/Call", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Example service

type ExampleServer interface {
	Call(context.Context, *go_api.Request) (*go_api.Response, error)
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(go_api.Request)
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
		return srv.(ExampleServer).Call(ctx, req.(*go_api.Request))
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
	Metadata: "github.com/jinbanglin/examples/api/default/proto/api.proto",
}

// Client API for Foo service

type FooClient interface {
	Bar(ctx context.Context, in *go_api.Request, opts ...grpc.CallOption) (*go_api.Response, error)
}

type fooClient struct {
	cc *grpc.ClientConn
}

func NewFooClient(cc *grpc.ClientConn) FooClient {
	return &fooClient{cc}
}

func (c *fooClient) Bar(ctx context.Context, in *go_api.Request, opts ...grpc.CallOption) (*go_api.Response, error) {
	out := new(go_api.Response)
	err := grpc.Invoke(ctx, "/Foo/Bar", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Foo service

type FooServer interface {
	Bar(context.Context, *go_api.Request) (*go_api.Response, error)
}

func RegisterFooServer(s *grpc.Server, srv FooServer) {
	s.RegisterService(&_Foo_serviceDesc, srv)
}

func _Foo_Bar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(go_api.Request)
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
		return srv.(FooServer).Bar(ctx, req.(*go_api.Request))
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
	Metadata: "github.com/jinbanglin/examples/api/default/proto/api.proto",
}

func init() {
	proto.RegisterFile("github.com/jinbanglin/examples/api/default/proto/api.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0x4f, 0xad, 0x48,
	0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0x49, 0x4d, 0x4b, 0x2c, 0xcd,
	0x29, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x07, 0x89, 0xe8, 0x81, 0x59, 0x52, 0xea, 0x18, 0xda,
	0xd2, 0xf3, 0x75, 0x41, 0x1a, 0xd0, 0x14, 0x1a, 0x99, 0x71, 0xb1, 0xbb, 0x42, 0x0c, 0x14, 0xd2,
	0xe6, 0x62, 0x71, 0x4e, 0xcc, 0xc9, 0x11, 0xe2, 0xd7, 0x4b, 0xcf, 0xd7, 0x03, 0xa9, 0x08, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x91, 0x12, 0x40, 0x08, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x2a,
	0x31, 0x18, 0x19, 0x72, 0x31, 0xbb, 0xe5, 0xe7, 0x0b, 0x69, 0x71, 0x31, 0x3b, 0x25, 0x16, 0x11,
	0xa5, 0x25, 0x89, 0x0d, 0x6c, 0xa3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x16, 0x19, 0x17, 0x97,
	0xd3, 0x00, 0x00, 0x00,
}
