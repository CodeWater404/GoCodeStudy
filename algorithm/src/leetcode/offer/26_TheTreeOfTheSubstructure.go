package offer

/**
  @author: CodeWater
  @since: 2023/6/20
  @desc: 树的子结构
**/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(a *TreeNode, b *TreeNode) bool {
	//1. a、b有一个为空就是false
	//2. 在递归遍历过程中有一个false就不行：
	//       b有可能是从a根节点开始匹配
	//       也有可能是从a的左半部分开始匹配
	//        也有可能是从a的右半部分开始匹配（所以三次递归，之间用或）
	return (a != nil && b != nil) && (recur3(a, b) || isSubStructure(a.Left, b) || isSubStructure(a.Right, b))
}

func recur3(a, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil || a.Val != b.Val {
		return false
	}
	return recur3(a.Left, b.Left) && recur3(a.Right, b.Right)
}
