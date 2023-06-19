package offer

/**
  @author: CodeWater
  @since: 2023/6/19
  @desc: 矩阵中的路径
**/
func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	words := []byte(word)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(board, words, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word []byte, i, j, k int) bool {
	if i >= len(board) || i < 0 || j < 0 || j >= len(board[0]) || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	board[i][j] = '*'
	res := dfs(board, word, i+1, j, k+1) || dfs(board, word, i, j+1, k+1) ||
		dfs(board, word, i-1, j, k+1) || dfs(board, word, i, j-1, k+1)
	board[i][j] = word[k]
	return res
}
