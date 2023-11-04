package _21_140

import "strings"

/**
  @author: CodedWater
  @since: 2023/11/4
  @desc: 验证回文串
**/

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < j && !check(s[i]) {
			i++
		}
		for i < j && !check(s[j]) {
			j--
		}
		if i < j && strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}
	}
	return true
}

// 检查是否是字母和数字
func check(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}
