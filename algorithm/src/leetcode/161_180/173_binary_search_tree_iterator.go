package _61_180

/**
  @author: CodeWater
  @since: 2024/1/24
  @desc: 173. 二叉搜索树迭代器
**/

/**
 * Definition for a binary tree node.
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stk []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	bst := BSTIterator{
		stk: make([]*TreeNode, 0),
	}
	for root != nil {
		bst.stk = append(bst.stk, root)
		root = root.Left
	}
	return bst
}

func (t *BSTIterator) Next() int {
	root := t.stk[len(t.stk)-1]
	t.stk = t.stk[:len(t.stk)-1]
	val := root.Val
	root = root.Right
	for root != nil {
		t.stk = append(t.stk, root)
		root = root.Left
	}
	return val
}

func (t *BSTIterator) HasNext() bool {
	return len(t.stk) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
