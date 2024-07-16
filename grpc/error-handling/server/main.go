package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ilovealt/goinaction/grpc/error-handling/ecommerce"
)

// 创建一个拦截其
// 函数名称自定义，参数和返回值需要按规定写
func orderUnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	s := time.Now()
	// Invoking the handler to complete the normal execution of a unary RPC.
	m, err := handler(ctx, req)
	// Post processing logic
	log.Printf("Method: %s, req: %s, resp: %s, latency: %s\n", info.FullMethod, req, m, time.Now().Sub(s))
	return m, err
}

func main() {
	// 测试环境，取消安全凭证
	s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()), grpc.UnaryInterceptor(orderUnaryServerInterceptor))

	pb.RegisterOrderManagementServer(s, &OrderManagementImpl{})

	lis, err := net.Listen("tcp", ":8009")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
