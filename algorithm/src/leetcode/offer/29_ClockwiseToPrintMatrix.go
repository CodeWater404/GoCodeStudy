package offer

import "math"

/**
  @author: CodeWater
  @since: 2023/6/16
  @desc: 顺时针打印矩阵
**/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	n, m := len(matrix), len(matrix[0])
	res := make([]int, n*m)
	//(x,y)当前遍历到的坐标     d当前要调整的方向
	x, y, d := 0, 0, 0
	//dx，dy方向数组，顺序为：右下左上
	dx, dy := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	for i := 0; i < n*m; i++ {
		res[i] = matrix[x][y]
		//更新走过的位置
		matrix[x][y] = math.MaxInt32
		//a，b下一个要遍历到的位置
		a, b := x+dx[d], y+dy[d]
		//判断下一个位置的正确性，如果不在范围内那就纠正方向
		if a < 0 || a >= n || b < 0 || b >= m || matrix[a][b] == math.MaxInt32 {
			d = (d + 1) % 4
			a = x + dx[d]
			b = y + dy[d]
		}
		x, y = a, b
	}
	return res
}
