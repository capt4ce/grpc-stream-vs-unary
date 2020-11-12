// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MainServiceClient is the client API for MainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MainServiceClient interface {
	SendUnaryRequest(ctx context.Context, in *UnaryRequest, opts ...grpc.CallOption) (*UnaryReply, error)
}

type mainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMainServiceClient(cc grpc.ClientConnInterface) MainServiceClient {
	return &mainServiceClient{cc}
}

func (c *mainServiceClient) SendUnaryRequest(ctx context.Context, in *UnaryRequest, opts ...grpc.CallOption) (*UnaryReply, error) {
	out := new(UnaryReply)
	err := c.cc.Invoke(ctx, "/proto.MainService/SendUnaryRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainServiceServer is the server API for MainService service.
// All implementations must embed UnimplementedMainServiceServer
// for forward compatibility
type MainServiceServer interface {
	SendUnaryRequest(context.Context, *UnaryRequest) (*UnaryReply, error)
	mustEmbedUnimplementedMainServiceServer()
}

// UnimplementedMainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMainServiceServer struct {
}

func (UnimplementedMainServiceServer) SendUnaryRequest(context.Context, *UnaryRequest) (*UnaryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUnaryRequest not implemented")
}
func (UnimplementedMainServiceServer) mustEmbedUnimplementedMainServiceServer() {}

// UnsafeMainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MainServiceServer will
// result in compilation errors.
type UnsafeMainServiceServer interface {
	mustEmbedUnimplementedMainServiceServer()
}

func RegisterMainServiceServer(s grpc.ServiceRegistrar, srv MainServiceServer) {
	s.RegisterService(&_MainService_serviceDesc, srv)
}

func _MainService_SendUnaryRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).SendUnaryRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MainService/SendUnaryRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).SendUnaryRequest(ctx, req.(*UnaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MainService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MainService",
	HandlerType: (*MainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendUnaryRequest",
			Handler:    _MainService_SendUnaryRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamServiceClient interface {
	SendStreamRequest(ctx context.Context, opts ...grpc.CallOption) (StreamService_SendStreamRequestClient, error)
}

type streamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServiceClient(cc grpc.ClientConnInterface) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) SendStreamRequest(ctx context.Context, opts ...grpc.CallOption) (StreamService_SendStreamRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/proto.StreamService/SendStreamRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceSendStreamRequestClient{stream}
	return x, nil
}

type StreamService_SendStreamRequestClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamReply, error)
	grpc.ClientStream
}

type streamServiceSendStreamRequestClient struct {
	grpc.ClientStream
}

func (x *streamServiceSendStreamRequestClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceSendStreamRequestClient) Recv() (*StreamReply, error) {
	m := new(StreamReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
// All implementations must embed UnimplementedStreamServiceServer
// for forward compatibility
type StreamServiceServer interface {
	SendStreamRequest(StreamService_SendStreamRequestServer) error
	mustEmbedUnimplementedStreamServiceServer()
}

// UnimplementedStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (UnimplementedStreamServiceServer) SendStreamRequest(StreamService_SendStreamRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method SendStreamRequest not implemented")
}
func (UnimplementedStreamServiceServer) mustEmbedUnimplementedStreamServiceServer() {}

// UnsafeStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServiceServer will
// result in compilation errors.
type UnsafeStreamServiceServer interface {
	mustEmbedUnimplementedStreamServiceServer()
}

func RegisterStreamServiceServer(s grpc.ServiceRegistrar, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_SendStreamRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).SendStreamRequest(&streamServiceSendStreamRequestServer{stream})
}

type StreamService_SendStreamRequestServer interface {
	Send(*StreamReply) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamServiceSendStreamRequestServer struct {
	grpc.ServerStream
}

func (x *streamServiceSendStreamRequestServer) Send(m *StreamReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceSendStreamRequestServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendStreamRequest",
			Handler:       _StreamService_SendStreamRequest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "main.proto",
}
