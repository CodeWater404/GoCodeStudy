package main

import (
	"fmt"
	"net"
)

/**
  @author: CodeWater
  @since: 2024/3/7
  @desc: $
**/

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 接收客户端发送的消息
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("接收消息失败:", err)
		return
	}
	clientMsg := string(buffer[:n])
	fmt.Println("客户端消息:", clientMsg)

	// 回复客户端消息
	reply := "你好"
	_, err = conn.Write([]byte(reply))
	if err != nil {
		fmt.Println("回复消息失败:", err)
		return
	}
	fmt.Println("已向客户端回复:", reply)
}

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("监听端口失败:", err)
		return
	}
	defer listener.Close()
	fmt.Println("服务器已启动，等待客户端连接...")

	// 接受客户端连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受客户端连接失败:", err)
			continue
		}
		fmt.Println("客户端连接成功:", conn.RemoteAddr())
		go handleConnection(conn)
	}
}
