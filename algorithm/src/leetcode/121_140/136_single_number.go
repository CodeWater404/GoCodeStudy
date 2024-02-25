package _21_140

/**
  @author: CodeWater
  @since: 2024/2/25
  @desc: 136. 只出现一次的数字
**/

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		// 异或：相异为1. 每对重复出现的数字异或之后变成0，最后只剩下出现单次的那个数。
		res ^= v
	}
	return res
}

// 另外一个解法：利用map记录每个数字出现的次数，然后遍历map找到出现一次的那个数。
