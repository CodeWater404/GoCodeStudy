package _1_60

/**
  @author: CodeWater
  @since: 2023/10/30
  @desc:最后一个单词长度
**/
func lengthOfLastWord(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			continue
		}
		j := i - 1
		for j >= 0 && s[j] != ' ' {
			j--
		}
		return i - j
	}
	return 0
}
