package main

import (
	"net"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/23
  @desc: 通过mutex来初始化单例对象
**/

// 使用互斥锁保证线程安全
var connMu sync.Mutex
var conn net.Conn

func getConn() net.Conn {
	connMu.Lock()
	defer connMu.Unlock()

	if conn != nil {
		return conn
	}
	conn, _ = net.DialTimeout("tcp", "baidu.com:80", 10*time.Second)
	return conn
}

func main() {
	conn := getConn()
	if conn == nil {
		panic("conn is nil")
	}
}
