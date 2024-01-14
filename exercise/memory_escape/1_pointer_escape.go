package main

/**
  @author: CodeWater
  @since: 2024/1/14
  @desc:
	编译器会根据变量是否被外部引用来决定是否逃逸：
	1、如果函数外部没有引用，则优先放到栈中：
	2、如果函数外部存在引用，则必定放到堆中：
	3、如果栈上放不下，则必定放到堆上
	内存逃逸场景：
		1. 指针逃逸
**/

func escape1() *int {
	var a int = 1
	return &a

}

func main() {
	escape1()
	// go build -gcflags=-m 1_pointer_escape.go
}
