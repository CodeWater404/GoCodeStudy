package _21_240

/**
  @author: CodeWater
  @since: 2024/3/31
  @desc: 221. 最大正方形
**/

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n, m := len(matrix), len(matrix[0])
	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, m+1)
	}
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if matrix[i-1][j-1] == '1' {
				f[i][j] = min(f[i-1][j], min(f[i][j-1], f[i-1][j-1])) + 1
				res = max(res, f[i][j])
			}
		}
	}
	return res * res
}
