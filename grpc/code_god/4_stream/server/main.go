package main

import (
	pb "code_god/4_stream/proto"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
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

func (p *ProductService) UpdateProductClientStream(stream pb.ProductService_UpdateProductClientStreamServer) error {
	count := 0
	for {
		// 源源不断的去接收客户端发来的消息
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		log.Printf("===>server received:%#+v , count:%v \n", recv, count)
		count++
		if count > 10 {
			resp := &pb.ProductResponse{
				ProdStack: 100,
				ProdPrice: 111,
			}
			err := stream.SendAndClose(resp)
			if err != nil {
				return err
			}
			return nil
		}
	}
}

// GetProductStockServerStream 服务端流,不断向客户端发送消息
func (p *ProductService) GetProductStockServerStream(req *pb.ProductRequest, stream pb.ProductService_GetProductStockServerStreamServer) error {
	count := 0
	for {
		resp := &pb.ProductResponse{
			ProdStack: 200,
			ProdPrice: 222,
		}
		err := stream.Send(resp)
		if err != nil {
			log.Fatalf("server send failed:%v , count:%v\n", err, count)
		}
		time.Sleep(time.Second)
		count++
		if count > 10 {
			return nil
		}
	}
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
