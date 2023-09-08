package __basicAlgorithm

import "fmt"

/**
  @author: CodeWater
  @since: 2023/9/9
  @desc: 最长连续不重复子序列
**/

const N int = 100010

var (
	n    int
	q, s [N]int
)

func main() {
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}

	res := 0
	for i, j := 0, 0; i < n; i++ {
		s[q[i]]++
		for j < i && s[q[i]] > 1 {
			s[q[j]]--
			j++
		}
		res = max(res, i-j+1)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
