package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2024/1/13
  @desc: 控制并发顺序
	采用channel
**/

var wg sync.WaitGroup

func print(goRoutingStr string, inputChan chan struct{}, outputChan chan struct{}) {
	// 模拟业务处理
	time.Sleep(time.Second * 1)
	select {
	case <-inputChan:
		fmt.Println("This is ", goRoutingStr)
		outputChan <- struct{}{}
	}
	wg.Done()
}

func main() {
	ch1, ch2, ch3 := make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)
	ch1 <- struct{}{}
	wg.Add(3)
	go print("go 1", ch1, ch2)
	go print("go 2", ch2, ch3)
	go print("go 3", ch3, ch1)
	wg.Wait()

}
