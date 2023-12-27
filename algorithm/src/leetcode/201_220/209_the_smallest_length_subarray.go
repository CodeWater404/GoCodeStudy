package _01_220

import "math"

/*
*

	@author: CodeWater
	@since: 2023/12/27
	@desc: 209. 长度最小的子数组

*
*/
func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt32
	//i往右扫描，j负责搜索左边界，sum表示j-i之间的总和
	for i, j, sum := 0, 0, 0; i < len(nums); i++ {
		sum += nums[i]
		//试探下一个j是否还大于等于target，是就更新
		for sum-nums[j] >= target {
			sum -= nums[j]
			j++
		}
		if sum >= target {
			res = min(res, i-j+1)
		}
	}

	if res == math.MaxInt32 {
		res = 0
	}
	return res
}
