package offer

/**
  @author: CodeWater
  @since: 2023/6/26
  @desc: 礼物的最大价值
**/
func maxValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			//起点直接下一次
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				//为最上边的格子只能由左边的格子走到
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				//位于左边的格子只能由上方格子走下来
				grid[i][j] += grid[i-1][j]
			} else {
				//正常有左边和上边的格子价值=本身价值加上max（左边，上边）
				grid[i][j] += max(grid[i][j-1], grid[i-1][j])
			}

		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
