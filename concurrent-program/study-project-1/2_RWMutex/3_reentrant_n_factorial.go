package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/21
  @desc: 读写锁的重入导致死锁
	情况三：n阶乘的例子
**/

// factorial 递归调用，计算 n 的阶乘
func factorial(m *sync.RWMutex, n int) int {
	if n < 1 { // 递归终止条件
		return 0
	}
	fmt.Println("rlock")
	m.RLock()
	defer func() {
		fmt.Println("RUnlock")
		m.RUnlock()
	}()
	time.Sleep(100 * time.Millisecond)
	return factorial(m, n-1) * n // 递归调用
}

func main() {
	var mu sync.RWMutex

	// wtiter，稍微等到，然后制造一个调用Lock的场景
	go func() {
		time.Sleep(200 * time.Millisecond)
		mu.Lock()
		fmt.Println("lock")
		time.Sleep(time.Millisecond * 100)
		mu.Unlock()
		fmt.Println("unlock")
	}()

	// reader先计算，一直在可重入，后面writer执行又在等reader执行完，
	// 而这时候reader又在重入导致有新的reader请求锁，需要等前面的writer
	// 执行；环路等待，死锁
	go func() {
		factorial(&mu, 10) // 计算10 的阶乘
	}()

	select {}
}
