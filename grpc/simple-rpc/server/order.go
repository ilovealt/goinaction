package main

import (
	"context"

	pb "github.com/ilovealt/goinaction/grpc/simple-rpc/ecommerce"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var _ pb.OrderManagementServer = &OrderManagementImpl{}

var orders = make(map[string]pb.Order)

type OrderManagementImpl struct {
	pb.UnimplementedOrderManagementServer
}

// 实现接口中承诺的接口方法
func (s *OrderManagementImpl) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
	ord, exists := orders[orderId.Value]
	if exists {
		return &ord, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Order does not exist. : ", orderId)
}
