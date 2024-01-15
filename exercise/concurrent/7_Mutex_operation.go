package main

import (
	"fmt"
	"sync"
)

/**
  @author: CodeWater
  @since: 2024/1/15
  @desc: 互斥锁mutex操作
	sync.Mutex(互斥锁)可以限制对临界资源的访问，保证只有一个goroutine访问共享资源
	使用场景：大虽读写，比如多个goroutine并发更新同一个资源，像计数器
**/

type Counter struct {
	mu    sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	var gNum = 1000
	wg.Add(gNum)
	for i := 0; i < gNum; i++ {
		go func() {
			defer wg.Done()
			counter.Incr()
		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
}
