package main

import (
	pb "code_god/4_stream/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/30
  @desc: grpc client
**/

var (
	conn   *grpc.ClientConn
	err    error
	client pb.ProductServiceClient
	ctx    = context.Background()
)

// func new() {
func init() {
	// 连接server，无加密
	conn, err = grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("client did not connect , err: %v\n", err)
		return
	}
	// 建立连接
	client = pb.NewProductServiceClient(conn)
	log.Printf("client init success!\n")
}

// NormalRpcCall 普通rpc调用
func NormalRpcCall() {
	// 执行rpc调用，方法在服务端实现并返回结果
	resp, err := client.GetProductStock(context.Background(), &pb.ProductRequest{ProdId: 1, ProdName: "codeIpad"})
	if err != nil {
		log.Fatalf("client request failed :%v\n", err)
		return
	}
	fmt.Printf("client received prodStack:%v , prodPrice:%v\n", resp.GetProdStack(), resp.ProdPrice)
}

// ClientStreamRpcCall	客户端流式rpc调用,服务端实现调用的方法
func ClientStreamRpcCall() {
	stream, err := client.UpdateProductClientStream(ctx)
	if err != nil {
		log.Fatalf("client stream get failed:%v\n", err)
	}
	resp := make(chan struct{}, 1)
	go prodRequest(stream, resp)
	select {
	case <-resp:
		recv, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("client stream close failed:%v\n", err)
		}
		log.Printf("===> client stream get resp:%#+v\n", recv)
	}
	//defer conn.Close()
}

// prodRequest 客户端不断发送请求
func prodRequest(stream pb.ProductService_UpdateProductClientStreamClient, resp chan struct{}) {
	count := 0
	for {
		request := &pb.ProductRequest{
			ProdId:   123,
			ProdName: "code_stream",
		}
		err = stream.Send(request)
		if err != nil {
			log.Fatalf("client stream prodId:%v , send failed:%v\n", request.ProdId, err)
		}
		time.Sleep(time.Second)
		count++
		if count > 10 {
			resp <- struct{}{}
			break
		}
	}
}
func main() {
	log.Println("client start...")
	//在Go语言中，“defer”语句会在函数完成执行后安排函数执行，但延迟函数的参数会立即计算，所以不能在这里close
	//new() // 这里可以改为init，不用显示调用
	defer conn.Close()
	//NormalRpcCall()
	ClientStreamRpcCall()
}
