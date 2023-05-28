package __20

import "sort"

/**
  @author: CodeWater
  @since: 2023/5/28
  @desc: 三数之和
**/

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 双指针
		for j, k := i+1, len(nums)-1; j < k; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k-1 && nums[i]+nums[j]+nums[k-1] >= 0 {
				k--
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}

	}

	return res
}
