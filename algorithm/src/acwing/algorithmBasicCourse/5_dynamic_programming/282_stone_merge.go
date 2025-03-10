package main

/*
  @Author: CodeWater
  @Since: 2025/03/10
  @desc: 282. 石子合并
*/

import (
	"fmt"
)

const N = 310

var (
	n int
	s [N]int
	//f[i][j]从第i堆到第j堆石子需要合并的最小代价
	f [N][N]int
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&s[i])
	}

	// 求前缀和
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + s[i]
	}
	//按照长度从小到大来枚举 ， 边界情况区间长度为1的时候不需要代价，f为全局默认0，所以从2开始
	for length := 2; length <= n; length++ {
		// 枚举起点1， 终点要小于等于n
		for i := 1; i+length-1 <= n; i++ {
			l, r := i, i+length-1
			// l - r的初始值要赋一个无穷大，不然下面计算永远是0
			f[l][r] = int(1e8)
			// 开始划分从l - r这之间的k分界线
			for k := l; k < r; k++ {
				f[l][r] = min(f[l][r], f[l][k]+f[k+1][r]+s[r]-s[l-1])
			}
		}
	}
	fmt.Println(f[1][n])
}
