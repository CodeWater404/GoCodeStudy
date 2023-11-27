package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/11/27
  @desc: channel的基本使用
**/

func main() {
	var ch = make(chan int, 10)
	for i := 0; i < 10; i++ {
		//select语句会随机选择一个可执行的case进行操作，如果有多个case同时满足条件，则会随机选择一个。
		//如果所有通道都没有准备好，那么执行 default 块中的代码。
		select {
		case ch <- i:
		case x := <-ch:
			fmt.Println("receive from ch:", x)
		}
	}
}
