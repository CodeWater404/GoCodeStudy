package _521_2540

import (
	"container/heap"
	"math"
)

/**
  @author: CodeWater
  @since: 2023/7/7
  @desc: 过桥的时间
**/

func findCrossingTime(n int, k int, time [][]int) int {
	var waitLeft, waitRight, workLeft, workRight queue
	var remain, curTime = n, 0
	for i, t := range time {
		heap.Push(&waitLeft, &item{id: i, priority: -(t[0]+t[2])*1001 - i})
	}

	for remain > 0 || workRight.Len() > 0 || waitRight.Len() > 0 {
		for workLeft.Len() > 0 && workLeft.Peek().time <= curTime {
			item := heap.Pop(&workLeft).(*item)
			item.priority = -(time[item.id][0]+time[item.id][2])*1001 - item.id
			heap.Push(&waitLeft, item)
		}

		for workRight.Len() > 0 && workRight.Peek().time <= curTime {
			item := heap.Pop(&workRight).(*item)
			item.priority = -(time[item.id][0]+time[item.id][2])*1001 - item.id
			heap.Push(&waitRight, item)
		}
		if waitRight.Len() > 0 {
			item := heap.Pop(&waitRight).(*item)
			curTime += time[item.id][2]
			item.priority = (curTime+time[item.id][3])*1001 + item.id
			item.time = curTime + time[item.id][3]
			heap.Push(&workLeft, item)
		} else if remain > 0 && waitLeft.Len() > 0 {
			item := heap.Pop(&waitLeft).(*item)
			curTime += time[item.id][0]
			item.priority = (curTime+time[item.id][1])*1001 + item.id
			item.time = curTime + time[item.id][1]
			heap.Push(&workRight, item)
			remain--
		} else {
			nextTime := math.MaxInt
			if workLeft.Len() > 0 {
				nextTime = min(nextTime, workLeft.Peek().time)
			}
			if workRight.Len() > 0 {
				nextTime = min(nextTime, workRight.Peek().time)
			}
			if nextTime != math.MaxInt {
				curTime = max(nextTime, curTime)
			}
		}
	}
	return curTime
}

type item struct {
	id       int
	priority int
	time     int
}

type queue []*item

func (q queue) Len() int {
	return len(q)
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q queue) Less(i, j int) bool {
	return q[i].priority < q[j].priority
}

func (pq *queue) Push(x any) {
	item := x.(*item)
	*pq = append(*pq, item)
}

func (pq *queue) Pop() any {
	n := pq.Len() - 1
	item := (*pq)[n]
	*pq = (*pq)[0:n]
	return item
}

func (pq *queue) Peek() *item {
	return (*pq)[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
