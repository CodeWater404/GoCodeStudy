package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/**
  @author: CodeWater
  @since: 2024/1/15
  @desc: cond条件变量
	sync.Cond可以让一组的Goroutine都在满足特定条件时被唤醒
	使用场景：利用等待/通知机制实现阻塞或者唤醒
**/

var status int64

func broadcast(c *sync.Cond) {
	c.L.Lock()
	atomic.StoreInt64(&status, 1)
	c.Signal()
	c.L.Unlock()
}

func listen(c *sync.Cond) {
	c.L.Lock()
	for atomic.LoadInt64(&status) != 1 {
		c.Wait()
	}
	fmt.Println("listen")
	c.L.Unlock()
}

func main() {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c)
	}
	time.Sleep(time.Second * 1)
	go broadcast(c)
	time.Sleep(time.Second * 1)
}
