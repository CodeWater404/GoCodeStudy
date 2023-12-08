package main

import (
	pb "code_god/3_token/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

/**
  @author: CodeWater
  @since: 2023/11/30
  @desc: grpc server
	token 认证
**/

type ProductService struct {
	pb.UnimplementedProductServiceServer
}

// GetProductStock 实现服务端接口
func (p *ProductService) GetProductStock(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	log.Printf("server received prodId:%v , prodName:%v\n", req.GetProdId(), req.GetProdName())
	return &pb.ProductResponse{ProdStack: 1, ProdPrice: 100}, nil
}

func Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}
	var (
		user     string
		password string
	)
	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}
	if user != "code" || password != "123" {
		return status.Errorf(codes.Unauthenticated, "token invalid")
	}

	return nil
}

func main() {
	var authInterceptor grpc.UnaryServerInterceptor
	authInterceptor = func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		//拦截普通方法请求，验证 Token
		err = Auth(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}

	server := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	pb.RegisterProductServiceServer(server, &ProductService{})

	lister, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("failed to listen:%v\n", err)
	}

	log.Println("server start...")

	err = server.Serve(lister)
	if err != nil {
		log.Fatalf("failed to serve:%v\n", err)
	}
}
