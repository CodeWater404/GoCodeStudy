package main

/*
  @author: CodeWater
  @since: 2025/03/07
  @desc: 902. 最短编辑距离
*/

import (
	"fmt"
)

const N = 1010

var (
	n, m int
	// a，b可以用string来表示，下面main中处理输入就会少很多代码
	a, b [N]byte
	// f[i][j]:a的前i个字母和b的前j个字母相等需要的最少操作次数
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
		fmt.Scanf("%c", &a[i])
	}

	fmt.Scanf("\n")
	fmt.Scan(&m)

	for i := 1; i <= m; i++ {
		fmt.Scanf("%c", &b[i])
	}
	// 如果a，b是string，那么上面的输入可以直接一行搞定：
	// fmt.Scan(&n,&a, &m, &b)
	//但是下面for循环判断a[i]和b[j]就要改成a[i-1]和b[j-1]，不然会越界

	/**初始化边界情况
	 * 1. a的第0个字母和b的第i个字母匹配，只能对A用添加操作，b有i个就是i次,所以遍历m
	 * 2. a的第i个字母和b的第0个字母匹配，只能对A用删除操作，a有i个就是i次，所以遍历n
	 */
	for i := 0; i <= m; i++ {
		f[0][i] = i
	}
	for i := 0; i <= n; i++ {
		f[i][0] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = min(f[i-1][j]+1, f[i][j-1]+1)
			if a[i] == b[j] {
				// i、j位置处的字母一样，就不用加1
				f[i][j] = min(f[i][j], f[i-1][j-1])
			} else {
				f[i][j] = min(f[i][j], f[i-1][j-1]+1)
			}
		}
	}

	fmt.Println(f[n][m])
}
