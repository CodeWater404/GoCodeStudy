package offer

/**
  @author: CodeWater
  @since: 2023/6/27
  @desc: 最长不含重复字符的子字符串
**/
func lengthOfLongestSubstring(s string) int {
	//哈希表（记为 dic ）统计 各字符最后一次出现的索引位置 。
	dic := make(map[byte]int)
	res, tmp := 0, 0
	for j := 0; j < len(s); j++ {
		//i表示的是和j相同字符的索引位置
		i, exists := dic[s[j]]
		if !exists {
			//没有，说明该字符第一次出现
			i = -1
		}
		//更新索引j处的value
		dic[s[j]] = j
		//tmp记录到当前字符串的最长长度
		if tmp < j-i {
			//和j处相同的上一个字符i不在当前（i，j）范围之中
			tmp += 1
		} else {
			//和j处相同的上一个字符i在当前（i，j）范围之中,当前长度等于上一次i处字符最长的长度
			tmp = j - i
		}
		// max(dp[j - 1], dp[j])
		res = max(res, tmp)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
