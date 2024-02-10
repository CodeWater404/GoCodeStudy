package _01_220

/**
  @author: CodeWater
  @since: 2024/2/10
  @desc: 212. 单词搜索 II
**/

type Node struct {
	id  int
	son [26]*Node
}

var (
	root   *Node                                    //trie树
	ids    map[int]bool                             // 存放找到words中单词的小标，不能使用切片存，因为可能board中存在多个解，导致ids存多次
	g      [][]byte                                 // 存放board
	dx, dy = []int{-1, 0, 1, 0}, []int{0, 1, 0, -1} //上右下左四个方向
)

func insert(word string, id int) {
	p := root
	for _, c := range word {
		u := c - 'a'
		if p.son[u] == nil {
			p.son[u] = &Node{id: -1}
		}
		p = p.son[u]
	}
	p.id = id
}

func findWords(board [][]byte, words []string) []string {
	g, root, ids = board, &Node{id: -1}, make(map[int]bool)
	for i := 0; i < len(words); i++ {
		insert(words[i], i)
	}

	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			u := g[i][j] - 'a'
			if root.son[u] != nil {
				dfs(i, j, root.son[u])
			}
		}
	}

	res := make([]string, 0)
	for k, v := range ids {
		if v == true {
			res = append(res, words[k])
		}
	}
	return res
}

func dfs(x, y int, p *Node) {
	if p.id != -1 {
		ids[p.id] = true
	}
	t := g[x][y]
	g[x][y] = '.'
	for i := 0; i < 4; i++ {
		a, b := x+dx[i], y+dy[i]
		if a >= 0 && a < len(g) && b >= 0 && b < len(g[0]) && g[a][b] != '.' {
			u := g[a][b] - 'a'
			if p.son[u] != nil {
				dfs(a, b, p.son[u])
			}
		}
	}
	g[x][y] = t
}
