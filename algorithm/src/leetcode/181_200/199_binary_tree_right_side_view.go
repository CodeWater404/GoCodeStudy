package _81_200

/**
  @author: CodeWater
  @since: 2024/1/26
  @desc: 199. 二叉树的右视图
**/

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rightSideView 宽搜
func rightSideView(root *TreeNode) []int {
	q, res := make([]*TreeNode, 0), make([]int, 0)
	if root == nil {
		return res
	}
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			t := q[0]
			q = q[1:]
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
			//找到最右边的一个节点
			if i == n-1 {
				res = append(res, t.Val)
			}
		}
	}
	return res
}
