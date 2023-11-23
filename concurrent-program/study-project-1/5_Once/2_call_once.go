package main

import (
	"fmt"
	"sync"
)

/**
  @author: CodeWater
  @since: 2023/11/23
  @desc: once的使用
**/

func main() {
	var once sync.Once

	//可以多次调用 Do 方法，但是只有第一次调用 `Do` 方法时 f 参数才会执行,
	//即使第二次、第三次、第 n 次调用时 f 参数的值不一样，也不会被执行
	f1 := func() {
		fmt.Println("f1 called")
	}
	once.Do(f1)

	f2 := func() {
		fmt.Println("f2 called")
	}
	once.Do(f2)
}
