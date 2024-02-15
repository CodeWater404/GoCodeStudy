package _1_40

/**
  @author: CodeWater
  @since: 2024/2/15
  @desc: 33. 搜索旋转排序数组
**/

// 二分的本质：可以找到一个满足特定性质的边界点，所以对于序列来说不一定要满足单调性
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	// 二分找出两段的边界点
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] >= nums[0] { // mid>=0，说明mid在第一个区间中，边界点则在右边，更新l
			l = mid
		} else {
			r = mid - 1
		}
	}
	if target >= nums[0] { //看target是不是在第一个区间中，是的话重置l
		l = 0
	} else { // 如果target在第二个区间中，更新lr
		l, r = r+1, len(nums)-1
	}
	// 再次对一个确定的单调区间（l ， r）二分
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if nums[r] == target {
		return r
	}
	return -1
}
