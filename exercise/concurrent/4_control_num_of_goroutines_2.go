package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
  @author: CodeWater
  @since: 2024/1/14
  @desc: 控制并发数量2
	无缓冲channel
**/

var wg sync.WaitGroup

func read(ch chan bool, i int) {
	for _ = range ch {
		fmt.Printf("goroutine_num: %d , go func: %d\n", runtime.NumGoroutine(), i)
		wg.Done()
	}
}

func main() {
	// 模拟用户请求数量
	requestCount := 10
	fmt.Println("goroutine_num:", runtime.NumGoroutine())
	ch := make(chan bool)
	//手动限制，下面的for循环会阻塞，直到ch中有数据被读取
	for i := 0; i < 3; i++ {
		go read(ch, i)
	}

	for i := 0; i < requestCount; i++ {
		wg.Add(1)
		ch <- true
	}
	wg.Wait()
}
