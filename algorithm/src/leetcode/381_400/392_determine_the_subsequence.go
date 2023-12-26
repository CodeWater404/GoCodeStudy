package _81_400

/*
*

	@author: CodeWater
	@since: 2023/12/26
	@desc: 392. 判断子序列

*
*/
func isSubsequence(s string, t string) bool {
	// k遍历子串
	k := 0
	for i := 0; i < len(t); i++ {
		//t中有s的字符就移动k的位置
		if k < len(s) && s[k] == t[i] {
			k++
		}
	}
	//k如果和s长度相等就说明t包含s子串
	return k == len(s)
}
