package main

import "fmt"

/**
  @author: CodeWater
  @since: 2024/1/14
  @desc: 内存逃逸场景
	2. 栈空间不足逃逸
**/

// escape2 当栈空间足够时，不会发生逃逸，但是当变量过大时，已经完全超过栈空间的大小时，将
// 会发生逃逸到堆上分配内存。局部变量s占用内存过大，编译器会将其分配到堆上
func escape2() {
	s := make([]int, 0, 10000)
	for index, _ := range s {
		fmt.Printf("%d\n", index)
		s[index] = index

	}
}

func main() {
	escape2()
	// go build -gcflags=-m 2_out_of_stack_space.go
}
