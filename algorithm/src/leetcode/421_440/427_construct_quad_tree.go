package _21_440

/**
  @author: CodeWater
  @since: 2024/2/13
  @desc: 427. 建立四叉树
**/

/**
 * Definition for a QuadTree node.
 */

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

var s [][]int

func construct(w [][]int) *Node {
	n := len(w)
	s = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		s[i] = make([]int, n+1)
	}
	// 求二维数组的前缀和
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			s[i][j] = s[i-1][j] + s[i][j-1] - s[i-1][j-1] + w[i-1][j-1]
		}
	}
	return dfs(1, 1, n, n)
}

func dfs(x1, y1, x2, y2 int) *Node {
	n := x2 - x1 + 1
	// 用二维数组的前缀和判断一个区间内是不是全1或者全0
	sum := s[x2][y2] - s[x2][y1-1] - s[x1-1][y2] + s[x1-1][y1-1]
	if sum == 0 || sum == n*n {
		return &Node{Val: sum != 0, IsLeaf: true} // 叶子节点
	}
	node := &Node{Val: false, IsLeaf: false}
	m := n / 2 // 每次划分一般作为界限
	node.TopLeft = dfs(x1, y1, x1+m-1, y1+m-1)
	node.TopRight = dfs(x1, y1+m, x1+m-1, y2)
	node.BottomLeft = dfs(x1+m, y1, x2, y1+m-1)
	node.BottomRight = dfs(x1+m, y1+m, x2, y2)
	return node
}
