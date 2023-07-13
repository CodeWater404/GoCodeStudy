package _21_940

import "math"

/**
  @author: CodeWater
  @since: 2023/7/13
  @desc: 下降路径最小和
**/
func minFallingPathSum(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	//f[i][j]:走到f[i][j]最小的路径和
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, m)
		for j := 0; j < m; j++ {
			f[i][j] = math.MaxInt32
		}
	}
	//初始化边界
	for i := 0; i < m; i++ {
		f[0][i] = matrix[0][i]
	}
	//dp:从第二行开始遍历；（i。j）是位于数组中的坐标；k是位于ij上一行的列坐标，上一行坐标是（i-1，j-1），（i-1，j），（i-1，j+1），也就是k的范围是j-1<= k <= j + 1 ,但是同时需要满足在列坐标范围之内 0<= k <= m - 1 ,详见三层for
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			//这里就是遍历从上一行走到当前行，每个坐标的最小路径和。
			for k := max(0, j-1); k <= min(m-1, j+1); k++ {
				f[i][j] = min(f[i][j], f[i-1][k]+matrix[i][j])
			}
		}
	}
	res := math.MaxInt32
	//找出最小的和
	for i := 0; i < m; i++ {
		res = min(res, f[n-1][i])
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
