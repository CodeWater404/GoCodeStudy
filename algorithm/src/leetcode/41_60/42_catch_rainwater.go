package _1_60

/**
  @author: CodeWater
  @since: 2023/12/25
  @desc: 接雨水
**/
//单调栈
func trap(height []int) int {
	//stk存储的下标
	stk, res := []int{}, 0
	for i := 0; i < len(height); i++ {
		//last记录栈中栈顶元素的上一个元素高度值
		last := 0
		//栈非空 ； 栈顶元素小于等于当前遍历元素的高度
		for len(stk) > 0 && height[stk[len(stk)-1]] <= height[i] {
			//先加上栈顶元素和当前元素构成的凹槽的面积：高度（栈顶元素高度-上一个元素的高度） * 宽度（当前元素下标-栈顶元素-1）
			res += (height[stk[len(stk)-1]] - last) * (i - stk[len(stk)-1] - 1)
			last = height[stk[len(stk)-1]]
			//出栈
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			res += (i - stk[len(stk)-1] - 1) * (height[i] - last)

		}
		stk = append(stk, i)
	}
	return res
}
