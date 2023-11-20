package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

/**
  @author: CodeWater
  @since: 2023/11/19
  @desc: mutex：扩展额外功能
  1. 重写获取锁的机制：尝试获取锁，如果能够抢到锁，返回true;否则返回false。（原来的获取机制是：能拿到就拿，拿不到就阻塞等待直到拿到）
  在并发编程中，互斥锁（Mutex）是一种常用的同步机制，用于保护共享资源的访问。互斥锁有几种状态，包括加锁、唤醒和饥饿：
	a. 加锁（Locked）：当一个线程成功获取到互斥锁时，我们说互斥锁处于加锁状态。在这种状态下，其他试图获取锁的线程将会被阻塞，
	直到锁被释放
	b. 唤醒（Woken）：当一个等待锁的线程被唤醒（即被通知可以尝试获取锁）时，我们说互斥锁处于唤醒状态。这通常发生
	在锁被释放时，系统会从等待队列中唤醒一个或多个线程来尝试获取锁。
	c. 饥饿（Starving）：当一个线程长时间等待获取锁，但总是被其他线程抢先获取，我们说这个线程处于饥饿状态。为了防止饥饿，一些
	互斥锁的实现会提供公平锁机制，即按照线程到达的顺序分配锁，这样可以保证每个线程最终都能获取到锁。在下面代码中，mutexLocked、
	mutexWoken 和 mutexStarving 是用来表示互斥锁的这三种状态的常量。它们的值是通过位操作得到的，这样可以在一个 int32 变量中
	同时存储多个状态。
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

// TryLock 尝试获取锁
// 如果能够抢到锁，返回true;否则返回false
func (m *Mutex) TryLock() bool {
	/* 如果能成功抢到锁
	atomic.CompareAndSwapInt32 是一个原子操作，它会比较一个 int32 类型的值是否等于预期值，如果等于预期值，
	就将其替换为新值。在代码中，(*int32)(unsafe.Pointer(&m.Mutex)) 是将 m.Mutex 的
	地址转换为 *int32 类型的指针，然后传递给 CompareAndSwapInt32 函数。0 是预期值，mutexLocked 是新值。
	如果 m.Mutex 的值等于 0（表示未锁定），那么这个函数就会将其值设置为 mutexLocked（表示已锁定），并返回
	true。如果 m.Mutex 的值不等于 0（表示已经被其他线程锁定），那么这个函数就不会做任何事情，并返回 false。
	*/
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}

	/* 尝试获取锁：如果锁已经被其他 goroutine 所持有，或者被其他唤醒的 goroutine 准备持有，那么，就直接返回 false，不再请求。
	atomic.LoadInt32 函数用于在并发修改的情况下安全地读取一个 int32 类型的值。在代码中，(*int32)(unsafe.Pointer(&m.Mutex))
	是将 m.Mutex 的地址转换为 *int32 类型的指针，然后传递给LoadInt32 函数。这个函数会返回 m.Mutex 的当前值m返回的值（old）
	然后与 mutexLocked=1、mutexStarving=4 和 mutexWoken=2的组合进行位与操作。如果结果不为零，那么表示互斥锁已经被锁定、处于饥饿
	状态或已被唤醒，函数返回 false，表示无法获取锁。这是一种非阻塞的获取锁的方式。如果无法立即获取锁，它不会等待，而是立即返回。
	这种方式在某些场景下可以提高性能，但也可能导致更高的 CPU 使用率。
	*/
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex))) // 获取锁的状态，转成int32类型，loadInt32是原子操作，安全读取
	if old&(mutexWoken|mutexLocked|mutexStarving) != 0 {
		return false
	}

	/* 尝试在竞争的状态下请求锁
	如果没有被持有，也没有其它唤醒的 goroutine 来竞争锁，锁也不处于饥饿状态，就尝试获取这把锁（第 29 行），不论是否成功都将结果返回。
	因为，这个时候，可能还有其他的goroutine 也在竞争这把锁，所以，不能保证成功获取这把锁。
	*/
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32(
		(*int32)(unsafe.Pointer(&m.Mutex)),
		old,
		new,
	)
}

// try 用于测试 TryLock 的效果
// 测试工作机制：程序运行时会启动一个 goroutine 持有这把我们自
// 己实现的锁，经过随机的时间才释放。主 goroutine 会尝试获取这把锁。如果前一个
// goroutine 一秒内释放了这把锁，那么，主 goroutine 就有可能获取到这把锁了，输
// 出“got the lock”，否则没有获取到也不会被阻塞，会直接输出“can't get the lock”。
func try() {
	var mu Mutex
	go func() { // 启动一个goroutine持有一段时间的锁
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second) // 随机睡眠一段时间
		mu.Unlock()
	}()

	time.Sleep(time.Second)

	ok := mu.TryLock() // 尝试获取到锁
	if ok {
		fmt.Println("got the lock")
		//mock do something
		mu.Unlock()
		return
	}

	// 没有获取到锁，输出提示信息
	fmt.Println("can't get the lock")
}

func main() {
	try()
}
