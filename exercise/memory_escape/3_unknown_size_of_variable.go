package main

/**
  @author: CodeWater
  @since: 2024/1/15
  @desc: 内存逃逸场景
	3. 变量大小不确定
**/

func escape3() {
	number := 10
	s := make([]int, number) // 编译期间，编译器无法确定s的大小，所以s会逃逸到堆上
	// s := make([]int , 10) // 这样的， 编译期间，编译器可以确定s的大小，所以s不会逃逸到堆上
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
}

func main() {
	escape3()
	// go build -gcflags=-m 3_unknown_size_of_variable.go
}
