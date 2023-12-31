// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: brood.proto

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

const (
	Brood_Publish_FullMethodName   = "/pb.brood/Publish"
	Brood_Subscribe_FullMethodName = "/pb.brood/Subscribe"
)

// BroodClient is the client API for Brood service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BroodClient interface {
	Publish(ctx context.Context, in *Spore, opts ...grpc.CallOption) (*Acknowledgement, error)
	Subscribe(ctx context.Context, in *Hive, opts ...grpc.CallOption) (Brood_SubscribeClient, error)
}

type broodClient struct {
	cc grpc.ClientConnInterface
}

func NewBroodClient(cc grpc.ClientConnInterface) BroodClient {
	return &broodClient{cc}
}

func (c *broodClient) Publish(ctx context.Context, in *Spore, opts ...grpc.CallOption) (*Acknowledgement, error) {
	out := new(Acknowledgement)
	err := c.cc.Invoke(ctx, Brood_Publish_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broodClient) Subscribe(ctx context.Context, in *Hive, opts ...grpc.CallOption) (Brood_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Brood_ServiceDesc.Streams[0], Brood_Subscribe_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &broodSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Brood_SubscribeClient interface {
	Recv() (*Spore, error)
	grpc.ClientStream
}

type broodSubscribeClient struct {
	grpc.ClientStream
}

func (x *broodSubscribeClient) Recv() (*Spore, error) {
	m := new(Spore)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BroodServer is the server API for Brood service.
// All implementations must embed UnimplementedBroodServer
// for forward compatibility
type BroodServer interface {
	Publish(context.Context, *Spore) (*Acknowledgement, error)
	Subscribe(*Hive, Brood_SubscribeServer) error
	mustEmbedUnimplementedBroodServer()
}

// UnimplementedBroodServer must be embedded to have forward compatible implementations.
type UnimplementedBroodServer struct {
}

func (UnimplementedBroodServer) Publish(context.Context, *Spore) (*Acknowledgement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedBroodServer) Subscribe(*Hive, Brood_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedBroodServer) mustEmbedUnimplementedBroodServer() {}

// UnsafeBroodServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BroodServer will
// result in compilation errors.
type UnsafeBroodServer interface {
	mustEmbedUnimplementedBroodServer()
}

func RegisterBroodServer(s grpc.ServiceRegistrar, srv BroodServer) {
	s.RegisterService(&Brood_ServiceDesc, srv)
}

func _Brood_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Spore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroodServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Brood_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroodServer).Publish(ctx, req.(*Spore))
	}
	return interceptor(ctx, in, info, handler)
}

func _Brood_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Hive)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BroodServer).Subscribe(m, &broodSubscribeServer{stream})
}

type Brood_SubscribeServer interface {
	Send(*Spore) error
	grpc.ServerStream
}

type broodSubscribeServer struct {
	grpc.ServerStream
}

func (x *broodSubscribeServer) Send(m *Spore) error {
	return x.ServerStream.SendMsg(m)
}

// Brood_ServiceDesc is the grpc.ServiceDesc for Brood service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Brood_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.brood",
	HandlerType: (*BroodServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Brood_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Brood_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "brood.proto",
}
