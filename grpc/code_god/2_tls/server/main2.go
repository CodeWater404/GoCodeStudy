package main

import (
	pb "code_god/2_tls/proto"
	service "code_god/2_tls/proto"
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
)

/**
  @author: CodeWater
  @since: 2023/11/30
  @desc: grpc server
	双向认证
**/

type ProductService2 struct {
	pb.UnimplementedProductServiceServer
}

// GetProductStock 实现服务端接口
func (p *ProductService2) GetProductStock(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	log.Printf("server received prodId:%v , prodName:%v\n", req.GetProdId(), req.GetProdName())
	return &pb.ProductResponse{ProdStack: 1, ProdPrice: 100}, nil
}

func main() {
	// 双向认证，添加证书
	cert, err := tls.LoadX509KeyPair("F:/Code/GoCode/grpc/code_god/2_tls/key/server.pem", "F:/Code/GoCode/grpc/code_god/2_tls/key/server.key")
	if err != nil {
		log.Fatal("加载证书失败:", err)
		return
	}

	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("F:/Code/GoCode/grpc/code_god/2_tls/key/ca.crt")
	if err != nil {
		log.Fatal("failed to read ca.crt:", err)
		return
	}
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ClientAuth: tls.RequireAndVerifyClientCert,
		// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})

	rpcServer := grpc.NewServer(grpc.Creds(creds))

	// 我们传递一个指针（&ProductService2{}）而不是一个值（Product service 2{}）的原因是在ProductService2上定义的方法可能有指针接收器。在Go中，具有指针接收器的方法只能在指针上调用。通过传递指针，我们确保ProductService2的方法可以在必要时修改结构的字段。
	service.RegisterProductServiceServer(rpcServer, &ProductService2{})

	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("failed to listen:%v\n", err)
	}
	log.Println("server start success...")
	err = rpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve:%v\n", err)
	}
	log.Println("server start success...")

}
