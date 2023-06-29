package offer

/**
  @author: CodeWater
  @since: 2023/5/23
  @desc: 剑指 Offer II 004. 只出现一次的数字
		https://leetcode.cn/problems/WGki4K/
**/

/**singleNumber
** @Description: hash
** @param nums
** @return int
**/
func singleNumber(nums []int) int {
	hash := make(map[int]int)

	for _, num := range nums {
		hash[num]++
	}

	for value, count := range hash {
		if count == 1 {
			return value
		} else {
			continue
		}
	}

	return -1
}

/**singleNumber2
** @Description: 位运算
** @param nums
** @return int
**/
func singleNumber2(nums []int) int {
	ones, twos := 0, 0

	for _, num := range nums {
		ones = (ones ^ num) & (^twos)

		twos = (twos ^ num) & (^ones)
	}

	return ones
}
