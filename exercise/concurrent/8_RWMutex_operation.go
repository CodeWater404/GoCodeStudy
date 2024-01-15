package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2024/1/15
  @desc: RWMutex操作
	sync.RWMutex(读写锁)可以限制对临界资源的访问，保证只有一个goroutine写共享资源，可以有多个goroutine读共享资源
	使用场景：大量并发读，少量并发写，有强烈的性能要求
**/

type Counter struct {
	mu      sync.RWMutex
	counter uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.counter++
	c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counter
}

func main() {
	var counter Counter
	var gNum = 1000
	for i := 0; i < gNum; i++ {
		go func() { // read
			fmt.Println("read: ", counter.Count())
		}()
	}
	for { // write
		counter.Incr()
		fmt.Println("incr....")
		time.Sleep(time.Second * 1)
	}
}
