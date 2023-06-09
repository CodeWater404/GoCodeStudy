package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc_kuang/hello-server/proto"
	"log"
	"net"
)

/**
  @author: CodeWater
  @since: 2023/6/8
  @desc: 服务端代码
**/

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("server received: %v\n", req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	//开启端口
	listen, _ := net.Listen("tcp", "127.0.0.1:9090")
	//创建grpc服务
	grpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("grpc server run failed , err: %v\n", err)
		return
	}
}
