package _1_40

/**
  @author: CodeWater
  @since: 2024/2/15
  @desc: 34. 在排序数组中查找元素的第一个和最后一个位置
**/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		// 找左端点,排除多余的右边target，所以是>=target
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[r] != target {
		return []int{-1, -1}
	}
	L := r
	l, r = 0, len(nums)-1
	for l < r {
		// l= mid的是否就需要+1
		mid := (l + r + 1) >> 1
		// 找右端点，所以<=target
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return []int{L, r}
}
