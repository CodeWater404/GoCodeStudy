package _81_300

import "strings"

/*
*

	@author: CodeWater
	@since: 2024/1/1
	@desc: 290. 单词规律

*
*/
func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(pattern) != len(words) {
		return false
	}

	//pw映射p字符到对应word单词，wp映射word单词到对应的p字符
	pw, wp := make(map[byte]string), make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		a := pattern[i]
		b := words[i]

		//pw[a]存在并且val映射的值不等于当前单词b，匹配失败
		if val, ok := pw[a]; ok && val != b {
			return false
		}
		pw[a] = b
		//wp[b]存在并且val映射的值不等于当前字符a，匹配失败
		if val, ok := wp[b]; ok && val != a {
			return false
		}
		wp[b] = a
	}
	return true
}
