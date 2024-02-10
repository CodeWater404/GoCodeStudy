package _01_220

/**
  @author: CodeWater
  @since: 2024/2/10
  @desc: 211. 添加与搜索单词 - 数据结构设计
**/

type WordDictionary struct {
	is_end bool
	son    [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	p := this
	for _, c := range word {
		u := c - 'a'
		if p.son[u] == nil {
			p.son[u] = &WordDictionary{}
		}
		p = p.son[u]
	}
	p.is_end = true
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(this, word, 0)
}

func dfs(p *WordDictionary, word string, i int) bool {
	if i == len(word) {
		return p.is_end
	}
	if word[i] != '.' {
		u := word[i] - 'a'
		if p.son[u] == nil {
			return false
		}
		return dfs(p.son[u], word, i+1)
	} else { // 遇到万能字符‘.’
		for j := 0; j < 26; j++ {
			// p.son[j] !=nil随便匹配一个不为空的节点
			if p.son[j] != nil && dfs(p.son[j], word, i+1) {
				return true
			}
		}
		return false
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
