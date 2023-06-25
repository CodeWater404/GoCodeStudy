package offer

import (
	"container/heap"
	"sort"
)

/**
  @author: CodeWater
  @since: 2023/6/26
  @desc: 数据流中的中位数
**/
type MedianFinder struct {
	//queMin和queMax，它们分别用于存储较小的一半数和较大的一半数。
	queMin, queMax hp
}

/** initialize your data structure here. */
func Constructor4() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	minQ, maxQ := &mf.queMin, &mf.queMax
	//最小优先队列长度为0或者当前数比队头的负数还小
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {

		heap.Push(minQ, -num)
		//// 如果最大优先队列的长度小于最小优先队列的长度减1，则从最小优先队列中弹出队头并将其加入最大优先队列
		if maxQ.Len()+1 < minQ.Len() {
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		// 如果最大优先队列的长度大于最小优先队列的长度，则从最大优先队列中弹出队头并将其加入最小优先队列
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := mf.queMin, mf.queMax
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	}

	return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
}

type hp struct {
	sort.IntSlice
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}
func (h *hp) Pop() (v interface{}) {
	a := h.IntSlice
	v = a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
