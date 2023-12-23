package _1_40

/*
*

	@author: CoddeWater
	@since: 2023/11/3
	@desc: 找出字符串中第一个匹配项的下标

*
*/
func strStr(s string, p string) int {
	// kmp从1开始比较好处理
	n, m := len(s), len(p)
	s, p = " "+s, " "+p
	next := make([]int, m+1)

	// 求next数组，i在第二个字符的位置，j位于初始指向空，next[1] = 0 所以i从2开始
	for i, j := 2, 0; i <= m; i++ {
		// 注意i位置的字符是和j+1位置处的比较
		for j > 0 && p[i] != p[j+1] {
			j = next[j]
		}
		if p[i] == p[j+1] {
			j++
		}
		// 以i为结尾的字符串的最长前后缀是j
		next[i] = j
	}

	//匹配的过程，i匹配的是j+1位置，所以i从1开始
	for i, j := 1, 0; i <= n; i++ {
		for j > 0 && s[i] != p[j+1] {
			j = next[j]
		}
		if s[i] == p[j+1] {
			j++
		}
		if j == m {
			//因为下标错了一位，所以+1省略
			return i - m
		}
	}
	return -1
}
