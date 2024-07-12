package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/ilovealt/goinaction/grpc/simple-rpc/ecommerce"
)

func orderUnaryClientInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	s := time.Now()
	// 方法调用
	err := invoker(ctx, method, req, reply, cc, opts...)

	log.Printf("Clinet method: %s, req: %s, resp: %s, latency: %s\n", method, req, reply, time.Since(s))

	return err
}

func main() {
	// 测试环境，取消安全凭证
	conn, err := grpc.NewClient("127.0.0.1:8009", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(orderUnaryClientInterceptor))

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Get Order
	retrievedOrder, err := client.GetOrder(ctx, &wrapperspb.StringValue{Value: "101"})
	if err != nil {
		// 如果对服务端不存在这个数据，则执行次操作
		log.Println("说明gRPC已经通畅，只是数据错误！")
		panic(err)
	}

	log.Print("GetOrder Response -> : ", retrievedOrder)
}
