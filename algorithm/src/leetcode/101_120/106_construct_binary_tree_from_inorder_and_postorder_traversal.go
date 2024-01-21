package _01_120

/**
  @author: CodeWater
  @since: 2024/1/21
  @desc: 106. 从中序与后序遍历序列构造二叉树
**/

var pos = make(map[int]int)

func buildTree(inorder []int, postorder []int) *TreeNode {
	//中序元素的下标记录下,因为要通过后序的根元素找到在中序中的位置
	for i := 0; i < len(inorder); i++ {
		pos[inorder[i]] = i
	}
	return build(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

// build：il-ir中序范围，pl-pr后序范围
func build(inorder, postorder []int, il, ir, pl, pr int) *TreeNode {
	//中序的左端位置大于中序右端，说明该节点就是叶子
	if il > ir {
		return nil
	}
	//后序的最后一个就是根
	root := &TreeNode{Val: postorder[pr]}
	// 找到在中序中的位置
	k := pos[root.Val]
	//这里递归处理，对于当前节点root的左子树和右子树，它们的中序和后序范围的确定（画图好理解）
	root.Left = build(inorder, postorder, il, k-1, pl, pl+k-1-il)
	root.Right = build(inorder, postorder, k+1, ir, pl+k-1-il+1, pr-1)
	return root
}
