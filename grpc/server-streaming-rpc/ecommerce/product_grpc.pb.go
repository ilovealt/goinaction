// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.6.1
// source: product.proto

package ecommerce

import (
	context "context"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	OrderManagement_SearchOrders_FullMethodName = "/ecommerce.OrderManagement/searchOrders"
)

// OrderManagementClient is the client API for OrderManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderManagementClient interface {
	SearchOrders(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (OrderManagement_SearchOrdersClient, error)
}

type orderManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderManagementClient(cc grpc.ClientConnInterface) OrderManagementClient {
	return &orderManagementClient{cc}
}

func (c *orderManagementClient) SearchOrders(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (OrderManagement_SearchOrdersClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OrderManagement_ServiceDesc.Streams[0], OrderManagement_SearchOrders_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &orderManagementSearchOrdersClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderManagement_SearchOrdersClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderManagementSearchOrdersClient struct {
	grpc.ClientStream
}

func (x *orderManagementSearchOrdersClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderManagementServer is the server API for OrderManagement service.
// All implementations must embed UnimplementedOrderManagementServer
// for forward compatibility
type OrderManagementServer interface {
	SearchOrders(*wrappers.StringValue, OrderManagement_SearchOrdersServer) error
	mustEmbedUnimplementedOrderManagementServer()
}

// UnimplementedOrderManagementServer must be embedded to have forward compatible implementations.
type UnimplementedOrderManagementServer struct {
}

func (UnimplementedOrderManagementServer) SearchOrders(*wrappers.StringValue, OrderManagement_SearchOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchOrders not implemented")
}
func (UnimplementedOrderManagementServer) mustEmbedUnimplementedOrderManagementServer() {}

// UnsafeOrderManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderManagementServer will
// result in compilation errors.
type UnsafeOrderManagementServer interface {
	mustEmbedUnimplementedOrderManagementServer()
}

func RegisterOrderManagementServer(s grpc.ServiceRegistrar, srv OrderManagementServer) {
	s.RegisterService(&OrderManagement_ServiceDesc, srv)
}

func _OrderManagement_SearchOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrappers.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderManagementServer).SearchOrders(m, &orderManagementSearchOrdersServer{ServerStream: stream})
}

type OrderManagement_SearchOrdersServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type orderManagementSearchOrdersServer struct {
	grpc.ServerStream
}

func (x *orderManagementSearchOrdersServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

// OrderManagement_ServiceDesc is the grpc.ServiceDesc for OrderManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.OrderManagement",
	HandlerType: (*OrderManagementServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "searchOrders",
			Handler:       _OrderManagement_SearchOrders_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "product.proto",
}
