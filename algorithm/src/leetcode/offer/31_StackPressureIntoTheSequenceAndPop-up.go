package offer

/**
  @author: CodeWater
  @since: 2023/6/16
  @desc: 栈的压入、弹出序列
**/

func validateStackSequences(pushed []int, popped []int) bool {
	//  开一个辅助栈，放入压入栈的元素
	st := []int{}
	//遍历弹出序列的
	j := 0
	for _, x := range pushed {
		st = append(st, x)
		//辅助栈不空并且栈顶元素等于弹出的元素，那就一直把辅助栈顶元素弹出，同时更新弹出序列j
		for len(st) > 0 && st[len(st)-1] == popped[j] {
			st = st[:len(st)-1]
			j++
		}

	}
	//看最后辅助栈是否还有元素
	return len(st) == 0
}
