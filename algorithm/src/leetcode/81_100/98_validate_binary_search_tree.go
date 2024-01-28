package _1_100

/**
  @author: CodeWater
  @since: 2024/1/28
  @desc: 验证二叉搜索树
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// isValidBST 用定义的方法去做
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root)[0] == 1
}

// dfs 返回一个切片包含三个值，分别代表：当前子树有没有问题（0有问题，1没有问题），当前子树的最小值，当前子树的最大值
func dfs(root *TreeNode) []int {
	res := []int{1, root.Val, root.Val} //初始情况：没有问题， 最小值和最大值就是自己
	// 判断左子树
	if root.Left != nil {
		t := dfs(root.Left)
		//当前子树有问题或者子树最大值比当前节点值还大
		if t[0] == 0 || t[2] >= root.Val {
			res[0] = 0
		}
		//更新当前子树的最小和最大值
		res[1] = min(res[1], t[1])
		res[2] = max(res[2], t[2])
	}
	// 右子树同理
	if root.Right != nil {
		t := dfs(root.Right)
		if t[0] == 0 || t[1] <= root.Val {
			res[0] = 0
		}
		res[1] = min(res[1], t[1])
		res[2] = max(res[2], t[2])
	}
	return res
}
