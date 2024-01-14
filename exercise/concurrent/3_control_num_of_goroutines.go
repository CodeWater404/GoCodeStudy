package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
  @author: CodeWater
  @since: 2024/1/14
  @desc: 控制并发数量1
	采用channel
**/

var wg = sync.WaitGroup{}

func read(ch chan bool, i int) {
	fmt.Printf("goroutine_num:%d , go func:%d\n", runtime.NumGoroutine(), i)
	<-ch
	wg.Done()
}

func main() {
	// 模拟用户请求数量
	requestCount := 10
	fmt.Println("goroutine_num:", runtime.NumGoroutine())
	// 限制最大并发数量3
	ch := make(chan bool, 3)
	for i := 0; i < requestCount; i++ {
		wg.Add(1)
		ch <- true
		go read(ch, i)
	}
	wg.Wait()
}
