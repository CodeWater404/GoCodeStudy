package __basicAlgorithm

import "fmt"

/**
  @author: CodeWater
  @since: 2023/8/23
  @desc: 高精度乘法
**/

func mul(a []int, b int) []int {
	var c []int
	t := 0
	for i := 0; i < len(a) || t > 0; i++ {
		if i < len(a) {
			t = t + a[i]*b
		}
		c = append(c, t%10)
		t /= 10
	}

	for len(c) > 1 && c[len(c)-1] == 0 {
		c = c[:len(c)-1]
	}
	return c
}

func main() {
	var (
		a string
		b int
		A []int
	)
	fmt.Scan(&a, &b)
	for i := len(a) - 1; i >= 0; i-- {
		A = append(A, int(a[i]-'0'))
	}
	c := mul(A, b)

	for i := len(c) - 1; i >= 0; i-- {
		fmt.Print(c[i])
	}
	fmt.Println()
}
