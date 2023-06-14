package offer

/**
  @author: CodeWater
  @since: 2023/6/14
  @desc: 用两个栈实现队列
**/

type CQueue struct {
	inStack, outStack []int
}

func Constructor1() CQueue {
	return CQueue{}
}

// AppendTail 添加元素到队尾
func (this *CQueue) AppendTail(value int) {
	this.inStack = append(this.inStack, value)
}

//删除对头：输入栈全部入到输出栈，输出栈删除栈顶
func (this *CQueue) DeleteHead() int {
	//输出栈没有元素时才需要从输入栈拿元素
	if len(this.outStack) == 0 {
		//输入输出栈都为空的情况，队列相当于没有元素
		if len(this.inStack) == 0 {
			return -1
		}
		//输入栈元素全部放入到输出栈中
		this.in2out()
	}
	//出栈
	value := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return value
}

//互换元素
func (this *CQueue) in2out() {
	for len(this.inStack) > 0 {
		this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
		this.inStack = this.inStack[:len(this.inStack)-1]
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
