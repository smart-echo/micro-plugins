// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/broker.proto

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
	Broker_Publish_FullMethodName = "/Broker/Publish"
)

// BrokerClient is the client API for Broker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerClient interface {
	Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error)
}

type brokerClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerClient(cc grpc.ClientConnInterface) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Broker_Publish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServer is the server API for Broker service.
// All implementations must embed UnimplementedBrokerServer
// for forward compatibility.
type BrokerServer interface {
	Publish(context.Context, *Message) (*Empty, error)
	mustEmbedUnimplementedBrokerServer()
}

// UnimplementedBrokerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBrokerServer struct{}

func (UnimplementedBrokerServer) Publish(context.Context, *Message) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedBrokerServer) mustEmbedUnimplementedBrokerServer() {}
func (UnimplementedBrokerServer) testEmbeddedByValue()                {}

// UnsafeBrokerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerServer will
// result in compilation errors.
type UnsafeBrokerServer interface {
	mustEmbedUnimplementedBrokerServer()
}

func RegisterBrokerServer(s grpc.ServiceRegistrar, srv BrokerServer) {
	// If the following call pancis, it indicates UnimplementedBrokerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Broker_ServiceDesc, srv)
}

func _Broker_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Broker_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Publish(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// Broker_ServiceDesc is the grpc.ServiceDesc for Broker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Broker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Broker_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/broker.proto",
}
