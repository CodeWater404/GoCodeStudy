package _41_160

import "strings"

/*
*

	@author: CodeWater
	@since: 2023/12/17
	@desc: 反转字符串中的单词

*
*/
func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	j := len(s) - 1
	i, res := j, ""
	// 双指针：ij从尾往前扫描，i先走
	for i >= 0 {
		// i遇到空格后退出
		for i >= 0 && s[i] != ' ' {
			i--
		}
		//保存一个答案
		res += s[i+1:j+1] + " "
		for i >= 0 && s[i] == ' ' {
			i--
		}
		//i走到有字符的位置后退出，j更新到这个位置，实际上相当于单词末尾
		j = i
	}
	//去掉res最后的空格
	return strings.TrimSpace(res)
}
