package offer

/**
  @author: CodeWater
  @since: 2023/6/17
  @desc: 0～n-1中缺失的数字
**/
func missingNumber(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) >> 1
		//如果前半部分没有缺少值，那么说明中间值就正好是mid，缺少值只可能在右半部分
		if nums[mid] == mid {
			i = mid + 1
		} else {
			//前半部分缺少值，缩小右边界
			j = mid - 1
		}
	}
	return i
}
