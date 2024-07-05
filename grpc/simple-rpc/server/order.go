package main

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/ilovealt/goinaction/grpc/simple-rpc/ecommerce"
)

var _ pb.OrderManagementServer = &OrderManagementImpl{}

var orders = make(map[string]pb.Order)

type OrderManagementImpl struct {
	pb.UnimplementedOrderManagementServer
}

// Simple RPC
func (s *OrderManagementImpl) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
	//order := pb.Order{Id: "101", Items: []string{"iPhone XS", "Mac Book Pro"}, Destination: "San Jose, CA", Price: 2300.00}
	//orders["101"] = order
	ord, exists := orders[orderId.Value]
	if exists {
		return &ord, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Order does not exist. : ", orderId)
}
