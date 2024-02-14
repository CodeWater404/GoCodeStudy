package _1_60

/**
  @author: CodeWater
  @since: 2024/2/14
  @desc: 35. 搜索插入位置
**/

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
