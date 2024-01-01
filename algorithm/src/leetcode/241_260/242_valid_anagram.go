package _41_260

/**
  @author: CodeWater
  @since: 2024/1/1
  @desc: 242. 有效的字母异位词
**/

func isAnagram(s string, t string) bool {
	//两个字符串必须长度相等，只是字符位置不一样
	hash := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		hash[t[i]]--
	}
	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true

}
