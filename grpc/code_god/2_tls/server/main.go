package main

import (
	pb "code_god/1_example/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

/**
  @author: CodeWater
  @since: 2023/11/30
  @desc: grpc server
	单向认证:有中间人攻击
**/

type ProductService struct {
	pb.UnimplementedProductServiceServer
}

// GetProductStock 实现服务端接口
func (p *ProductService) GetProductStock(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	log.Printf("server received prodId:%v , prodName:%v\n", req.GetProdId(), req.GetProdName())
	return &pb.ProductResponse{ProdStack: 1, ProdPrice: 100}, nil
}

func main() {
	// 添加证书
	file, err := credentials.NewServerTLSFromFile("F:/Code/GoCode/grpc/code_god/2_tls/key/server.pem", "F:/Code/GoCode/grpc/code_god/2_tls/key/server.key")
	if err != nil {
		log.Fatal("加载证书失败:", err)
		return
	}

	server := grpc.NewServer(grpc.Creds(file)) // 相比没有加密的，这里只是多了一个Creds
	pb.RegisterProductServiceServer(server, &ProductService{})

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("server start success...")

	// 阻塞监听
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
