// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Transport_Stream_FullMethodName = "/go.micro.transport.grpc.Transport/Stream"
)

// TransportClient is the client API for Transport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransportClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Message, Message], error)
}

type transportClient struct {
	cc grpc.ClientConnInterface
}

func NewTransportClient(cc grpc.ClientConnInterface) TransportClient {
	return &transportClient{cc}
}

func (c *transportClient) Stream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Message, Message], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Transport_ServiceDesc.Streams[0], Transport_Stream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Message, Message]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Transport_StreamClient = grpc.BidiStreamingClient[Message, Message]

// TransportServer is the server API for Transport service.
// All implementations must embed UnimplementedTransportServer
// for forward compatibility.
type TransportServer interface {
	Stream(grpc.BidiStreamingServer[Message, Message]) error
	mustEmbedUnimplementedTransportServer()
}

// UnimplementedTransportServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransportServer struct{}

func (UnimplementedTransportServer) Stream(grpc.BidiStreamingServer[Message, Message]) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedTransportServer) mustEmbedUnimplementedTransportServer() {}
func (UnimplementedTransportServer) testEmbeddedByValue()                   {}

// UnsafeTransportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransportServer will
// result in compilation errors.
type UnsafeTransportServer interface {
	mustEmbedUnimplementedTransportServer()
}

func RegisterTransportServer(s grpc.ServiceRegistrar, srv TransportServer) {
	// If the following call pancis, it indicates UnimplementedTransportServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Transport_ServiceDesc, srv)
}

func _Transport_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TransportServer).Stream(&grpc.GenericServerStream[Message, Message]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Transport_StreamServer = grpc.BidiStreamingServer[Message, Message]

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
