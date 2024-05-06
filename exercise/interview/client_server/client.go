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

func main() {
	// 连接服务器
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("连接服务器失败:", err)
		return
	}
	defer conn.Close()

	// 发送消息给服务器
	msg := "Hello"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("发送消息失败:", err)
		return
	}
	fmt.Println("已向服务器发送消息:", msg)

	// 接收服务器回复的消息
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("接收服务器回复消息失败:", err)
		return
	}
	serverMsg := string(buffer[:n])
	fmt.Println("服务器回复消息:", serverMsg)
}
