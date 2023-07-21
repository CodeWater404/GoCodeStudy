package _481_1500

import "container/heap"

/**
  @author: CodeWater
  @since: 2023/7/21
  @desc: 满足不等式的最大值
**/
type PriorityQueue [][]int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i][0] != pq[j][0] {
		return pq[i][0] < pq[j][0]
	}
	return pq[i][1] < pq[j][1]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() any {
	n, old := len(*pq), *pq
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func (pq PriorityQueue) Top() []int {
	return pq[0]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func findMaxValueOfEquation(points [][]int, k int) int {
	res := -0x3f3f3f3f
	pq := &PriorityQueue{}
	for _, point := range points {
		x, y := point[0], point[1]
		for pq.Len() > 0 && x-pq.Top()[1] > k {
			heap.Pop(pq)
		}
		if pq.Len() > 0 {
			res = max(res, x+y-pq.Top()[0])
		}
		heap.Push(pq, []int{x - y, x})
	}
	return res
}
