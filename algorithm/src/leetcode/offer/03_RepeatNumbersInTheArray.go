package offer

/**
  @author: CodeWater
  @since: 2023/6/16
  @desc: 数组中重复的数字
**/

func findRepeatNumber(nums []int) int {
	//因为所有元素数值都在0到n-1范围内，所以可以用nums的长度
	hash := make([]int, len(nums))
	for _, value := range nums {
		if hash[value] > 0 {
			return value
		} else {
			hash[value] = 1
		}
	}
	return -1
}