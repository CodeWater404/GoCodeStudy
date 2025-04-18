package main

/**
  @author: CodeWater
  @since: 2025/03/03
  @desc: 多重背包问题
  和完全背包非常类似，但是物品有限个，所以集合划分按照每个物品选择了多少个来；由于物品是有重量，背包有容量，
  所以不能无限选择，假设选择到k个。
  这样就把f[i][j]分成了k类， 第一个部分是第i个物品选了0个的选法、 第二个部分是第i个物品选了1个的选法、 第三个部分是第i个物品选了2个的选法...。
  然后看每个部分该如何算。
  - 0个： 当前不选，等于上一层选择的情况。f[i - 1][j]
  - 1个往后的状态： 跟01背包分析一样，当前状态由之前的状态的推算而来，所以可以先减去k个物品然后加回来k个物品。 f[i - 1][j - k * v[i]] + k * w[i]
  当k = 0时，其实就是0个的表达式。所以这两个状态的表达式可以合并为第二个表达式。
  状态方程都是一样的，区别是多了个条件
**/

/* =====================朴素三重for===================== */

import (
	"fmt"
)

const N = 110

var (
	n, m int
	// v体积 w价值 s个数
	v, w, s [N]int
	// f[i][j]: 前i种物品体积不超过j的最大价值
	f [N][N]int
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
		fmt.Scan(&v[i], &w[i], &s[i])
	}

	// f[0][0 ~ m - 1]:选取0种物品，价值是0，用数组默认值即可，不显示初始化。所以下面i从1开始

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			// 每种物品不超过s[i]件并且当前选取物品重量不能超过当前背包的容量
			for k := 0; k <= s[i] && k*v[i] <= j; k++ {
				f[i][j] = max(f[i][j], f[i-1][j-k*v[i]]+k*w[i])
			}
		}
	}

	fmt.Println(f[n][m])
}
