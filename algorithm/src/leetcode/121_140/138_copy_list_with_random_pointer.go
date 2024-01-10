package _21_140

/**
  @author: CodeWater
  @since: 2024/1/9
  @desc: 138. 复制带随机指针的链表
**/

/**
 * Definition for a Node.
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	//1.在每个点后面复制出来一个小弟节点
	for p := head; p != nil; p = p.Next.Next { // 因为在后面复制出来一个，所以跳两下
		q := &Node{Val: p.Val}
		q.Next = p.Next
		p.Next = q //p指向小弟q
	}

	//2.复制random指针
	for p := head; p != nil; p = p.Next.Next { //还是跳过小弟
		if p.Random != nil { //random可能指向nil
			p.Next.Random = p.Random.Next //p后面时小弟节点，所以小弟的random就是p原来的random节点后面的小弟节点
		}
	}

	//3.拆分两个链表
	dummy := &Node{Val: -1}
	cur := dummy //新链表
	for p := head; p != nil; p = p.Next {
		q := p.Next                    //小弟节点
		cur.Next = q                   //新链表后面加上小弟节点
		cur, p.Next = cur.Next, q.Next //新链表移动下一个节点；恢复原来的老链表节点next指向

	}
	return dummy.Next
}
