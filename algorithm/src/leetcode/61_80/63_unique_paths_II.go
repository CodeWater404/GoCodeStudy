package _1_80

/**
  @author: CodeWater
  @since: 2024/3/11
  @desc: 63. 不同路径 II
**/

func uniquePathsWithObstacles(o [][]int) int {
	n := len(o)
	if n == 0 {
		return 0
	}
	m := len(o[0])
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if o[i][j] == 0 { //没有障碍物
				if i == 0 && j == 0 { //起点
					f[i][j] = 1
				} else {
					if i > 0 { //当前位置只能由上边走下来
						f[i][j] += f[i-1][j]
					}
					if j > 0 { //当前位置只能由左边走过来
						f[i][j] += f[i][j-1]
					}
				}
			}
		}
	}
	return f[n-1][m-1]
}
