package __20

import "sort"

/**
  @author: CodeWater
  @since: 2023/7/15
  @desc: 四数之和
**/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k, u := j+1, len(nums)-1; k < u; k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				for k < u-1 && nums[i]+nums[j]+nums[k]+nums[u-1] >= target {
					u--
				}
				if nums[i]+nums[j]+nums[k]+nums[u] == target {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[u]})
				}
			}
		}
	}
	return res
}
