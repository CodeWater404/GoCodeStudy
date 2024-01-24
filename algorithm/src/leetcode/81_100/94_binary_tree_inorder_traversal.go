package _1_100

/**
  @author: CodeWater
  @since: 2024/1/24
  @desc: 94. 二叉树的中序遍历
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var ans []int

// inorderTraversal recursive method
func inorderTraversal(root *TreeNode) []int {
	ans = make([]int, 0)
	dfs(root)
	return ans
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	ans = append(ans, root.Val)
	dfs(root.Right)
}

// inorderTraversal iterative method
func inorderTraversal2(root *TreeNode) []int {
	res, stk := make([]int, 0), make([]*TreeNode, 0)
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
