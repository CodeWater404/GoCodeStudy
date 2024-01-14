package main

import "fmt"

/**
  @author: CodeWater
  @since: 2024/1/15
  @desc: 内存逃逸场景
	4. 动态类型变量
**/

func escape4() {
	//fmt.Println(a ...any)的函数参数a是一个动态类型变量，编译器无法确定a的类型，所以a会逃逸到堆上
	fmt.Println(111)
}

func main() {
	escape4()
	// go build -gcflags=-m 4_dynamic_type.go
}
