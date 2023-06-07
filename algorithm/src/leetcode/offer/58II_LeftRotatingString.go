package offer

import "strings"

/**
  @author: CodeWater
  @since: 2023/6/7
  @desc: 左旋转字符串
**/
func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	//三次反转
	reverse(b, 0, len(b)-1)
	reverse(b, 0, len(b)-n-1)
	reverse(b, len(b)-n, len(b)-1)
	//byte直接转string
	return string(b)
}

func reverse(s []byte, left, right int) {
	for left <= right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func reverseLeftWords2(s string, n int) string {
	res := []string{}
	for i := n; i < n+len(s); i++ {
		res = append(res, string(s[i%len(s)]))
	}
	return strings.Join(res, "")
}

func reverseLeftWords3(s string, n int) string {
	return s[n:] + s[:n]
}
