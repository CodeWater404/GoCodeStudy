package _21_140

/**
  @author: CodeWater
  @since: 2024/3/5
  @desc: 139. 单词拆分
**/

func wordBreak(s string, wordDict []string) bool {
	const P = 131 //99%的概率不会冲突
	hash := make(map[int64]bool)
	// 字符串哈希：把所有单词变成数值放入哈希表中（这一步是为了优化查询单词）
	for _, word := range wordDict {
		h := int64(0)
		for _, c := range word {
			h = h*P + int64(c)
		}
		hash[h] = true
	}

	n := len(s)
	//f[i]:字符串1-i位置所有合法的划分方案是否非空
	f := make([]bool, n+1)
	f[0] = true //f[0]：一个字符串都没有的情况，划分合法
	s = " " + s
	for i := 0; i < n; i++ {
		if f[i] { // 合法的方案
			h := int64(0)
			for j := i + 1; j <= n; j++ {
				h = h*P + int64(s[j])
				if hash[h] { // 存在p进制表示的单词，说明当前划分单词方案合法（也就是从i-j这个位置的单词）
					f[j] = true
				}
			}
		}
	}
	return f[n]
}
