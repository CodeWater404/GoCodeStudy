package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/8/5
  @desc: 让goroutine停止的几种方法
	1. 使用全局变量（本例）
		缺点：单个文件的时候还行；但是一旦变成工程化的时候就很容易冲突
	2. 使用channel
		缺点：值传来传去，不确定。
	3. 使用context，用法本质类似channel里面传一个空结构体
**/
var wg1 sync.WaitGroup

func worker11() {
	defer wg1.Done()
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
}

//waitgroup的基本使用
func runWorker11() {
	wg1.Add(1)
	go worker11()
	wg1.Wait()
	fmt.Println("worker1 over")
}

var exit bool

//使用全局变量让goroutine停止
func worker12() {
	defer wg1.Done()
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		//break之后，函数就执行换成，相对应的，goroutine就会退出
		if exit {
			break
		}
	}
}

func runWorker12() {
	wg1.Add(1)
	go worker12()
	time.Sleep(5 * time.Second)
	exit = true //要在wait之前改边exit值，不然goroutine里面一直是false就会一直阻塞执行，不会停止
	wg1.Wait()
	fmt.Println("worker2 over")

}

func main() {
	//runWorker11()
	runWorker12()
}
