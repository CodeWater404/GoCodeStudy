package offer

/**
  @author: CodeWater
  @since: 2023/6/22
  @desc: 二叉搜索树的最近公共祖先
**/
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		//说明pq在root的左边
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			//说明pg在root的右边
			root = root.Right
		} else {
			//找到答案（也就是pq分布在root的两边）
			break
		}

	}
	return root
}
