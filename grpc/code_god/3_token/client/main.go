package main

import (
	pb "code_god/3_token/proto"
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
	token认证
**/

// Authentication 客户端需要实现 `PerRPCCredentials` 接口:GetRequestMetadata、RequireTransportSecurity
type Authentication struct {
	User     string
	Password string
}

// GetRequestMetadata 返回认证需要的必要信息
func (a *Authentication) GetRequestMetadata(ctx context.Context, str ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}

// RequireTransportSecurity 方法表示是否启用安全链接，在生产环境中，一般都是启用的，但为了测试方便，暂时这里不启用了。
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

func main() {
	log.Println("client start...")

	user := &Authentication{
		User:     "code",
		Password: "123",
	}

	// 连接server，无加密
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithPerRPCCredentials(user))
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
