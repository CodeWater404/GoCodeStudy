package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
)

/**
  @author: CodeWater
  @since: 2024/1/14
  @desc: go实现可重入锁（本身不支持）
**/

type ReentrantLock struct {
	sync.Mutex
	recursion int32 //记录重入次数
	owner     int64 //记录当前持有锁的goroutine的id
}

// GetGoroutineID 获取当前持有锁的goroutine的id
func GetGoroutineID() int64 {
	var buf [64]byte
	/*stack获取到的：
	goroutine 1 [running]:
	main.GetGoroutineID()
	        F:/Code/GoCode/exe
	*/
	var s = buf[:runtime.Stack(buf[:], false)]

	s = s[len("goroutine"):]
	a := bytes.IndexByte(s, ' ')
	s = s[1 : a+2] //s[1:3]，即取出goroutine id
	gid, _ := strconv.ParseInt(string(s), 10, 64)
	return gid
}

func NewReentrantLock() sync.Locker {
	res := &ReentrantLock{
		Mutex:     sync.Mutex{},
		recursion: 0,
		owner:     0,
	}
	return res
}

// ReentrantMutex 可重入锁
type ReentrantMutex struct {
	sync.Mutex
	owner     int64 //当前持有锁的goroutine的id
	recursion int32 //重入次数
}

func (m *ReentrantMutex) Lock() {
	gid := GetGoroutineID()
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

func (m *ReentrantMutex) Unlock() {
	gid := GetGoroutineID()
	// 非持有锁的goroutine尝试释放锁，错误的使用
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d):%d!", m.owner, gid))
	}
	// 调用次数减1
	m.recursion--
	if m.recursion != 0 { // 如果这个goroutine还没有完全释放，则直接返回
		return
	}
	// 此goroutine最后一次调用，需要释放锁
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

func main() {
	var mutex = &ReentrantMutex{}
	mutex.Lock()
	mutex.Lock()
	fmt.Println("lock before =============== lock after")
	mutex.Unlock()
	mutex.Unlock()
}
