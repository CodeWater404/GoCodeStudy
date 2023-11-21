package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/21
  @desc: wait-group例子
	启动了 10 个 worker，分别对计数值加一，10 个 worker 都完成后，我们期望输出计数器的值。
**/

// Counter 线程安全的计数器
type Counter struct {
	mu    sync.Mutex
	count uint64
}

// Incr 计数值加1
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// Count 获取计数值
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// worker
func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go worker(&counter, &wg)
	}
	// 检查点，等待所有的worker都完成任务
	wg.Wait()
	fmt.Println("count:", counter.Count())
}
