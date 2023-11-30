package main

import (
	pb "code_god/1_example/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

/**
  @author: CodeWater
  @since: 2023/11/30
  @desc: grpc server
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
	log.Println("server start...")
	server := grpc.NewServer()
	pb.RegisterProductServiceServer(server, &ProductService{})

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	_ = server.Serve(listener)
}
