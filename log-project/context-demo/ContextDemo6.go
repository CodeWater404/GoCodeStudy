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
  @desc: context.withvalue的使用：
	记录一次请求的上下文(日志系统的场景)
**/

type TraceCode string

var wg6 sync.WaitGroup

func worker61(ctx context.Context) {
	key := TraceCode("TRACE_CODE")           //把普通字符串类型转为treenode类型
	traceCode, ok := ctx.Value(key).(string) //在子goroutine中取trace code
	if !ok {
		fmt.Println("invalid trace code")
		return
	}
LOOP:
	for {
		fmt.Printf("worker61 , trace code: %v\n", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	fmt.Println("worker61 done!!!!")
	wg6.Done()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	//在系统的入口中设置trace code传递给后续启动的goroutine:实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "111111111")
	wg6.Add(1)
	go worker61(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg6.Wait()
	fmt.Println("main61 over....")

}
