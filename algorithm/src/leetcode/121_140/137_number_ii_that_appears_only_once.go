package _21_140

/**
  @author: CodeWater
  @since: 2024/2/26
  @desc: 137. 只出现一次的数字 II
**/

func singleNumberII(nums []int) int {
	one, two := 0, 0
	for _, x := range nums {
		one = (one ^ x) & (^two)
		two = (two ^ x) & (^one)
	}
	// 哪些位为1，就是个数出现一次的数，也就是one
	return one
}
