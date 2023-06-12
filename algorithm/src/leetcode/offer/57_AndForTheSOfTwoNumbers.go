package offer

/**
  @author: CodeWater
  @since: 2023/6/12
  @desc: 和为s的两个数字
**/
func twoSum(nums []int, target int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if sum == target {
			return []int{nums[i], nums[j]}
		} else if sum < target {
			i++
		} else if sum > target {
			j--
		}

	}
	return []int{0}
}
