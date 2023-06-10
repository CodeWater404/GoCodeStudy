package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc_kuang/hello-server/proto"
	"log"
)

/**
  @author: CodeWater
  @since: 2023/6/8
  @desc: grpc客户端代码编写
**/

func main() {
	//tls
	//creds , _ := credentials.NewClientTLSFromFile("" , "*.codewater.com")

	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds))

	//连接到远程，这里没有加密
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("client did not connect , err: %v\n", err)
		return
	}
	defer conn.Close()

	//建立连接
	client := pb.NewSayHelloClient(conn)
	//执行rpc调用，方法在服务端实现并返回结果
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "codewater"})
	log.Printf("client received msg: %v\n", resp.GetResponseMsg())

}
