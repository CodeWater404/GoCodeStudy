package offer

import "math"

/**
  @author: CodeWater
  @since: 2023/6/24
  @desc: 打印从1到最大的n位数
**/
func printNumbers(n int) []int {
	var res []int
	max := int(math.Pow10(n)) - 1

	for i := 1; i <= max; i++ {
		res = append(res, i)
	}
	return res
}