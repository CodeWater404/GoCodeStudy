package _01_920

/**
  @author: CodeWater
  @since: 2024/2/1
  @desc: 蛇梯棋
**/

type pair struct {
	x, y int
}

var (
	id  [][]int
	cor []pair
)

func snakesAndLadders(board [][]int) int {
	n, m := len(board), len(board[0])
	// id存放矩阵中的数字，cor[k]记录跳跃到k位置位于id中的坐标(x,y)
	id, cor = make([][]int, n), make([]pair, n*m+1)
	for i := 0; i < m; i++ {
		id[i] = make([]int, m)
	}
	// k要填充的数字 ， s标志当前行从哪一边开始填
	for i, k, s := n-1, 1, 0; i >= 0; i, s = i-1, s+1 {
		if s%2 == 0 { // 偶数行是从左往右填充
			for j := 0; j < m; j, k = j+1, k+1 {
				id[i][j] = k
				cor[k] = pair{i, j}
			}
		} else {
			for j := m - 1; j >= 0; j, k = j-1, k+1 {
				id[i][j] = k
				cor[k] = pair{i, j}
			}
		}
	}
	// q队列，dist[x][y]存储从遍历的起点到(x,y)所花费的距离是多少
	q, dist := make([]pair, 0), make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dist[i][j] = 1e9
		}
	}
	q, dist[n-1][0] = append(q, pair{n - 1, 0}), 0
	for len(q) > 0 {
		t := q[0]
		q = q[1:]

		k := id[t.x][t.y]
		if k == n*m { // 判断是否走到最后一个编号
			return dist[t.x][t.y]
		}
		// 遍历下一个可能的走向
		for i := k + 1; i <= k+6 && i <= n*m; i++ {
			x, y := cor[i].x, cor[i].y
			if board[x][y] == -1 { // 下一个位置是个普通位置
				//如果下一个位置大于当前位置+1（实际上就是正无穷），那么更新下一个位置dist
				if dist[x][y] > dist[t.x][t.y]+1 {
					dist[x][y] = dist[t.x][t.y] + 1
					q = append(q, pair{x, y})
				}
			} else { // 如果是梯子或者蛇，说明可以跳跃移动
				r := board[x][y]
				//取到跳跃位置编号为r处的坐标(x,y)
				x, y = cor[r].x, cor[r].y
				//如果下一个位置大于当前位置+1（实际上就是正无穷），那么更新下一个位置dist
				if dist[x][y] > dist[t.x][t.y]+1 {
					dist[x][y] = dist[t.x][t.y] + 1
					q = append(q, pair{x, y})
				}
			}
		}
	}
	return -1
}
