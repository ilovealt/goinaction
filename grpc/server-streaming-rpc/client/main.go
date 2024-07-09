package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/ilovealt/goinaction/grpc/server-streaming-rpc/ecommerce"
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

	// Get Order
	stream, err := client.SearchOrders(ctx, &wrapperspb.StringValue{Value: "Google"})
	if err != nil {
		// 如果对服务端不存在这个数据，则执行次操作
		log.Println("说明gRPC已经通畅，只是数据错误！")
		panic(err)
	}

	for {
		order, err := stream.Recv()
		if err == io.EOF {
			break
		}

		log.Println("Search Result: ", order)
	}

}
