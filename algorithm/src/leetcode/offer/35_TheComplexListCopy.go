package offer

/**
  @author: CodeWater
  @since: 2023/6/10
  @desc: 35. 复杂链表的复制
**/

//Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//哈希表
var cacheNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	//在哈希表中已经存在，说明已经拷贝过
	if n, has := cacheNode[node]; has {
		return n
	}
	//没有拷贝过的拷贝
	newNode := &Node{Val: node.Val}
	//先标记哈希表中已经拷贝过
	cacheNode[node] = newNode
	//然后再递归处理下一个和随机节点
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode

}

func copyRandomList(head *Node) *Node {
	cacheNode = map[*Node]*Node{}
	return deepCopy(head)
}