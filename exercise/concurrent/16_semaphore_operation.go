package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"runtime"
	"time"
)

/**
  @author: CodeWater
  @since: 2024/1/17
  @desc: semaphore信号量操作
	Semaphore带权重的信号量，控制多个goroutine同时访问资源
	使用场景：控制goroutine的阻塞和唤醒
**/

var (
	maxWorkers = runtime.GOMAXPROCS(0)
	sema       = semaphore.NewWeighted(int64(maxWorkers)) //信号量
	task       = make([]int, maxWorkers*4)
)

func main() {
	ctx := context.Background()
	for i := range task {
		//如果没有worker可用，会阻塞在这里，直到有worker可用
		if err := sema.Acquire(ctx, 1); err != nil {
			break
		}
		//启动goroutine处理任务
		go func(i int) {
			defer sema.Release(1)
			time.Sleep(100 * time.Millisecond) //模拟耗时操作
			task[i] = i + 1
		}(i)
	}
	//请求所有的worker，这样可以确保前面的goroutine都执行完
	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("获取所有的worker失败: %v\n", err)
	}
	fmt.Println("maxWorkers: ", maxWorkers, ", task: ", task)
}
