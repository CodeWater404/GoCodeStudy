package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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
	//token认证的
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("未传输token")
	}
	var appid string
	var appkey string
	if v, ok := md["appid"]; ok {
		appid = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appkey = v[0]
	}

	//一些校验逻辑
	if appid != "code" || appkey != "123" {
		return nil, errors.New("token error")
	}

	log.Printf("server received: %v\n", req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello " + req.RequestName}, nil
}

func main() {
	//todo：生成文件失败
	//tls认证
	//creds, _ := credentials.NewServerTLSFromFile("F:\\Code\\GoCode\\grpc_kuang\\key\\test.pem" , "F:\\Code\\GoCode\\grpc_kuang\\key\\test.key")

	//开启端口
	listen, _ := net.Listen("tcp", "127.0.0.1:9090")

	//1.tls
	//grpcServer := grpc.NewServer(grpc.Creds(creds))

	//2. token
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	//0.创建grpc服务(没有认证的)
	//grpcServer := grpc.NewServer()
	//注册服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("grpc server run failed , err: %v\n", err)
		return
	}
}
