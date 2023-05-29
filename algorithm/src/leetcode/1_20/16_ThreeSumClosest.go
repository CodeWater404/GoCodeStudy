package __20

import (
	"math"
	"sort"
)

/**
  @author: CodeWater
  @since: 2023/5/29
  @desc: 最接近的三数之和
**/
func threeSumClosest(nums []int, target int) int {
	// 第一个表示差值，第二个表示三数之和
	res := [2]int{math.MaxInt32, math.MaxInt32}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		for j, k := i+1, len(nums)-1; j < k; j++ {
			for j < k-1 && nums[i]+nums[j]+nums[k-1] >= target {
				k--
			}
			sum := nums[i] + nums[j] + nums[k]
			res = min(res, [2]int{diff(sum, target), sum})
			if j < k-1 {
				sum = nums[i] + nums[j] + nums[k-1]
				res = min(res, [2]int{diff(sum, target), sum})
			}
		}
	}
	return res[1]
}

func min(a, b [2]int) [2]int {
	if a[0] < b[0] {
		return a
	}
	return b
}
func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
