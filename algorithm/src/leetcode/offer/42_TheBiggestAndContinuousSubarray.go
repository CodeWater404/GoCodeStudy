package offer

/**
  @author: CodeWater
  @since: 2023/6/26
  @desc: 连续子数组的最大和
**/
func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		//求连续子数组的最大和，这里直接在数组里面保存前i个和的最大值
		nums[i] += max(nums[i-1], 0)
		//然后res记录即可
		res = max(res, nums[i])
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
