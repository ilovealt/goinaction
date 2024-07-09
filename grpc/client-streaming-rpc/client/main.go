package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ilovealt/goinaction/grpc/client-streaming-rpc/ecommerce"
)

func main() {
	// 测试环境，取消安全凭证
	conn, err := grpc.NewClient("127.0.0.1:8009", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementClient(conn)

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	// 业务逻辑
	stream, err := client.UpdateOrders(ctx)

	if stream.Send(&pb.Order{
		Id:          "00",
		Items:       []string{"A", "B"},
		Description: "A with B",
		Price:       0.11,
		Destination: "ABC",
	}); err != nil {
		panic(err)
	}

	if err := stream.Send(&pb.Order{
		Id:          "01",
		Items:       []string{"C", "D"},
		Description: "C with D",
		Price:       1.11,
		Destination: "ABCDEFG",
	}); err != nil {
		panic(err)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}

	log.Printf("Update Orders Res : %s", res)

}
