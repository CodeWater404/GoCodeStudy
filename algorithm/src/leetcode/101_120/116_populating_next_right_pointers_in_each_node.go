package _01_120

/**
  @author: CodeWater
  @since: 2024/1/22
  @desc: 116. 填充每个节点的下一个右侧节点指针
**/

/**
 * Definition for a Node.
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 本质是宽搜，由于下一个点保存在next里面，所以队列也不用开了
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	source := root
	//遍历每一层其实就是看左节点是否存在
	for root.Left != nil {
		for p := root; p != nil; p = p.Next {
			p.Left.Next = p.Right
			//判断一下next是否存在，存在继续赋值
			if p.Next != nil {
				p.Right.Next = p.Next.Left
			}
		}
		//一层遍历完成，更新root，去下一层
		root = root.Left
	}
	//返回原始根节点
	return source
}
