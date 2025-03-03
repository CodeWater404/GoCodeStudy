package __dynamic_programming

/**
  @author: CodeWater
  @since: 2025/03/03
  @desc: 完全背包问题
  和01背包非常类似，但是物品有无限个，所以集合划分按照每个物品选择了多少个来；由于物品是有重量，背包有容量，
  所以不能无限选择，假设选择到k个。
  这样就把f[i][j]分成了k类， 第一个部分是第i个物品选了0个的选法、 第二个部分是第i个物品选了1个的选法、 第三个部分是第i个物品选了2个的选法...。
  然后看每个部分该如何算。
  - 0个： 当前不选，等于上一层选择的情况。f[i - 1][j]
  - 1个往后的状态： 跟01背包分析一样，当前状态由之前的状态的推算而来，所以可以先减去k个物品然后加回来k个物品。 f[i - 1][j - k * v[i]] + k * w[i]
  当k = 0时，其实就是0个的表达式。所以这两个状态的表达式可以合并为第二个表达式。
**/

/* =====================朴素===================== */
// import (
// 	"fmt"
// )

// const N = 1010

// var (
// 	// n中物品，背包容量为m
// 	n, m int
// 	v, w [N]int
// 	f    [N][N]int
// )

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	fmt.Scan(&n, &m)
// 	for i := 1; i <= n; i++ {
// 		fmt.Scan(&v[i], &w[i])
// 	}

// 	for i := 1; i <= n; i++ {
// 		for j := 0; j <= m; j++ {
// 			for k := 0; k*v[i] <= j; k++ {
// 				f[i][j] = max(f[i][j], f[i-1][j-k*v[i]]+k*w[i])
// 			}
// 		}
// 	}

// 	fmt.Println(f[n][m])
// }

/* =====================优化一层for===================== */
// import (
// 	"fmt"
// )

// const N = 1010

// var (
// 	// n中物品，背包容量为m
// 	n, m int
// 	v, w [N]int
// 	f    [N][N]int
// )

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	fmt.Scan(&n, &m)
// 	for i := 1; i <= n; i++ {
// 		fmt.Scan(&v[i], &w[i])
// 	}

// 	for i := 1; i <= n; i++ {
// 		for j := 0; j <= m; j++ {
// 			f[i][j] = f[i-1][j]
// 			if j >= v[i] {
// 				// 状态方程变形一下，可以去掉k，优化一层for
// 				f[i][j] = max(f[i][j], f[i][j-v[i]]+w[i])
// 			}
// 		}
// 	}

// 	fmt.Println(f[n][m])
// }

/* =====================状态方程优化成一维===================== */
import (
	"fmt"
)

const N = 1010

var (
	// n中物品，背包容量为m
	n, m int
	v, w [N]int
	f    [N]int
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

	for i := 1; i <= n; i++ {
		// 继续优化，状态方程变成一维
		for j := v[i]; j <= m; j++ {
			f[j] = max(f[j], f[j-v[i]]+w[i])
		}
	}

	fmt.Println(f[m])
}
