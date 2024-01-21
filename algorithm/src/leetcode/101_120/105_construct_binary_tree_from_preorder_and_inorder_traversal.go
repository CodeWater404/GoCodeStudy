package _01_120

/**
  @author: CodeWater
  @since: 2024/1/21
  @desc: 105. 从前序与中序遍历序列构造二叉树
**/

var pos = make(map[int]int)

func buildTree(preorder []int, inorder []int) *TreeNode {
	for i := 0; i < len(inorder); i++ {
		pos[inorder[i]] = i
	}
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func build(preorder, inorder []int, pl, pr, il, ir int) *TreeNode {
	if pl > pr {
		return nil
	}
	root := &TreeNode{Val: preorder[pl]}
	k := pos[root.Val]
	root.Left = build(preorder, inorder, pl+1, pl+1+k-1-il, il, k-1)
	root.Right = build(preorder, inorder, pl+1+k-1-il+1, pr, k+1, ir)
	return root
}
