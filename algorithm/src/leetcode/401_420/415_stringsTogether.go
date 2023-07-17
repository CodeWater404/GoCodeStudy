package _01_420

import "strconv"

/**
  @author: CodeWater
  @since: 2023/7/17
  @desc: 字符串相加
**/
func add(a, b []int) []int {
	c := []int{}
	//t表示进位
	for i, t := 0, 0; i < len(a) || i < len(b) || t > 0; i++ {
		//同位相加，以及上一次的进位
		if i < len(a) {
			t += a[i]
		}
		if i < len(b) {
			t += b[i]
		}
		c = append(c, t%10)
		t /= 10
	}
	return c
}

func addStrings(num1 string, num2 string) string {
	a, b := []int{}, []int{}
	//字符串转换为数组存储
	for i := len(num1) - 1; i >= 0; i-- {
		a = append(a, int(num1[i]-'0'))
	}
	for i := len(num2) - 1; i >= 0; i-- {
		b = append(b, int(num2[i]-'0'))
	}
	c := add(a, b)
	var res string
	for i := len(c) - 1; i >= 0; i-- {
		//数组转换为字符串
		res += strconv.Itoa(c[i])
	}
	return res
}
