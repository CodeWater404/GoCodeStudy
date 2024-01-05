package _41_160

/**
  @author: CodeWater
  @since: 2024/1/5
  @desc: 155. 最小栈
**/

type MinStack struct {
	stk, temp []int //temp保存最小值
}

func Constructor() MinStack {
	return MinStack{
		stk:  make([]int, 0),
		temp: make([]int, 0),
	}
}

func (ms *MinStack) Push(val int) {
	ms.stk = append(ms.stk, val)
	if len(ms.temp) == 0 || ms.temp[len(ms.temp)-1] >= val {
		ms.temp = append(ms.temp, val)
	}
}

func (ms *MinStack) Pop() {
	//stk栈顶元素小于等于temp栈顶的是否，temp也要弹出
	if ms.stk[len(ms.stk)-1] <= ms.temp[len(ms.temp)-1] {
		ms.temp = ms.temp[:len(ms.temp)-1]
	}
	ms.stk = ms.stk[:len(ms.stk)-1]
}

func (ms *MinStack) Top() int {
	return ms.stk[len(ms.stk)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.temp[len(ms.temp)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
