package main

import (
	"fmt"
	"sync/atomic"
)

/**
  @author: CodeWater
  @since: 2024/1/15
  @desc: 原子操作,最基本的并发术语
**/

var opts int64 = 0

func add(addr *int64, delta int64) {
	atomic.AddInt64(addr, delta)
	fmt.Println("add opts:", *addr)
}

func load(addr *int64) {
	fmt.Println("load opts:", atomic.LoadInt64(addr))
}

func compareAndSwap(addr *int64, oldValue int64, newValue int64) {
	if atomic.CompareAndSwapInt64(addr, oldValue, newValue) {
		fmt.Println("cas opts:", *addr)
		return
	}
}

func swap(addr *int64, newValue int64) {
	atomic.SwapInt64(addr, newValue)
	fmt.Println("swap opts:", *addr)
}

func store(addr *int64, newValue int64) {
	atomic.StoreInt64(addr, newValue)
	fmt.Println("stort opts:", *addr)
}

func main() {
	add(&opts, 3)
	load(&opts)
	compareAndSwap(&opts, 3, 4)
	swap(&opts, 5)
	store(&opts, 6)
}
