package offer

/**
  @author: CodeWater
  @since: 2023/6/17
  @desc: 旋转数组的最小数字
**/
func minArray(nums []int) int {
	//1.排序
	// sort.Ints(nums)
	// return nums[0]

	//2.二分
	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) >> 1
		if nums[mid] < nums[high] {
			high = mid
		} else if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return nums[low]
}
