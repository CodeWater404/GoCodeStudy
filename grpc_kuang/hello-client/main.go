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

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "code",
		"appkey": "1231",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return false
}

func main() {
	//1.tls认证
	//creds , _ := credentials.NewClientTLSFromFile("" , "*.codewater.com")
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds))

	//2.token认证
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))
	conn, err := grpc.Dial("127.0.0.1:9090", opts...)

	//3.连接到远程，这里没有加密(没有认证的)
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("client did not connect , err: %v\n", err)
		return
	}
	defer conn.Close()

	//建立连接
	client := pb.NewSayHelloClient(conn)
	//执行rpc调用，方法在服务端实现并返回结果
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "codewater"})
	if err != nil {
		log.Fatalf("client request failed :%v\n", err)
		return
	}
	log.Printf("client received msg: %v\n", resp.GetResponseMsg())

}
