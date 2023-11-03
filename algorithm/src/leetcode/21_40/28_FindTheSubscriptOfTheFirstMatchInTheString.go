package _1_40

/**
  @author: CoddeWater
  @since: 2023/11/3
  @desc: 找出字符串中第一个匹配项的下标
**/
func strStr(s string, p string) int {
	if len(p) == 0 {
		return 0
	}
	n, m := len(s), len(p)
	s, p = " "+s, " "+p

	next := make([]int, m+1)
	for i, j := 2, 0; i <= m; i++ {
		for j > 0 && p[i] != p[j+1] {
			j = next[j]
		}
		if p[i] == p[j+1] {
			j++
		}
		next[i] = j
	}

	for i, j := 1, 0; i <= n; i++ {
		for j > 0 && s[i] != p[j+1] {
			j = next[j]
		}
		if s[i] == p[j+1] {
			j++
		}
		if j == m {
			return i - m
		}
	}
	return -1
}
