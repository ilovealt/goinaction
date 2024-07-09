package main

import (
	"io"
	"log"

	pb "github.com/ilovealt/goinaction/grpc/client-streaming-rpc/ecommerce"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

func (s *OrderManagementImpl) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	ordersStr := "Updated Order IDs : "
	for {
		order, err := stream.Recv()
		if err == io.EOF {
			// Finished reading the order stream.
			return stream.SendAndClose(&wrapperspb.StringValue{Value: "Orders processed " + ordersStr})
		}
		// Update order
		orders[order.Id] = *order

		log.Println("Order ID ", order.Id, ": Updated")
		ordersStr += order.Id + ", "
	}
}
