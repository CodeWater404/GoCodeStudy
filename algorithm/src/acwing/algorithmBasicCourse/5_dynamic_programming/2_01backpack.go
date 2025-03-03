package __dynamic_programming

/**
  @author: CodeWater
  @since: 2025/02/28
  @desc: 01背包问题
**/

import (
	"fmt"
)

const N = 1010

var (
	// n个物品 背包容量为m
	n, m int
	// v存每个物品的体积 w存每个物品的价值
	v, w [N]int
	// f[i][j] = a: 在前i个物品中所选择的总体积小于等于j的最大价值a
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
		fmt.Scan(&v[i], &w[i])
	}

	// f[0][0-m]，不选择物品时的价值是0，而f为包级变量默认数组值就是0，
	// 所以这里不用对这个进行初始化

	// 直接从1开始
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			/*本题状态有两种：
			  1. f[i - 1][j]: 不选择当前物品，即上面一种恒存在状态，无需判断
			  2. f[i][j]: 选择当前物品，前提是当前背包容量得大于当前物品重量，不然下面的回退上一个状态计算当前状态会数组越界f[i-1][j-v[i]]，所以有如下判断。正常应该是j - v[i] >= 0,这里为了方便做了个变形。
			*/
			if j >= v[i] {
				f[i][j] = max(f[i-1][j], f[i-1][j-v[i]]+w[i])
			}
		}
	}

	fmt.Println(f[n][m])
}
