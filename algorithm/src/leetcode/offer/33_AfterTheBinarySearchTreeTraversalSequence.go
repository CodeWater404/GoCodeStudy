package offer

/**
  @author: CodeWater
  @since: 2023/6/24
  @desc: 二叉搜索树的后序遍历序列
**/
func verifyPostorder(postorder []int) bool {
	//后序遍历：0时左子树；len(postorder) - 1是根节点
	return recur(postorder, 0, len(postorder)-1)
}

//i左子树起点。j根节点
func recur(postorder []int, i, j int) bool {
	//此子树节点数量 ≤1 ，无需判别正确性，因此直接返回 true
	if i >= j {
		return true
	}
	p := i
	//找到左子树的范围边界
	for postorder[p] < postorder[j] {
		p++
	}
	//m记录右子树的起点
	m := p
	//找到右子树的范围
	for postorder[p] > postorder[j] {
		p++
	}
	//p==j：p找完左右子树此时应该和j根节点相等；不等则不是
	//递归遍历左子树（i , m - 1）和右子树(m , j - 1)序列是否符合
	return p == j && recur(postorder, i, m-1) && recur(postorder, m, j-1)
}