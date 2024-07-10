package main

import (
	"io"
	"log"

	pb "github.com/ilovealt/goinaction/grpc/bidirectional-streaming-rpc/ecommerce"
)

const (
	orderBatchSize = 3
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

func (s *OrderManagementImpl) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {

	batchMarker := 1
	var combinedShipmentMap = make(map[string]pb.CombinedShipment)
	for {
		orderId, err := stream.Recv()
		// 客户端停止发送数据时，把所有服务器端结果返回
		if err == io.EOF {
			for _, shipment := range combinedShipmentMap {
				if err := stream.Send(&shipment); err != nil {
					return err
				}
			}
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}

		// 这个位置很明显的数据可能为空
		destination := orders[orderId.GetValue()].Destination
		log.Println("集合orders： 如果不存在order，直接获取Destination会发生什么？", destination)
		shipment, found := combinedShipmentMap[destination]
		log.Println("集合combinedShipmentMap： 如果不存在destination，直接获取会发生什么？", shipment)

		if found {
			ord := orders[orderId.GetValue()]
			shipment.OrderList = append(shipment.OrderList, &ord)
			combinedShipmentMap[destination] = shipment
		} else {
			comShip := pb.CombinedShipment{Id: "cmb - " + (orders[orderId.GetValue()].Destination), Status: "Processed!"}
			ord := orders[orderId.GetValue()]
			comShip.OrderList = append(shipment.OrderList, &ord)
			combinedShipmentMap[destination] = comShip
			log.Print("new add:", len(comShip.OrderList), comShip.GetId())
		}

		// 信息达到限定值
		if batchMarker == orderBatchSize {
			for _, comb := range combinedShipmentMap {
				log.Printf("Shipping : %v -> %v", comb.Id, len(comb.OrderList))
				if err := stream.Send(&comb); err != nil {
					return err
				}
			}
			batchMarker = 0
			combinedShipmentMap = make(map[string]pb.CombinedShipment)
		} else {
			batchMarker++
		}

	}

}
