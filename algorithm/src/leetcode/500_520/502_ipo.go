package _00_520

import (
	"container/heap"
	"sort"
)

/**
  @author: CodeWater
  @since: 2024/2/19
  @desc: 502. IPO
**/

type pair struct {
	first, second int
}

type pairList []pair

func (p pairList) Len() int           { return len(p) }
func (p pairList) Less(i, j int) bool { return p[i].first < p[j].first }
func (p pairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	q := make(pairList, 0, n)
	for i := 0; i < n; i++ {
		q = append(q, pair{capital[i], profits[i]})
	}
	// 需要实现Len，Less，Swap方法
	sort.Sort(q)
	h := &maxHeap{} //大根堆
	heap.Init(h)
	i := 0
	for ; k > 0; k-- { // 最多选k次
		for i < n && q[i].first <= w {
			// 把成本小于利润w的放入堆中自动排序，利润w会扩大，所以每次选取都需要遍历一遍q存储的成本
			heap.Push(h, q[i].second)
			i++
		}
		if h.Len() == 0 {
			break
		}
		//在所有小于利润个数中，选取最大的那个
		w += heap.Pop(h).(int)

	}
	return w
}
