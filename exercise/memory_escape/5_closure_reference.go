package main

/**
  @author: CodeWater
  @since: 2024/1/15
  @desc: 内存逃逸场景
	5. 闭包引用对象
**/

func escape5() func() int {
	var i int = 1
	return func() int {
		// i是一个局部变量，但是返回的函数引用了i，所以i会逃逸到堆上
		i++
		return i
	}
}

func main() {
	escape5()
	// go build -gcflags=-m 5_closure_reference.go
}
