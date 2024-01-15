package main

import (
	"fmt"
	"time"
)

/**
  @author: CodeWater
  @since: 2024/1/15
  @desc: 通道操作
	channel管道，高级同步原语，goroutine之间通信的桥梁
	使用场景：消息队列、数据传递、信号通知、任务编排、锁
**/

func main() {
	c := make(chan struct{}, 1)
	for i := 0; i < 10; i++ {
		go func() {
			c <- struct{}{}
			time.Sleep(time.Second * 1)
			fmt.Println("visit critical zone by ch")
			<-c
		}()
	}

	for {
	}
}
