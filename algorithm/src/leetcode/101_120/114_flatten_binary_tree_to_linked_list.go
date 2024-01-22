package _01_120

/**
  @author: CodeWater
  @since: 2024/1/22
  @desc: 114. 二叉树展开为链表
**/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	for root != nil {
		p := root.Left
		//存在左子树，把左子树的右链插入到当前节点的右边
		if p != nil {
			//一直移到左子树的最右下方节点
			for p.Right != nil {
				p = p.Right
			}
			//插入过程：左子树的最右下方的节点的右指针指向当前节点的右节点
			p.Right = root.Right
			//更新当前节点的右节点为左节点
			root.Right = root.Left
			//左节点置空
			root.Left = nil
		}
		//做下一次右移
		root = root.Right
	}

}
