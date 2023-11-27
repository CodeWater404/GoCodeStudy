package main

import (
	"fmt"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/27
  @desc: 用channel实现锁
	先初始化一个 capacity 等于 1
的 Channel，然后再放入一个元素。这个元素就代表锁，谁取得了这个元素，就相当于获
取了这把锁
**/

// Mutex 是一个互斥锁
type Mutex struct {
	ch chan struct{}
}

// NewMutex 创建一个Mutex
func NewMutex() *Mutex {
	mu := &Mutex{make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}

// Lock 请求锁，直到获取锁
func (m *Mutex) Lock() {
	<-m.ch
}

// Unlock 释放锁
func (m *Mutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock of unlocked mutex")
	}
}

// TryLock 尝试获取锁
func (m *Mutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:

	}
	return false
}

// LockTimeout 加入一个超时的设置
func (m *Mutex) LockTimeout(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:
	}
	return false
}

// IsLocked 判断锁时候已被持有
func (m *Mutex) IsLocked() bool {
	return len(m.ch) == 0
}

func main() {
	m := NewMutex()
	ok := m.TryLock()
	fmt.Printf("locked v %v\n", ok)
	// 第一次能够获取到锁，所以这次就获取不到 ;要想获取到，可以用m.Unlock()释放掉
	ok = m.TryLock()
	fmt.Printf("locked v %v\n", ok)
}
