package main

import (
	"context"
	"fmt"
	"time"
)

/**
  @author: CodeWater
  @since: 2024/1/17
  @desc: context.Context操作
	sync.Context可以进行上下文信息传递、提供超时和取消机制、控制子goroutine的执行
	使用场景：取消一个goroutine的执行
**/

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer func() {
			fmt.Println("goroutine exit")
		}()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("receive cancel signal")
				return
			default:
				fmt.Println("running default...")
				time.Sleep(time.Second * 1)
			}
		}
	}()

	time.Sleep(time.Second * 1)
	cancel()
	//等待goroutine的退出
	time.Sleep(time.Second * 2)
}
