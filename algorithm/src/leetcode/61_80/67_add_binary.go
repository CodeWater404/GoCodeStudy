package _1_80

import "fmt"

/**
  @author: CodeWater
  @since: 2024/2/22
  @desc: 67. 二进制求和
**/

func addBinary(a string, b string) string {
	a, b = reverseString(a), reverseString(b)
	c := ""
	for i, t := 0, 0; i < len(a) || i < len(b) || t > 0; i++ {
		if i < len(a) {
			t += int(a[i] - '0')
		}
		if i < len(b) {
			t += int(b[i] - '0')
		}
		c += fmt.Sprintf("%d", t%2)
		t /= 2
	}
	return reverseString(c)
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
