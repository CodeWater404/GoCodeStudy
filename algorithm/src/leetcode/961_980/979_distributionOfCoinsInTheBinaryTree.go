package _61_980

/**
  @author: CodeWater
  @since: 2023/7/14
  @desc: 在二叉树中分配硬币
**/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	return dfs(root)[2]
}

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0, 0}
	}
	//递归求左右子树
	l, r := dfs(root.Left), dfs(root.Right)
	//x表示当前结点下的子树里面的结点数 ； y表示当前结点下的子树里面的金币数
	x, y := l[0]+r[0]+1, l[1]+r[1]+root.Val
	return []int{x, y, abs(x-y) + l[2] + r[2]}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}