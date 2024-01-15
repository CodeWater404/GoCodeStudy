package main

import (
	"fmt"
	"sync"
)

/**
  @author: CodeWater
  @since: 2024/1/15
  @desc: once一次性操作
	synC.Once可以保证在G0程序运行期间的某段代码只会执行一次
	使用场景：常常用于单例对象的初始化场景
**/

func main() {
	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println("only once")
		})
	}
}
