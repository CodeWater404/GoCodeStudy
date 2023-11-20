package main

import "sync"

/**
  @author: CodeWater
  @since: 2023/11/20
  @desc: 用mutex实现一个线程安全的slice
	 Mutex 经常会和其他非线程安全（对于 Go 来说，我们其
	实指的是 goroutine 安全）的数据结构一起，组合成一个线程安全的数据结构。新数据结
	构的业务逻辑由原来的数据结构提供，而 Mutex 提供了锁的机制，来保证线程安全。
	比如队列，我们可以通过 Slice 来实现，但是通过 Slice 实现的队列不是线程安全的，出队
	（Dequeue）和入队（Enqueue）会有 data race 的问题。这个时候，Mutex 就要隆重
	出场了，通过它，我们可以在出队和入队的时候加上锁的保护。
	因为标准库中没有线程安全的队列数据结构的实现，所以，你可以通过 Mutex 实现一个简
	单的队列。通过 Mutex 我们就可以为一个非线程安全的 data interface{}实现线程安全的
	访问
**/

// SliceQueue 线程安全的队列
type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

// NewSliceQueue 创建一个队列
func NewSliceQueue(n int) *SliceQueue {
	return &SliceQueue{
		data: make([]interface{}, 0, n),
	}
}

// Enqueue 入队
func (q *SliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	q.data = append(q.data, v)
	q.mu.Unlock()
}

// Dequeue 出队,如果队列为空，返回nil;否则返回队头元素
func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	if len(q.data) == 0 {
		q.mu.Unlock()
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mu.Unlock()
	return v
}
