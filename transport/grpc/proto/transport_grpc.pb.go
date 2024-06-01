// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/transport.proto

package proto

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

const (
	Transport_Stream_FullMethodName = "/go.micro.transport.grpc.Transport/Stream"
)

// TransportClient is the client API for Transport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransportClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (Transport_StreamClient, error)
}

type transportClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportClient(cc grpc.ClientConnInterface) TransportClient {
	return &transportClient{cc}
}

func (c *transportClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Transport_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Transport_ServiceDesc.Streams[0], Transport_Stream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &transportStreamClient{stream}
	return x, nil
}

type Transport_StreamClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type transportStreamClient struct {
	grpc.ClientStream
}

func (x *transportStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *transportStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TransportServer is the server API for Transport service.
// All implementations must embed UnimplementedTransportServer
// for forward compatibility
type TransportServer interface {
	Stream(Transport_StreamServer) error
	mustEmbedUnimplementedTransportServer()
}

// UnimplementedTransportServer must be embedded to have forward compatible implementations.
type UnimplementedTransportServer struct {
}

func (UnimplementedTransportServer) Stream(Transport_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedTransportServer) mustEmbedUnimplementedTransportServer() {}

// UnsafeTransportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransportServer will
// result in compilation errors.
type UnsafeTransportServer interface {
	mustEmbedUnimplementedTransportServer()
}

func RegisterTransportServer(s grpc.ServiceRegistrar, srv TransportServer) {
	s.RegisterService(&Transport_ServiceDesc, srv)
}

func _Transport_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransportServer).Stream(&transportStreamServer{stream})
}

type Transport_StreamServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type transportStreamServer struct {
	grpc.ServerStream
}

func (x *transportStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *transportStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Transport_ServiceDesc is the grpc.ServiceDesc for Transport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.transport.grpc.Transport",
	HandlerType: (*TransportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Transport_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/transport.proto",
}
