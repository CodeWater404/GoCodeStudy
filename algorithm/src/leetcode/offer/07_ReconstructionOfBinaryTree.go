package offer

/**
  @author: CodeWater
  @since: 2023/6/23
  @desc: 重建二叉树
**/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var (
	pre []int
	dic map[int]int
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	//注意，这里需要为全局变量dic分配内存空间！！！！
	pre, dic = preorder, make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		//dic作为map：key为val值，value为在中序中的索引，方便下面在递归中通过前序的val值找到在中序中对应的索引位置从而划分树的左右范围
		dic[inorder[i]] = i
	}
	return recur(0, 0, len(inorder)-1)
}

func recur(root, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	//建立根节点
	node := &TreeNode{Val: pre[root], Left: nil, Right: nil}
	//通过前序划分左右结点范围，i是在中序中划分左右的一个位置
	i := dic[pre[root]]
	//左右子树递归
	node.Left = recur(root+1, left, i-1)
	node.Right = recur(root+i-left+1, i+1, right)
	return node
}
