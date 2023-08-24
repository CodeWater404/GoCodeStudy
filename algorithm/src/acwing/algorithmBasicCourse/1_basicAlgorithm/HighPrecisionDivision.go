package __basicAlgorithm

import "fmt"

/**
  @author: CodeWater
  @since: 2023/8/24
  @desc: 高精度除法
**/

func div(a []int, b int, r *int) []int {
	var c []int
	*r = 0
	for i := len(a) - 1; i >= 0; i-- {
		*r = *r*10 + a[i]
		c = append(c, *r/b)
		*r %= b
	}
	var temp []int
	for i := len(c) - 1; i >= 0; i-- {
		temp = append(temp, c[i])
	}
	//反转c
	c = temp
	//去除前导0
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
		r int
	)
	fmt.Scan(&a, &b)
	for i := len(a) - 1; i >= 0; i-- {
		A = append(A, int(a[i]-'0'))
	}

	c := div(A, b, &r)

	for i := len(c) - 1; i >= 0; i-- {
		fmt.Print(c[i])
	}
	fmt.Printf("\n%d", r)

}
