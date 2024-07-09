package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/ilovealt/goinaction/grpc/server-streaming-rpc/ecommerce"
)

func main() {
	// 测试环境，取消安全凭证
	s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	pb.RegisterOrderManagementServer(s, &OrderManagementImpl{})

	lis, err := net.Listen("tcp", ":8009")
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
