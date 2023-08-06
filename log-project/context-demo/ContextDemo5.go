package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/8/6
  @desc: context.withtimeout的使用：
	让子goroutine退出
**/

var wg sync.WaitGroup

func worker51(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting....")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	fmt.Println("worker51 done!")
	wg.Done()
}

func main() {
	//50毫秒之后，ctx就会结束，就会像子goroutine发出退出信号，然后子goroutine就会退出
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker51(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("main51 over...")
}
