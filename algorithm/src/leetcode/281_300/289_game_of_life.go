package _81_300

/*
*

	@author: CodeWater
	@since: 2023/12/31
	@desc: 289. 生命游戏

*
*/
func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m := len(board), len(board[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			//统计（i，j）细胞周围活的细胞数
			live := 0
			//遍历细胞周围
			for x := max(0, i-1); x <= min(n-1, i+1); x++ {
				for y := max(0, j-1); y <= min(m-1, j+1); y++ {
					//去除掉当前位置(i,j)的细胞
					if (x != i || y != j) && (board[x][y]&1) == 1 {
						live++
					}
				}
			}
			//当前细胞的状态和下一步的状态
			cur, next := board[i][j]&1, 0
			if cur == 1 { //当前活着时，周围活得少于2或超过3就死亡
				if live < 2 || live > 3 {
					next = 0
				} else { //周围活得为2或者3，当前仍然存活
					next = 1
				}
			} else { //当前死亡，周围活得为3，复活
				if live == 3 {
					next = 1
				} else { //死细胞周围活的不是3个，仍然死亡
					next = 0
				}
			}
			//board数组本身只有0、1两种状态，占了个位；这里或上next左移一位后的数，就是
			//board十位上表示该细胞下一状态的情况，这样就不用多开空间
			board[i][j] |= next << 1
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			//变成下一状态，也就是把十位移到个位
			board[i][j] >>= 1
		}
	}
}
