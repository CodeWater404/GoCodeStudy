package _1_60

import "math"

/**
  @author: CodeWater
  @since: 2024/2/14
  @desc: 53. 最大子序和
**/

func maxSubArray(nums []int) int {
	res := math.MinInt32
	for i, last := 0, 0; i < len(nums); i++ {
		//last记录上一次f(i)的最大值
		last = nums[i] + max(last, 0)
		res = max(res, last)
	}
	return res
}
