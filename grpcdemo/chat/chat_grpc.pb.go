// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chat

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	Average(ctx context.Context, opts ...grpc.CallOption) (ChatService_AverageClient, error)
	Max(ctx context.Context, opts ...grpc.CallOption) (ChatService_MaxClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/chat.ChatService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (ChatService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/chat.ChatService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceAverageClient{stream}
	return x, nil
}

type ChatService_AverageClient interface {
	Send(*AverageMessage) error
	CloseAndRecv() (*AverageMessage, error)
	grpc.ClientStream
}

type chatServiceAverageClient struct {
	grpc.ClientStream
}

func (x *chatServiceAverageClient) Send(m *AverageMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceAverageClient) CloseAndRecv() (*AverageMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) Max(ctx context.Context, opts ...grpc.CallOption) (ChatService_MaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], "/chat.ChatService/Max", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceMaxClient{stream}
	return x, nil
}

type ChatService_MaxClient interface {
	Send(*MaxMessage) error
	Recv() (*MaxMessage, error)
	grpc.ClientStream
}

type chatServiceMaxClient struct {
	grpc.ClientStream
}

func (x *chatServiceMaxClient) Send(m *MaxMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceMaxClient) Recv() (*MaxMessage, error) {
	m := new(MaxMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
	Average(ChatService_AverageServer) error
	Max(ChatService_MaxServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) SayHello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedChatServiceServer) Average(ChatService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (UnimplementedChatServiceServer) Max(ChatService_MaxServer) error {
	return status.Errorf(codes.Unimplemented, "method Max not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Average(&chatServiceAverageServer{stream})
}

type ChatService_AverageServer interface {
	SendAndClose(*AverageMessage) error
	Recv() (*AverageMessage, error)
	grpc.ServerStream
}

type chatServiceAverageServer struct {
	grpc.ServerStream
}

func (x *chatServiceAverageServer) SendAndClose(m *AverageMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceAverageServer) Recv() (*AverageMessage, error) {
	m := new(AverageMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_Max_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Max(&chatServiceMaxServer{stream})
}

type ChatService_MaxServer interface {
	Send(*MaxMessage) error
	Recv() (*MaxMessage, error)
	grpc.ServerStream
}

type chatServiceMaxServer struct {
	grpc.ServerStream
}

func (x *chatServiceMaxServer) Send(m *MaxMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceMaxServer) Recv() (*MaxMessage, error) {
	m := new(MaxMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ChatService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Average",
			Handler:       _ChatService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Max",
			Handler:       _ChatService_Max_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat/chat.proto",
}