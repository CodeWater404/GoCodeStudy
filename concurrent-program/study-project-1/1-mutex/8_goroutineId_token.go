package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
  @author: CodeWater
  @since: 2023/11/20
  @desc: 通过token标记一个goroutine
	6、7是用 goroutine id 做 goroutine 的标识，我们也可以让 goroutine 自己来提供标识。不管怎么说，Go 开发者不期望你利用 goroutine id 做一些不确定的东西，所
	以，他们没有暴露获取 goroutine id 的方法。下面的代码是第二种方案。调用者自己提供一个 token，获取锁的时候把这个 token 传入，释放锁的时候也需要把这个 token 传
	入。通过用户传入的 token 替换方案一中goroutine id，其它逻辑和方案一一致。
**/

// TokenRecursiveMutex 带token的可重入锁
type TokenRecursiveMutex struct {
	sync.Mutex
	token     int64
	recursion int32
}

// Lock 加锁，需要传入token
func (m *TokenRecursiveMutex) Lock(token int64) {
	if atomic.LoadInt64(&m.token) == token {
		m.recursion++
		return
	}
	// 传入的token和持有锁的token不一致，说明不是第一次调用
	m.Mutex.Lock()
	// 抢到锁之后记录这个token
	atomic.StoreInt64(&m.token, token)
	m.recursion = 1
	m.Mutex.Unlock()
}

// Unlock 解锁，需要传入token
func (m *TokenRecursiveMutex) Unlock(token int64) {
	if atomic.LoadInt64(&m.token) != token { // 非持有锁的goroutine尝试释放锁，错误的使用
		panic(fmt.Sprintf("wrong the owner(%d): %d", m.token, token))
	}
	m.recursion--         // 当前持有这个锁的token释放锁
	if m.recursion != 0 { // 如果这个goroutine还没有完全释放，则直接返回
		return
	}
	atomic.StoreInt64(&m.token, 0) // 此goroutine最后一次调用，需要释放锁,没有递归了
	m.Mutex.Unlock()
}

func main() {

}
