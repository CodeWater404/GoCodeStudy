package __20

import "strconv"

/**
  @author: CodeWater
  @since: 2024/2/27
  @desc: 9. 回文数
**/

// 1. 反转数字
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	//y保存为一个副本，res为x从个位逐渐变大的数（相当于反转x的每一位）
	y, res := x, 0
	for x > 0 {
		res = res*10 + x%10
		x /= 10
	}
	return res == y
}

// 2. 转换为字符串
func isPalindrome2(x int) bool {
	str := strconv.Itoa(x)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
