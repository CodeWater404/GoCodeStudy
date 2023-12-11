package main

import (
	"fmt"
	"math/big"
)

/**
  @author: CodeWater
  @since: 2023/8/21
  @desc: 高精度减法
**/

// 比较两个数谁大，大的返回true。
func cmp(a, b []int) bool {
	if len(a) != len(b) {
		return len(a) > len(b)
	}
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	//走到这里，说明一样大
	return true
}

func sub(a, b []int) []int {
	var c []int
	for i, t := 0, 0; i < len(a); i++ {
		//t表示借位
		t = a[i] - t
		if i < len(b) {
			t -= b[i]
		}
		//到这里，a-b就减完了。因为a-b的结果可能是负数，所以要加10取模；如果是正数，这里其实不影响
		c = append(c, (t+10)%10)
		//这里判断a-b的时候是不是用到了上一位，如果t<0说明借位了，把负的变成1，这样下一轮的时候a就会减掉借位的
		if t < 0 {
			t = 1
		} else {
			// 如果没有借位，就把t变成0，不要影响下一轮的计算。t大于0的情况：说明a[i]本身就比b[i]大，不需要借位
			t = 0
		}
	}
	//去除前导0
	for len(c) > 1 && c[len(c)-1] == 0 {
		c = c[:len(c)-1]
	}
	return c
}

// sub2 用math/big包实现高精度减法
func sub2() {
	var (
		a, b string
		A, B big.Int
	)
	fmt.Scan(&a, &b)
	A.SetString(a, 10)
	B.SetString(b, 10)
	fmt.Println(A.Sub(&A, &B))
}

func main() {
	//var (
	//	a string
	//	b string
	//	A []int
	//	B []int
	//	c []int
	//)
	//fmt.Scan(&a, &b)
	//
	////AB从个位开始读取
	//for i := len(a) - 1; i >= 0; i-- {
	//	A = append(A, int(a[i]-'0'))
	//}
	//for i := len(b) - 1; i >= 0; i-- {
	//	B = append(B, int(b[i]-'0'))
	//}
	//
	//if cmp(A, B) {
	//	c = sub(A, B)
	//} else {
	//	//小-大
	//	c = sub(B, A)
	//	fmt.Print("-")
	//}
	////c也是从个位开始，所以从尾遍历
	//for i := len(c) - 1; i >= 0; i-- {
	//	fmt.Print(c[i])
	//}
	//fmt.Println()

	sub2()
}
