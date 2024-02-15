package _61_180

/**
  @author: CodeWater
  @since: 2024/2/15
  @desc: 162. 寻找峰值
**/

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] > nums[mid+1] {
			//如果中点大于右边的数，说明从中点开始出现下降趋势，峰值在左边，更新r
			r = mid
		} else {
			// 如果中点小于右边的数，那么峰值就还是在右边区间，l更新
			l = mid + 1
		}
	}

	return r
}
