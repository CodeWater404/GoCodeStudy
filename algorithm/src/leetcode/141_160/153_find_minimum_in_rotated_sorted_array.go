package _41_160

/**
  @author: CodeWater
  @since: 2024/2/16
  @desc: 153. 寻找旋转排序数组中的最小值
**/

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	// 本身是个升序序列，直接返回
	if nums[r] >= nums[l] {
		return nums[0]
	}
	for l < r {
		mid := (l + r) >> 1
		// mid在下边的上升区间中，最小值在左边，更新r
		if nums[mid] < nums[0] {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return nums[l]
}
