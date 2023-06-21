package offer

/**
  @author: CodeWater
  @since: 2023/6/21
  @desc: 二叉搜索树与双向链表
**/
/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;

    Node() {}

    Node(int _val) {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node* _left, Node* _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/

//本题不支持go，牛客有
func Convert(root *TreeNode) *TreeNode {
	var dfs func(cur *TreeNode)
	var pre, head *TreeNode
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		//按照左根右的顺序来建立双向链表
		dfs(cur.Left)
		//前一个结点不空的话，前一个结点的下一个结点就是当前结点
		if pre != nil {
			pre.Right = cur
		} else {
			//前一个结点为空，那么链表头就是当前结点
			head = cur
		}
		//当前的结点上一个结点就是pre
		cur.Left = pre
		//上一个结点移动到当前结点
		pre = cur
		//左子树构建链表完成，开始遍历右边
		dfs(cur.Right)

	}
	//调用函数
	dfs(root)
	return head
}
