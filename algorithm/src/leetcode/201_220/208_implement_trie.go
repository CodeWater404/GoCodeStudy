package _01_220

/*
*

	@author: CodeWater
	@since: 2024/2/9
	@desc: 208. 实现 Trie (前缀树)

*
*/
type Trie struct {
	is_end bool      // 打一个标记，true表示以当前字母结尾有一个单词存在
	son    [26]*Trie // 每个节点的可能性有26种
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	p := this
	for _, c := range word {
		u := c - 'a'
		if p.son[u] == nil { // 没有该节点就创造该节点
			p.son[u] = &Trie{}
		}
		// 有该节点就往下走
		p = p.son[u]
	}
	// 遍历结束后，打一个标记
	p.is_end = true
}

func (this *Trie) Search(word string) bool {
	p := this
	for _, c := range word {
		u := c - 'a'
		if p.son[u] == nil {
			return false
		}
		p = p.son[u]
	}
	return p.is_end
}

func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for _, c := range prefix {
		u := c - 'a'
		if p.son[u] == nil {
			return false
		}
		p = p.son[u]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
