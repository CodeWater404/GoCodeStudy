package main

import (
	"fmt"
)

/*
  @Author: CodeWater
  @since: 2025/03/11
  @desc: 900. 整数划分
	按照完全背包思路来划分状态,f[i][j]前1到 i- 1个整数和是j的方案有多少个；优化成一维f[j]整数和是j。
的有多少个。
*/

const N, mod = 1010, 1e9 + 7

var (
	n int
	// f[j]:是整数j的划分数量
	f [N]int
)

func main() {
	fmt.Scan(&n)

	f[0] = 1
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			f[j] = (f[j] + f[j-i]) % mod
		}
	}

	fmt.Println(f[n])
}
