package offer

import "math"

/**
  @author: CodeWater
  @since: 2023/6/14
  @desc: 包含min函数的栈
**/

//method a================================================================
type MinStack struct {
	stack []int
}

/** initialize your data structure here. */
func Constructor2() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	min := math.MaxInt32
	for i := len(this.stack) - 1; i >= 0; i-- {
		if this.stack[i] < min {
			min = this.stack[i]
		}
	}
	return min
}

//method b: 辅助栈================================================================

type MinStack2 struct {
	stack []int
	//辅助栈，记录当前stack中min的数
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack2 {
	return MinStack2{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack2) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	//这里一直拿辅助栈栈顶和当前入栈元素比较，栈顶元素更小说明之前就有最小值，当前x还不够小
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack2) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack2) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack2) Min() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
