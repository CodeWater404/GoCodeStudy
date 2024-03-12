package __20

/**
  @author: CodeWater
  @since: 2024/3/12
  @desc: 5. 最长回文子串
**/

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		//假设回文串是奇数，那么lr围绕i两边开始遍历走
		l, r := i-1, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if len(res) < r-l-1 {
			//合适的范围[l+1,r-1]
			res = s[l+1 : r]
		}
		//假设回文串是奇数，那么l从i开始往左，r从i+1开始往右
		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if len(res) < r-l-1 {
			res = s[l+1 : r]
		}
	}
	return res
}
