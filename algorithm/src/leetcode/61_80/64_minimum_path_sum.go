package _1_80

import "math"

/**
  @author: CodeWater
  @since: 2024/3/9
  @desc: 64. 最小路径和
**/

func minPathSum(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, m)
		for j := 0; j < m; j++ {
			f[i][j] = math.MaxInt32
		}
	}
	// f[i][j]:从左上角走到(i,j)位置处需要的最少路径数
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				f[i][j] = grid[i][j]
			} else {
				if i > 0 {
					//横向上判断，走到当前位置只有从左边走过来，加上当前的值，取最小
					f[i][j] = min(f[i][j], f[i-1][j]+grid[i][j])
				}
				if j > 0 {
					//纵向上判断，走到当前位置只有从上边走过来，加上当前的值，取最小
					f[i][j] = min(f[i][j], f[i][j-1]+grid[i][j])
				}
			}
		}
	}
	return f[n-1][m-1]
}
