package main

import (
	"sync/atomic"
	"unsafe"
)

/**
  @author: CodeWater
  @since: 2023/11/26
  @desc: $
**/

// LKQueue lock-free的queue
type LKQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

// node 通过链表实现，这个数据结构代表链表中的节点
type node struct {
	value interface{}
	next  unsafe.Pointer
}

// NewLKQueue 创建一个LKQueue，头指针所指向的是个空节点，不存储数据，只是为了方便操作
func NewLKQueue() *LKQueue {
	n := unsafe.Pointer(&node{})
	return &LKQueue{head: n, tail: n}
}

// Enqueue 入队
func (q *LKQueue) Enqueue(v interface{}) {
	n := &node{value: v}
	for {
		tail := load(&q.tail)
		next := load(&tail.next)
		if tail == nil { // 尾还是尾
			if cas(&tail.next, next, n) { // 增加到队尾
				cas(&q.tail, tail, n) // 入队成功，移动尾巴指针
				return
			}
		} else { // 已有新数据加到队列的后面的，需要移动尾指针
			cas(&q.tail, tail, next)
		}
	}
}

// Dequeue 出队，没有元素返回nil
func (q *LKQueue) Dequeue() interface{} {
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)
		if head == load(&q.head) { // head还是那个head
			if head == tail { // head和tail一样
				if next == nil { // 说明是空队列
					return nil
				}
				// 只是尾指针还没有调整，尝试调整它指向下一个
				cas(&q.tail, tail, next)
			} else {
				// 读取出队的数据
				v := next.value
				//出队，头指针移动到下一个
				if cas(&q.head, head, next) {
					return v
				}
			}
		}
	}
}

// load 将unsafe.Pointer原子加载为node
func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

// cas 封装cas，避免直接将*node转为unsafe.Pointer
func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}
