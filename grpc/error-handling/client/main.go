package main

import (
	"context"
	"log"
	"time"

	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "github.com/ilovealt/goinaction/grpc/error-handling/ecommerce"
)

func main() {
	// 测试环境，取消安全凭证
	conn, err := grpc.NewClient("127.0.0.1:8009", grpc.WithTransportCredentials(insecure.NewCredentials()))

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

		st, ok := status.FromError(err)
		/*
			if ok && st.Code() == codes.NotFound {
				log.Println(st.Code(), st.Message())
			} else {
				log.Println(err)
			}
		*/
		if !ok {
			log.Println(err)
			return
		}

		switch st.Code() {
		case codes.NotFound:
			for _, d := range st.Details() {
				switch info := d.(type) {
				case *epb.BadRequest_FieldViolation:
					log.Printf("Request Field Invalid: %s", info)
				default:
					log.Printf("Unexpected error type: %s", info)
				}
			}
		default:
			log.Printf("Unhandled error : %s ", st.String())
		}

		panic(err)
	}

	log.Print("GetOrder Response -> : ", retrievedOrder)
}
