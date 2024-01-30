package _21_140

/**
  @author: CodeWater
  @since: 2024/1/30
  @desc: 克隆图
**/

/**
 * Definition for a Node.
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

// 存所有的点 key是原点node，val是新复制的点node
var hash map[*Node]*Node

func cloneGraph(node *Node) *Node {
	hash = make(map[*Node]*Node)
	if node == nil {
		return nil
	}
	dfs(node) //复制所有的点到hash中
	//复制所有的边
	for s, d := range hash {
		//通过原点连接的点遍历，然后复制出相应的新边到新的点上
		for _, ver := range s.Neighbors {
			d.Neighbors = append(d.Neighbors, hash[ver])
		}
	}
	return hash[node]
}

func dfs(node *Node) {
	hash[node] = &Node{Val: node.Val}

	for _, ver := range node.Neighbors {
		//判重：这个点是第一次遍历到才继续遍历
		if _, ok := hash[ver]; ok == false {
			dfs(ver)
		}
	}
}
