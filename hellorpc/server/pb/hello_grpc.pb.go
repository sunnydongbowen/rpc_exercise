// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: pb/hello.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloClient interface {
	SayHello(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
	// 服务端流式rpc  流式数据
	ServerStreamHello(ctx context.Context, in *Req, opts ...grpc.CallOption) (Hello_ServerStreamHelloClient, error)
	// 客户端流式rpc
	ClientStreamHello(ctx context.Context, opts ...grpc.CallOption) (Hello_ClientStreamHelloClient, error)
	// 双向rpc
	BudiStreamHello(ctx context.Context, opts ...grpc.CallOption) (Hello_BudiStreamHelloClient, error)
}

type helloClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloClient(cc grpc.ClientConnInterface) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) SayHello(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/pb.hello/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloClient) ServerStreamHello(ctx context.Context, in *Req, opts ...grpc.CallOption) (Hello_ServerStreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &Hello_ServiceDesc.Streams[0], "/pb.hello/ServerStreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServerStreamHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hello_ServerStreamHelloClient interface {
	Recv() (*Res, error)
	grpc.ClientStream
}

type helloServerStreamHelloClient struct {
	grpc.ClientStream
}

func (x *helloServerStreamHelloClient) Recv() (*Res, error) {
	m := new(Res)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloClient) ClientStreamHello(ctx context.Context, opts ...grpc.CallOption) (Hello_ClientStreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &Hello_ServiceDesc.Streams[1], "/pb.hello/ClientStreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloClientStreamHelloClient{stream}
	return x, nil
}

type Hello_ClientStreamHelloClient interface {
	Send(*Req) error
	CloseAndRecv() (*Res, error)
	grpc.ClientStream
}

type helloClientStreamHelloClient struct {
	grpc.ClientStream
}

func (x *helloClientStreamHelloClient) Send(m *Req) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloClientStreamHelloClient) CloseAndRecv() (*Res, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Res)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloClient) BudiStreamHello(ctx context.Context, opts ...grpc.CallOption) (Hello_BudiStreamHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &Hello_ServiceDesc.Streams[2], "/pb.hello/BudiStreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloBudiStreamHelloClient{stream}
	return x, nil
}

type Hello_BudiStreamHelloClient interface {
	Send(*Req) error
	Recv() (*Res, error)
	grpc.ClientStream
}

type helloBudiStreamHelloClient struct {
	grpc.ClientStream
}

func (x *helloBudiStreamHelloClient) Send(m *Req) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloBudiStreamHelloClient) Recv() (*Res, error) {
	m := new(Res)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServer is the server API for Hello service.
// All implementations must embed UnimplementedHelloServer
// for forward compatibility
type HelloServer interface {
	SayHello(context.Context, *Req) (*Res, error)
	// 服务端流式rpc  流式数据
	ServerStreamHello(*Req, Hello_ServerStreamHelloServer) error
	// 客户端流式rpc
	ClientStreamHello(Hello_ClientStreamHelloServer) error
	// 双向rpc
	BudiStreamHello(Hello_BudiStreamHelloServer) error
	mustEmbedUnimplementedHelloServer()
}

// UnimplementedHelloServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (UnimplementedHelloServer) SayHello(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedHelloServer) ServerStreamHello(*Req, Hello_ServerStreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamHello not implemented")
}
func (UnimplementedHelloServer) ClientStreamHello(Hello_ClientStreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamHello not implemented")
}
func (UnimplementedHelloServer) BudiStreamHello(Hello_BudiStreamHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method BudiStreamHello not implemented")
}
func (UnimplementedHelloServer) mustEmbedUnimplementedHelloServer() {}

// UnsafeHelloServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServer will
// result in compilation errors.
type UnsafeHelloServer interface {
	mustEmbedUnimplementedHelloServer()
}

func RegisterHelloServer(s grpc.ServiceRegistrar, srv HelloServer) {
	s.RegisterService(&Hello_ServiceDesc, srv)
}

func _Hello_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.hello/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).SayHello(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hello_ServerStreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Req)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServer).ServerStreamHello(m, &helloServerStreamHelloServer{stream})
}

type Hello_ServerStreamHelloServer interface {
	Send(*Res) error
	grpc.ServerStream
}

type helloServerStreamHelloServer struct {
	grpc.ServerStream
}

func (x *helloServerStreamHelloServer) Send(m *Res) error {
	return x.ServerStream.SendMsg(m)
}

func _Hello_ClientStreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServer).ClientStreamHello(&helloClientStreamHelloServer{stream})
}

type Hello_ClientStreamHelloServer interface {
	SendAndClose(*Res) error
	Recv() (*Req, error)
	grpc.ServerStream
}

type helloClientStreamHelloServer struct {
	grpc.ServerStream
}

func (x *helloClientStreamHelloServer) SendAndClose(m *Res) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloClientStreamHelloServer) Recv() (*Req, error) {
	m := new(Req)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Hello_BudiStreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServer).BudiStreamHello(&helloBudiStreamHelloServer{stream})
}

type Hello_BudiStreamHelloServer interface {
	Send(*Res) error
	Recv() (*Req, error)
	grpc.ServerStream
}

type helloBudiStreamHelloServer struct {
	grpc.ServerStream
}

func (x *helloBudiStreamHelloServer) Send(m *Res) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloBudiStreamHelloServer) Recv() (*Req, error) {
	m := new(Req)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Hello_ServiceDesc is the grpc.ServiceDesc for Hello service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hello_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Hello_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamHello",
			Handler:       _Hello_ServerStreamHello_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamHello",
			Handler:       _Hello_ClientStreamHello_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BudiStreamHello",
			Handler:       _Hello_BudiStreamHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/hello.proto",
}