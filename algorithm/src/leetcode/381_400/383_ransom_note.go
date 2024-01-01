package _81_400

/*
*

	@author: CodeWater
	@since: 2024/1/1
	@desc: 383. 赎金信

*
*/
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if m[ransomNote[i]] == 0 {
			return false
		} else {
			m[ransomNote[i]]--
		}
	}

	return true
}
