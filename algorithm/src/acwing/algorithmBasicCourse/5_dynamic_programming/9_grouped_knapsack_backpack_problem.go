package main

/*
  @author: CodeWater
  @since: 2025/03/04
  @desc: 9. 分组背包
*/

import "fmt"

const N = 110

var (
	// n组物品 背包容量m
	n, m int
	// v[i][j], w[i][j]: 第i组第j个物品的体积、价值
	v, w [N][N]int
	// s[i]:第i组有几个物品 f[j]: 体积为j时的最大价值
	s, f [N]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Scan(&n, &m)

	for i := 1; i <= n; i++ {
		fmt.Scan(&s[i])
		for j := 0; j < s[i]; j++ {
			fmt.Scan(&v[i][j], &w[i][j])
		}
	}

	// 枚举组
	for i := 1; i <= n; i++ {
		// 体积，直接采用01背包的优化
		for j := m; j >= 0; j-- {
			// 枚举物品数量，不超过s[i]个
			for k := 0; k <= s[i]; k++ {
				// 物品体积小于等于j的才选择
				if j >= v[i][k] {
					f[j] = max(f[j], f[j-v[i][k]]+w[i][k])
				}
			}
		}
	}

	fmt.Println(f[m])
}
