package _61_280

import "sort"

/**
  @author: CodeWater
  @since: 2023/10/25
  @desc: h指数
**/
/*
这题需要理解h指数：就是一个人发表了n篇文章，他至少有j篇文章被引用j次，那h指数就是j（h指数一般取最大）。比如本题案例，他发5篇，但其实并不是每一篇都被引用5次，所以h指数不是5；4同理；而3就正好，有3篇文章引用次数超过3次，所以h指数是3.
*/
func hIndex(citations []int) int {
	// sort.Ints(citations) 默认递增
	// 降序
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	// h指数最多从数组长度开始，倒序遍历，如果一个元素值大于等于当前h指数那就说明找到一个最大的h指数（因为数组降序排列了又从尾部最小开始遍历，如果当前元素大于h，那之前的元素肯定都是大于h的，所以也就找到解）
	for h := len(citations); h > 0; h-- {
		if citations[h-1] >= h {
			return h
		}
	}
	// 遍历完就表示没有
	return 0
}
