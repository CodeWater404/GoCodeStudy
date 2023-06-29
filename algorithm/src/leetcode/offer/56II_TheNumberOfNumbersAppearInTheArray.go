package offer

/**
  @author: CodeWater
  @since: 2023/6/29
  @desc: 数组中数字出现的次数 II
**/
func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = ones ^ num & ^twos
		twos = twos ^ num & ^ones
	}

	return ones
}