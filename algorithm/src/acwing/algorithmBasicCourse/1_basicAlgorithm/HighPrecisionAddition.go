package __basicAlgorithm

import (
	. "fmt"
	"math/big"
)

/**
  @author: CodeWater
  @since: 2023/8/21
  @desc: 高精度加法
**/

// use math/big
func compute() {
	var a, b big.Int
	var s1, s2 string
	Scan(&s1, &s2)
	a.SetString(s1, 10)
	b.SetString(s2, 10)
	Println(a.Add(&a, &b))
}

// simulation compute
func add(a, b []int) []int {
	if len(a) < len(b) {
		return add(b, a)
	}

	var c []int
	t := 0
	for i := 0; i < len(a); i++ {
		t += a[i]
		if i < len(b) {
			t += b[i]
		}
		c = append(c, t%10)
		t /= 10
	}
	if t > 0 {
		c = append(c, t)
	}
	return c

}

func compute2() {
	var a, b string
	var A, B []int
	Scan(&a, &b)
	for i := len(a) - 1; i >= 0; i-- {
		A = append(A, int(a[i]-'0'))
	}
	for i := len(b) - 1; i >= 0; i-- {
		B = append(B, int(b[i]-'0'))
	}

	c := add(A, B)
	for i := len(c) - 1; i >= 0; i-- {
		Print(c[i])
	}
	Println()
}
