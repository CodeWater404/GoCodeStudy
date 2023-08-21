package __basicAlgorithm

import "fmt"

/**
  @author: CodeWater
  @since: 2023/8/21
  @desc: 高精度减法
**/

func cmp(a, b []int) bool {
	if len(a) != len(b) {
		return len(a) > len(b)
	}
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return true
}

func sub(a, b []int) []int {
	var c []int
	for i, t := 0, 0; i < len(a); i++ {
		t = a[i] - t
		if i < len(b) {
			t -= b[i]
		}
		c = append(c, (t+10)%10)
		if t < 0 {
			t = 1
		} else {
			t = 0
		}
	}
	//去除前导0
	for len(c) > 1 && c[len(c)-1] == 0 {
		c = c[:len(c)-1]
	}
	return c
}

func main() {
	var (
		a string
		b string
		A []int
		B []int
		c []int
	)
	fmt.Scan(&a, &b)

	//AB从个位开始读取
	for i := len(a) - 1; i >= 0; i-- {
		A = append(A, int(a[i]-'0'))
	}
	for i := len(b) - 1; i >= 0; i-- {
		B = append(B, int(b[i]-'0'))
	}

	if cmp(A, B) {
		c = sub(A, B)
	} else {
		//小-大
		c = sub(B, A)
		fmt.Print("-")
	}
	//c也是从个位开始，所以从尾遍历
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Print(c[i])
	}
	fmt.Println()

}
