package main

import (
	"fmt"
	"sync"
)

/**
  @author: CodeWater
  @since: 2023/11/21
  @desc: 读写锁的重入导致死锁
	情况二
**/

func foo(l *sync.RWMutex) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l *sync.RWMutex) {
	fmt.Println("in bar")
	l.Lock()
	fmt.Println("in bar, get lock")
	l.Unlock()
}

func main() {
	//output:
	//in foo
	//in bar
	//fatal error: all goroutines are asleep - deadlock!
	//
	//goroutine 1 [sync.Mutex.Lock]:
	foo(&sync.RWMutex{})
}
