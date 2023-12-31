package _1_60

/*
*

	@author: CodeWater
	@since: 2023/12/31
	@desc: 54. 螺旋矩阵

*
*/
func spiralOrder(matrix [][]int) []int {
	res, n := make([]int, 0), len(matrix)
	if n == 0 {
		return res
	}
	m := len(matrix[0])
	//方向：左下右上
	dx, dy := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}
	//标记是否走过
	st := make([][]bool, n)
	for i := range st {
		st[i] = make([]bool, m)
	}
	for i, x, y, d := 0, 0, 0, 0; i < n*m; i++ {
		res = append(res, matrix[x][y])
		st[x][y] = true

		a, b := x+dx[d], y+dy[d] //下一个要走的位置
		//下一个位置在范围内并且走过，继续更新位置
		if a < 0 || a >= n || b < 0 || b >= m || st[a][b] {
			d = (d + 1) % 4
			a, b = x+dx[d], y+dy[d]
		}
		x, y = a, b
	}
	return res
}
