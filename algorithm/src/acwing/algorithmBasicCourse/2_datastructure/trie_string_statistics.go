package __datastructure

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/12/25
  @desc: trie字符串统计
**/

// 字符串总长度不超过 10^5，也就相当于trie树最多有这么多层
const N int = 100010

var (
	n, idx int    // idx当前trie树插入元素用到的下标
	cnt    [N]int // 以某个字符结尾的单词有多少个（存的时候用acsii码值减去‘a’的）
	//题目中说都是小写字母，以某个字符表示的26个子节点（一维就是最大可能的字符串长度）
	son            [N][26]int
	str            string //读取的字符串
	reader, writer = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
)

// insert trie插入
func insert(str string) {
	//根节点p=0
	p := 0
	//遍历trie树，有没有str顺序的记录，有就继续往下搜索，没有就创建对应的节点
	for i := 0; i < len(str); i++ {
		u := str[i] - 'a'
		//创建节点
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx
		}
		//p指向当前字符表示的节点
		p = son[p][u]
	}
	// 插入完成后在对应的结尾字符标记数量
	cnt[p]++
}

// query 查询某字符结尾的单词也是类似，只不过在找不到对应的字符之后就返回
func query(str string) int {
	p := 0
	for i := 0; i < len(str); i++ {
		u := str[i] - 'a'
		if son[p][u] == 0 {
			return 0
		}
		p = son[p][u]
	}
	// 如果找到就返回cnt中记录的数量
	return cnt[p]
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &n)
	// 模拟while( n-- ){}
	for ; n > 0; n-- {
		var op string
		fmt.Fscan(reader, &op, &str)
		if op == "I" {
			insert(str)
		} else {
			fmt.Fprintf(writer, "%d\n", query(str))
		}
	}
}
