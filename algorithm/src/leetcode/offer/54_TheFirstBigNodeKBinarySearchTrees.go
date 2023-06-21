package offer

/**
  @author: CodeWater
  @since: 2023/6/21
  @desc: 二叉搜索树的第k大节点
**/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res, k int

func kthLargest(root *TreeNode, kk int) int {
	k = kk
	dfs(root)
	return res
}

//二叉搜索树的中序遍历是一个有序的序列。这里是找出第k大的数，直接中序的倒序：右根左，这样k减小到0即是答案
func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Right)
	if k == 0 {
		return
	}
	k--
	if k == 0 {
		res = root.Val
	}
	dfs(root.Left)
}
