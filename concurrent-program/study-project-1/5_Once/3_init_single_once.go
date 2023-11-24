package main

import (
	"net"
	"sync"
)

/**
  @author: CodeWater
  @since: 2023/11/23
  @desc: $
**/

func main() {
	var addr = "baidu.com"

	var conn net.Conn
	var err error
	var once sync.Once

	once.Do(func() {
		conn, err = net.Dial("tcp", addr)
	})

	if err != nil {
		panic(err)
	}
	if conn == nil {
		panic("conn is nil")
	}
}
