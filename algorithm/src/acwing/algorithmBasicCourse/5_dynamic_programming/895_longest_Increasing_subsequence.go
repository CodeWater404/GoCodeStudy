package main

/*
  @author: CodeWater
  @since: 2025/03/05
  @desc: 898. 895. 最长上升子序列
*/

import "fmt"

const N = 1010

var (
	n int
	// f[i]以 第i个数结尾的最长子序列
	a, f [N]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 1; i <= n; i++ {
		// 极端情况：只有以i为结尾的一个，就是i自己。
		f[i] = 1
		// 枚举i前面每一个值以j结尾的情况，所有情况的最大值
		for j := 0; j < i; j++ {
			// 前提是前面的值都必须小于i结尾的这个值，题目要求严格上升。
			if a[j] < a[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
	}

	res := 0
	// 看所有数字结尾的情况是多少，取最大值。
	for i := 1; i <= n; i++ {
		res = max(res, f[i])
	}

	fmt.Println(res)
}
