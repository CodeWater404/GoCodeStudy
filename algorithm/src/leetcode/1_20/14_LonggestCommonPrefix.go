package __20

/**
  @author: CodeWater
  @since: 2023/10/31
  @desc: 最长公共前缀
**/
func longestCommonPrefix(strs []string) string {
	res := ""
	if len(strs) == 0 {
		return res
	}
	for i := 0; ; i++ {
		// 第一个单词不存在这样的字符
		if i >= len(strs[0]) {
			return res
		}
		// 拿出第一个单词的第i个字符
		c := strs[0][i]
		// 依次枚举每个单词是否存在这样的字符
		for _, v := range strs {
			// 1. 当前单词的长度在i之下，说明已经是找到一个最长的了 2. 当前单词的第i个字符不等于c，也是找到一个解
			if len(v) <= i || v[i] != c {
				return res
			}

		}
		// c这个字符，所有单词都有，res加上
		res += string(c)
	}
	return res
}
