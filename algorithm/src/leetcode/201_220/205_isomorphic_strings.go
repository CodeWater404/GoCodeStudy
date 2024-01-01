package _01_220

/**
  @author: CodeWater
  @since: 2024/1/1
  @desc: 205. 同构字符串
**/

func isIsomorphic(s string, t string) bool {
	//注意，不同字符不能映射到同一个字符上！相同字符必须映射到相同字符上。
	st, ts := make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]
		if val, ok := st[a]; ok && val != b {
			return false
		}
		st[a] = b
		if val, ok := ts[b]; ok && val != a {
			return false
		}
		ts[b] = a
	}
	return true
}
