package _81_300

import "container/heap"

/**
  @author: CodeWater
  @since: 2024/2/21
  @desc: 295. 数据流的中位数
**/

type minHeap []int
type maxHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type MedianFinder struct {
	up   minHeap
	down maxHeap
}

func Constructor() MedianFinder {
	a, b := &minHeap{}, &maxHeap{}
	heap.Init(a)
	heap.Init(b)
	return MedianFinder{*a, *b}
}

// 对顶堆：左半部分用大根堆存储，右半部分用小根堆存储。求中位数直接比较两个堆的个数和堆顶
func (m *MedianFinder) AddNum(num int) {
	if m.down.Len() == 0 || num <= m.down[0] {
		heap.Push(&m.down, num)
		if m.down.Len() > m.up.Len()+1 {
			heap.Push(&m.up, heap.Pop(&m.down).(int))
		}
	} else {
		heap.Push(&m.up, num)
		if m.up.Len() > m.down.Len() {
			heap.Push(&m.down, heap.Pop(&m.up).(int))
		}
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if (m.down.Len()+m.up.Len())%2 > 0 {
		return float64(m.down[0])
	}
	return float64(m.down[0]+m.up[0]) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
