package main

import (
	"fmt"
)

/**
  @author: CodeWater
  @since: 2023/4/17
  @desc: error  vs  panic:
		1.意料之中的：使用error。如：文件打不开
		2.意料之外的：使用panic。如：数组越界
**/

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred: ", err)
		} else {
			panic(err)
		}
	}() //这里”（）“是表示这个匿名函数的调用

	//1.error
	//panic(errors.New("this is an error"))

	//2.error
	b := 0
	a := 5 / b
	fmt.Println("a:", a)

	//3.not a error, enter else
	//panic(123)
}

func main() {
	tryRecover()
}
