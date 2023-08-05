package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/8/5
  @desc:  让goroutine停止的几种方法
	1. 使用全局变量（本例）
		缺点：单个文件的时候还行；但是一旦变成工程化的时候就很容易冲突
	2. 使用channel
		缺点：值传来传去，不确定。
	3. 使用context，用法本质类似channel里面传一个空结构体
**/

var wg2 sync.WaitGroup

func worker21(exitChan chan struct{}) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			//退出到标记点
			break LOOP
		default:

		}
	}
	wg2.Done()
}

func main() {
	var exitChan = make(chan struct{})
	wg2.Add(1)
	go worker21(exitChan)
	time.Sleep(5 * time.Second)
	//让子goroutine发送退出信号
	exitChan <- struct{}{}
	close(exitChan)
	wg2.Wait()
	fmt.Println("over: use channel to close goroutine")

}
