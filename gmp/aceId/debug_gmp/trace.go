package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

/**
  @author: CodeWater
  @since: 2023/12/18
  @desc: trace的编程过程
	1. 创建文件
	2. 启动
	3. 停止
	go run trace.go 			===> trace.out
	go tool trace trace.out		===> 2023/12/18 22:05:25 Opening browser. Trace viewer is listening on http://127.0.0.1:63031
**/

func main() {
	// 1. 创建一个trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 2. 启动trace
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	// 正常要调式的业务代码
	fmt.Println("Hello gmp")

	// 3. 停止trace
	trace.Stop()
}
