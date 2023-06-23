package offer

import "sort"

/**
  @author: CodeWater
  @since: 2023/6/23
  @desc: 字符串的排列
**/
//题目并没有说明输入字符不会重复，所以有可能得到的答案序列中有重复的：但是题目要求返回的答案里面不能有重复的！！！！这是一个坑
//1.解决该问题的一种较为直观的思路是，我们首先生成所有的排列，然后进行去重。
//2.而另一种思路是我们通过修改递归函数，使得该递归函数只会生成不重复的序列。解法使用这个：具体地，我们只要在递归函数中设定一个规则，保证在填每一个空位的时候重复字符只会被填入一次。具体地，我们首先对原字符串排序，保证相同的字符都相邻，在递归函数中，我们限制每次填入的字符一定是这个字符所在重复字符集合中「从左往右第一个未被填入的字符」，

func permutation(s string) (ans []string) {
	t := []byte(s)
	//这里进行排序为了保证只使用重复字符的第一个
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	n := len(t)
	//临时记录一组解
	perm := make([]byte, 0, n)
	//标记是否访问过
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(i int) {
		//遍历到一组解
		if i == n {
			ans = append(ans, string(perm))
			return
		}
		//j：index   b:bool
		for j, b := range vis {
			//vis[j] || (j > 0 && !vis[j - 1] && s[j - 1] == s[j])：当前字符已经使用过；或者当前字符没有使用但是它是一个重复字符，上一个重复的字符已经使用过！！！
			if b || j > 0 && !vis[j-1] && t[j-1] == t[j] {
				continue
			}
			vis[j] = true
			perm = append(perm, t[j])
			backtrack(i + 1)
			perm = perm[:len(perm)-1]
			vis[j] = false
		}
	}
	backtrack(0)
	return
}
