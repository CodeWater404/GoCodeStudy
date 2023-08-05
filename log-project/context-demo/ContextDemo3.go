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
  @desc:  让goroutine停止的几种方法
	1. 使用全局变量（本例）
		缺点：单个文件的时候还行；但是一旦变成工程化的时候就很容易冲突
	2. 使用channel
		缺点：值传来传去，不确定。
	3. 使用context，用法本质类似channel里面传一个空结构体(本例)
**/

var wg3 sync.WaitGroup

func work31(ctx context.Context) {
	defer wg3.Done()
	//使用context可以在嵌套使用goroutine
	go work32(ctx)
LABEL:
	for {
		fmt.Println("worker32...")
		time.Sleep(time.Second)
		select {
		//查看文档，可以发现其实返回的就是一个空的结构体
		case <-ctx.Done():
			break LABEL
		default:

		}
	}
}

func runWorker31() {
	ctx, cancel := context.WithCancel(context.Background())
	wg3.Add(1)
	go work31(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg3.Wait()
	fmt.Println("runWorker31 over.....")

}

func work32(ctx context.Context) {
	defer wg3.Done()
LABEL:
	for {
		fmt.Println("worker31...")
		time.Sleep(time.Second)
		select {
		//查看文档，可以发现其实返回的就是一个空的结构体
		case <-ctx.Done():
			break LABEL
		default:

		}
	}
}

func main() {
	runWorker31()
}
