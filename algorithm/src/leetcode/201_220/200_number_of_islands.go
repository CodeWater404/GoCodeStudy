package _01_220

/**
  @author: CodeWater
  @since: 2024/1/29
  @desc: 岛屿数量
**/

var (
	g [][]byte
	//上右下左
	dx, dy = []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
)

// numIslands 遍历每一个可以走的点，走出dfs之后就是一个块，然后继续下一个可以走的点（上一步已经改了走过的点，
// 所以不用担心重复走过的点）
func numIslands(grid [][]byte) int {
	g = grid
	cnt := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			if g[i][j] == '1' {
				dfs(i, j)
				cnt++
			}
		}
	}
	return cnt
}

func dfs(x, y int) {
	g[x][y] = '0'
	//遍历周围四个可以走的方向
	for i := 0; i < 4; i++ {
		a, b := x+dx[i], y+dy[i]
		if a >= 0 && a < len(g) && b >= 0 && b < len(g[a]) && g[a][b] == '1' {
			dfs(a, b)
		}
	}
}
