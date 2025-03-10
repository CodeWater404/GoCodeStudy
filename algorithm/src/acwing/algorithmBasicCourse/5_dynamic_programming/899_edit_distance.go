package main

/*
  @Author: CodeWater
  @Since: 2025/03/10
  @desc: 899. 编辑距离
*/

import (
	"fmt"
)

/*
跟最短编辑距离一样，变化是多了m次问询操作。
*/

const N, M = 15, 1010

var (
	n, m int
	// 存储问询的字符串，一个M次
	strs [M]string
	// f[i][j]:a的前i个字母和b的前j个字母相等需要的最少操作次数
	f [N][N]int
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func editDistance(a, b string) int {
	lenA, lenB := len(a), len(b)

	for i := 0; i <= lenB; i++ {
		f[0][i] = i
	}
	for i := 0; i <= lenA; i++ {
		f[i][0] = i
	}

	for i := 1; i <= lenA; i++ {
		for j := 1; j <= lenB; j++ {
			f[i][j] = min(f[i-1][j]+1, f[i][j-1]+1)
			if a[i-1] == b[j-1] {
				f[i][j] = min(f[i][j], f[i-1][j-1])
			} else {
				f[i][j] = min(f[i][j], f[i-1][j-1]+1)
			}
		}
	}

	return f[lenA][lenB]
}

func main() {
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		fmt.Scan(&strs[i])
	}

	// 处理实际上的m次问询
	for ; m > 0; m-- {
		var s string
		var limit int
		fmt.Scan(&s, &limit)

		// 答案
		res := 0
		for i := 0; i < n; i++ {
			// 编辑距离在合法的limit范围内，res+1
			if editDistance(strs[i], s) <= limit {
				res++
			}
		}
		fmt.Println(res)
	}
}
