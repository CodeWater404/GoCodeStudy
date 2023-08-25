package __basicAlgorithm

import "fmt"

/**
  @author: CodeWater
  @since: 2023/8/25
  @desc: 前缀和
**/

const N int = 100010

var (
	n int
	m int
	a [N]int
	s [N]int
)

func main() {

	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i <= n; i++ {
		s[i] = a[i] + s[i-1]
	}

	var l, r int
	for m > 0 {
		fmt.Scan(&l, &r)
		fmt.Println(s[r] - s[l-1])
		m--
	}
}
