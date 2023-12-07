package main

import (
	pb "code_god/1_example/proto"
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

/**
  @author: CodeWater
  @since: 2023/11/30
  @desc: grpc client
	双向认证
**/

func main() {
	// 证书认证-双向认证
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, _ := tls.LoadX509KeyPair("F:/Code/GoCode/grpc/code_god/2_tls/key/client.pem", "F:/Code/GoCode/grpc/code_god/2_tls/key/client.key")
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("F:/Code/GoCode/grpc/code_god/2_tls/key/ca.crt")
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ServerName: "*.codewater.com",
		RootCAs:    certPool,
	})

	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal("failed to dial server:", err)
	}
	defer conn.Close()

	prodClient := pb.NewProductServiceClient(conn)
	request := &pb.ProductRequest{
		ProdId:   2,
		ProdName: "golang",
	}
	stockResponse, err := prodClient.GetProductStock(context.Background(), request)
	if err != nil {
		log.Fatal("failed to get stock:", err)
	}

	log.Printf("stock:%#+v\n", stockResponse)
}
