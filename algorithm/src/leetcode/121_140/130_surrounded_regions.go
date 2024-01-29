package _21_140

/**
  @author: CodeWater
  @since: 2024/1/29
  @desc: 被围绕的区域
**/

var (
	board [][]byte
	// 上右下左
	dx, dy = []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
	n, m   int // n行m列
)

// solve 题目要求把围住的O变成X，正常思路是去遍历然后改变；这里方向遍历边界上的O，这部分不变，其余地方的O
// 全部变成X即可
func solve(_board [][]byte) {
	board = _board
	n = len(board)
	if n == 0 {
		return
	}
	m = len(board[0])
	//特判左右两个边界 ， 把O变成#
	for i := 0; i < n; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][m-1] == 'O' {
			dfs(i, m-1)
		}
	}
	//特判上下两个边界 ， 把O变成#
	for i := 0; i < m; i++ {
		if board[0][i] == 'O' {
			dfs(0, i)
		}
		if board[n-1][i] == 'O' {
			dfs(n-1, i)
		}
	}
	// 重新生成：把边界上的#换成o，其余地方的就是被围住的O，全部变成X
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}
	_board = board
}

func dfs(x, y int) {
	board[x][y] = '#'
	for i := 0; i < 4; i++ {
		a, b := x+dx[i], y+dy[i]
		if a >= 0 && a < n && b >= 0 && b < m && board[a][b] == 'O' {
			dfs(a, b)
		}
	}
}
