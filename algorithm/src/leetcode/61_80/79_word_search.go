package _1_80

/**
  @author: CodeWater
  @since: 2024/2/11
  @desc: 单词搜索
**/

var (
	dx, dy = []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
)

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(i, j, 0, word, board) {
				return true
			}
		}
	}
	return false
}

func dfs(x, y, u int, word string, board [][]byte) bool {
	if board[x][y] != word[u] {
		return false
	}
	if u == len(word)-1 {
		return true
	}
	t := board[x][y]
	board[x][y] = '.'
	for i := 0; i < 4; i++ {
		a, b := x+dx[i], y+dy[i]
		if a >= 0 && a < len(board) && b >= 0 && b < len(board[0]) && board[a][b] != '.' {
			if dfs(a, b, u+1, word, board) {
				return true
			}
		}
		//两种写法
		// if a < 0 || a >= len(board) || b < 0 || b >= len(board[0]) || board[a][b] == '.' {
		//     continue
		// }
		// if dfs(a , b , u + 1 , word , board)  {
		//     return true
		// }
	}
	board[x][y] = t
	return false
}
