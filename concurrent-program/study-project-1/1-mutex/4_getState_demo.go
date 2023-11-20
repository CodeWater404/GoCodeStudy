package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

/**
  @author: CodeWater
  @since: 2023/11/20
  @desc: mutex：扩展额外功能
  2. 获取等待者的数量等指标:(这里的获取指标依赖go版本，如果go底层mutex的state各个位的含义变化，那么这里的获取指标也会变化)
	Mutex 结构中的 state 字段有很多个含义，通过 state 字段，可以知道锁是否已经被某
	个 goroutine 持有、当前是否处于饥饿状态、是否有等待的 goroutine 被唤醒、等待者的
	数量等信息。但是，state 这个字段并没有暴露出来，所以，我们需要想办法获取到这个字段，并进行解析.
	可以通过 unsafe 的方式实现.
	state 这个字段的第一位是用来标记锁是否被持有，第二位用来标记是否已经唤醒了一个等
	待者，第三位标记锁是否处于饥饿状态，通过分析这个 state 字段我们就可以得到这些状
	态信息。我们可以为这些状态提供查询的方法，这样就可以实时地知道锁的状态了。
**/

// 复制Mutex定义的常量
const (
	mutexLocked      = 1 << iota // 加锁标识位置,iota=0，1左移0位，结果为1。1表示加锁，0表示未加锁
	mutexWoken                   // 唤醒标识位置，2：1左移1位，用上面的表达式计算出的
	mutexStarving                // 锁饥饿标识位置，4：1左移2位，用上面的表达式计算出的
	mutexWaiterShift = iota      // 标识waiter的起始bit位置,3:使用的是iota递增的值
)

// Mutex 扩展额外功能
type Mutex struct {
	sync.Mutex
}

// Count 当前持有和等待这把锁的 goroutine 的总数。
// 通过 unsafe 操作，我们可以得到 state 字段的值。然后，通过右移操作，将 state 的值右移
// 三位（这里的常量 mutexWaiterShift 的值为 3），就得到了当前等待者的数量。如果当前
// 的锁已经被其他 goroutine 持有，那么，我们就稍微调整一下这个值，加上一个 1
func (m *Mutex) Count() int {
	// 获取 state 字段的值
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> mutexWaiterShift //得到等待者的数量
	v = v + (v & mutexLocked) //再加上锁持有者的数量，0或者1
	return int(v)
}

// IsLocked 锁是否被持有
func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	/*
		mutexLocked 的值是1，它的二进制表示是 0001（假设我们只考虑4位二进制数）。所以，只有 state 的最低位是1时，
		state & mutexLocked 的结果才会是1。因此，state & mutexLocked == mutexLocked 这个表达式的意思是，如
		果 state 的最低位是1，那么锁就被持有（因为 mutexLocked 的值是1）。如果state 的最低位是0，那么锁就没有被持有。
		下面唤醒和饥饿同理
	*/
	return state&mutexLocked == mutexLocked
}

// IsWoken 判断是否有等待者被唤醒
func (m *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

// IsStarving 锁是否处于饥饿状态
func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}

// 测试一下，比如，在 1000 个 goroutine 并发访问的情况下，我们可
// 以把锁的状态信息输出出来：
func main() {
	var mu Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}

	time.Sleep(time.Second)
	/*
		有一点需要注意一下，在获取 state 字段的时候，并没有通过 Lock 获取这把锁，所以获
		取的这个 state 的值是一个瞬态的值，可能在解析出这个字段之后，锁的状态已经发生
		了变化。不过没关系，因为查看的就是调用的那一时刻的锁的状态。
	*/
	fmt.Printf("waitings: %d , isLocked: %t , isWorken: %t , isStarving: %t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}
