package main

import (
	pb "code_god/1_example/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

/**
  @author: CodeWater
  @since: 2023/11/30
  @desc: grpc client
**/

func main() {
	log.Println("client start...")

	// 连接server，无加密
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("client did not connect , err: %v\n", err)
		return
	}
	defer conn.Close()

	// 建立连接
	client := pb.NewProductServiceClient(conn)
	// 执行rpc调用，方法在服务端实现并返回结果
	resp, err := client.GetProductStock(context.Background(), &pb.ProductRequest{ProdId: 1, ProdName: "codeIpad"})
	if err != nil {
		log.Fatalf("client request failed :%v\n", err)
		return
	}
	fmt.Printf("client received prodStack:%v , prodPrice:%v\n", resp.GetProdStack(), resp.ProdPrice)
}
