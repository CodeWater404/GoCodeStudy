package _1_60

import "sort"

/*
*

	@author: CodeWater
	@since: 2024/1/2
	@desc: 49. 字母异位词分组

*
*/
func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)
	for _, v := range strs {
		nstr := []byte(v)
		//把字符串按照字典序排列，这样乱序的排序之后就是一组的
		sort.Slice(nstr, func(i, j int) bool {
			return nstr[i] < nstr[j]
		})
		s := string(nstr)
		hash[s] = append(hash[s], v)
	}
	res := [][]string{}
	for _, v := range hash {
		res = append(res, v)
	}
	return res
}
