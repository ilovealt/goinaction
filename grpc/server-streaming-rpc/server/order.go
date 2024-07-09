package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/ilovealt/goinaction/grpc/server-streaming-rpc/ecommerce"
)

var _ pb.OrderManagementServer = &OrderManagementImpl{}

var orders = map[string]pb.Order{
	"101": {
		Id: "101",
		Items: []string{
			"Google",
			"Baidu",
		},
		Description: "example",
		Price:       0,
		Destination: "example",
	},
}

type OrderManagementImpl struct {
	pb.UnimplementedOrderManagementServer
}

// Simple RPC
func (s *OrderManagementImpl) SearchOrders(query *wrapperspb.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for _, order := range orders {
		for _, str := range order.Items {
			if strings.Contains(str, query.Value) {
				err := stream.Send(&order)
				if err != nil {
					return fmt.Errorf("## error send: %v", err)
				}
			}
		}
	}

	return nil
}
