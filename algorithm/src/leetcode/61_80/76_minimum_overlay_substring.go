package _1_80

/*
*

	@author: CodeWater
	@since: 2023/12/30
	@desc: 76. 最小覆盖子串

*
*/
func minWindow(s string, t string) string {
	// hs记录窗口内各字符对应的数量；ht记录t中各字符对用的数量
	hs := make(map[byte]int)
	ht := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		ht[t[i]]++
	}

	//res答案，cnt记录hs窗口（j，i）这一段中有效能够覆盖t中字符的个数
	res, cnt := "", 0
	for i, j := 0, 0; i < len(s); i++ {
		hs[s[i]]++
		//窗口内字符数小于等于t中时，cnt才加，大于的时候加了也没用，因为跟这个cnt来判断窗口时候包含t
		if hs[s[i]] <= ht[s[i]] {
			cnt++
		}
		//j往前走的情况:只有窗口内该字符数量大于t中时，j才可以往前走
		for j < i && hs[s[j]] > ht[s[j]] { //s和t长度都为1时，会越界
			hs[s[j]]--
			j++
		}
		//找到一组解，res为0或者窗口长度比当前res小的时候才更新
		if cnt == len(t) {
			if res == "" || i-j+1 < len(res) {
				res = s[j : i+1] //(j , i-j+1)
			}
		}

	}
	return res
}
