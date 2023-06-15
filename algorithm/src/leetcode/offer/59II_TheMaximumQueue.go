package offer

/**
  @author: CodeWater
  @since: 2023/6/15
  @desc: 队列的最大值
**/
type MaxQueue struct {
	queue []int
	//辅助优先队列，从大到小。因为题目要求的是max
	maxQueue []int
}

func Constructor3() MaxQueue {
	return MaxQueue{
		queue:    []int{},
		maxQueue: []int{0},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.maxQueue[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	//优先队列不空，并且插入的值从优先队尾开始比较直到小于前一个元素，否则队尾一直弹出元素,优先队列为了保存从大到小,中间一部分元素被提前弹出了
	for len(this.maxQueue) != 0 && value > this.maxQueue[len(this.maxQueue)-1] {
		this.maxQueue = this.maxQueue[:len(this.maxQueue)-1]
	}
	this.maxQueue = append(this.maxQueue, value)
}

func (this *MaxQueue) Pop_front() int {
	n := -1
	if len(this.queue) != 0 {
		n = this.queue[0]
		this.queue = this.queue[1:]
		//弹出的元素值跟优先队头一样大，优先队头也要弹出：如果队列与优先队列，队头不相同，说明这个值在优先队列中已经被弹出了，就无需再弹出,也就是在push的时候
		if this.maxQueue[0] == n {
			this.maxQueue = this.maxQueue[1:]
		}
	}
	return n
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */