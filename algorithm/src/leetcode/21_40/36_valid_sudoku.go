package _1_40

/*
*

	@author: CodeWater
	@since: 2023/12/31
	@desc: 36. 有效的数独

*
*/
func isValidSudoku(board [][]byte) bool {
	st := make([]bool, 9) //标记每个数字是否出现过

	//判断行
	for i := 0; i < 9; i++ {
		reset(st)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				t := board[i][j] - '1'
				if st[t] {
					return false
				}
				st[t] = true
			}
		}
	}

	//判断列
	for i := 0; i < 9; i++ {
		reset(st)
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				t := board[j][i] - '1'
				if st[t] {
					return false
				}
				st[t] = true
			}
		}
	}

	//判断小方格
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			reset(st)
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if board[i+x][j+y] != '.' {
						t := board[i+x][j+y] - '1'
						if st[t] {
							return false
						}
						st[t] = true
					}
				}
			}
		}
	}

	return true
}

func reset(st []bool) {
	for i := range st {
		st[i] = false
	}
}
