package main

import (
	"fmt"
	"github.com/petermattis/goid"
	"sync"
	"sync/atomic"
)

/**
  @author: CodeWater
  @since: 2023/11/20
  @desc:  获取goroutineId
  方案二：通过底层暴力获取
		获取运行时的 g 指针，反解出对应的 g 的结构。每个运行的 goroutine 结构的g 指针保存在当前 goroutine 的一个叫做 TLS 对象中。
		第一步：我们先获取到 TLS 对象；
		第二步：再从 TLS 中获取 goroutine 结构的 g 指针；
		第三步：再从 g 指针中取出 goroutine id
	需要注意的是，不同 Go 版本的 goroutine 的结构可能不同，所以需要根据 Go 的不同
	版本进行调整。当然了，如果想要搞清楚各个版本的 goroutine 结构差异，所涉及的内容
	又过于底层而且复杂，学习成本太高,直接使用第三方的库来获取 goroutine id 就可以了。
**/

// RecursiveMutex 包装一个mutex，实现可重入
// 给Mutex 打一个补丁，解决了记录锁的持有者的问题。用 owner 字段，记
// 录当前锁的拥有者 goroutine 的 id；recursion 是辅助字段，用于记录重入的次数。
// 尽管拥有者可以多次调用 Lock，但是也必须调用相同次数的Unlock，这样
// 才能把锁释放掉。这是一个合理的设计，可以保证 Lock 和 Unlock 一一对应
type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // 当前持有锁的goroutine id
	recursion int32 // 这个goroutine 重入的次数
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()

	fmt.Printf("===>Lock goroutine id:%d\n", gid)

	// 如果当前持有锁的goroutine就是这次调用的goroutine，说明是重入
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// 获得锁的goroutine第一次调用，记录下它的goroutine id,调用次数加1
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()

	fmt.Printf("===>Unlock goroutine id:%d\n", gid)

	// 非持有锁的goroutine尝试释放锁，错误的使用
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d", m.owner, gid))
	}
	// 调用次数减1
	m.recursion--
	if m.recursion != 0 { // 如果这个goroutine还没有完全释放，则直接返回
		return
	}
	// 此goroutine最后一次调用，需要释放锁.原子地将 m.owner 的值设置为 -1。
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

// foo 和 bar 函数都会调用 locker 的 Lock 方法，foo拥有锁，然后在bar中又请求这一把锁，重入锁了。go没有提供重入锁，所以会死锁。
// 这里为了演示获取goroutine id
func foo(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	defer l.Unlock()
}

func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

func main() {
	foo(&RecursiveMutex{})
}
