package _1_100

/**
  @author: CodeWater
  @since: 2024/1/19
  @desc: 100. 相同的树
**/

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	//都是空树的情况
	if p == nil && q == nil {
		return true
	}
	// 其中有一棵树为空(从上面if走到这必然有棵树不为空)，或者值不一样
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
