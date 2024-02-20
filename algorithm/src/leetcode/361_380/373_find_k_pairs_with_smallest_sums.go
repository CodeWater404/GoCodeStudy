package _61_380

import "container/heap"

/**
  @author: CodeWater
  @since: 2024/2/20
  @desc: 373. 查找和最小的K对数字
**/
// [][0]:从两个数组中选取的元素和；[][1]:第一个数组选取元素的下标；[][2]第二个数组选取元素的下标
type pairHeap [][3]int

func (h pairHeap) Len() int { return len(h) }

// 小根堆，比较和最小用第一个位置的比
func (h pairHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h pairHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(x interface{}) { *h = append(*h, x.([3]int)) }
func (h *pairHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

// 本质上是多路归并
func kSmallestPairs(a []int, b []int, k int) [][]int {
	if len(a) == 0 || len(b) == 0 {
		return [][]int{}
	}
	n, m := len(a), len(b)
	h := &pairHeap{}
	heap.Init(h)
	for i := 0; i < m; i++ {
		//把第二个数组每一个元素和第一个数组第一个元素的组合放入堆中
		heap.Push(h, [3]int{b[i] + a[0], 0, i})
	}
	res := make([][]int, 0)
	for k > 0 && h.Len() > 0 {
		t := heap.Pop(h).([3]int)
		// 选取一个最小的组合
		res = append(res, []int{a[t[1]], b[t[2]]})
		if t[1]+1 < n {
			// 如果第一个数组还有元素，把第一个数组的下一个元素和第二个数组当前元素组合放入堆
			heap.Push(h, [3]int{a[t[1]+1] + b[t[2]], t[1] + 1, t[2]})
		}
		k--
	}
	return res
}
