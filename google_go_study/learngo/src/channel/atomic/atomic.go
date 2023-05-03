package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/5/4
  @desc: traditional synchronization mechanism(传统的同步机制)
	1. use metux as example
	2. WaitGroup
	3. Cond
**/

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	//1.
	//a.lock.Lock()
	//defer a.lock.Unlock()
	//
	//a.value++

	//2. another way to write it: an area of code can be securely accessed using anonymous functions(可以使用匿名函数安全访问一段代码区)
	fmt.Println("safe increment")
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
