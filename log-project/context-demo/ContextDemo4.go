package main

import (
	"context"
	"fmt"
)

/**
  @author: CodeWater
  @since: 2023/8/6
  @desc: withcancel的使用
	context控制关闭goroutine的另外一个案例
**/

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		select {
		case <-ctx.Done():
			return
		case dst <- n:
			n++
		}
	}()
	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		//当n为5的时候，break退出for，main也就结束，执行cancel函数，这个时候子goroutine就会收到结束的信号，响应的dst中的
		//子goroutine也会退出
		if n == 5 {
			break
		}
	}
}
