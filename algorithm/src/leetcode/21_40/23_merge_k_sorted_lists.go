package _1_40

import "container/heap"

/**
  @author: CodeWater
  @since: 2024/2/13
  @desc: 23. 合并K个升序链表
**/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int {
	return len(h)
}

func (h ListNodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ListNodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 比较k个最小值，然后放到最前面。比较这个操作用堆来维护
func mergeKLists(lists []*ListNode) *ListNode {
	h := &ListNodeHeap{}
	heap.Init(h)
	dummy := &ListNode{Val: -1}
	tail := dummy
	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	for h.Len() > 0 {
		// 弹出最小的数，放到tail后面
		t := heap.Pop(h).(*ListNode)
		tail.Next = t
		tail = tail.Next
		if t.Next != nil {
			// 同时把弹出的这个列表的下一个元素放入堆中，然后进行下一轮的比较
			heap.Push(h, t.Next)
		}
	}
	return dummy.Next
}
