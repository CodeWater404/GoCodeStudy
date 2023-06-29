package offer

/**
  @author: CodeWater
  @since: 2023/6/29
  @desc: 构建乘积数组
**/
func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return []int{}
	}
	//把b数组每个元素的等式一行行写下来，然后划分上下三角
	b := make([]int, n)
	b[0] = 1
	temp := 1
	//计算下三角
	for i := 1; i < n; i++ {
		b[i] = b[i-1] * a[i-1]
	}
	//计算上三角
	for i := n - 2; i >= 0; i-- {
		temp *= a[i+1]
		//乘上下三角的值就是该元素
		b[i] *= temp
	}
	return b
}
