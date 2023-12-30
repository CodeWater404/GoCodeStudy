package __20

/*
*

	@author: CodeWater
	@since: 2023/12/30
	@desc: 3. 无重复字符的最长子串

*
*/
func lengthOfLongestSubstring(s string) int {
	heap, res := make(map[byte]int), 0
	for i, j := 0, 0; i < len(s); i++ {
		heap[s[i]]++
		// 出现重复字符，收缩左边界j
		for heap[s[i]] > 1 {
			heap[s[j]]--
			j++
		}
		//1.21的版本有max，力扣支持，acwing不支持
		res = max(res, i-j+1)

	}
	return res
}
