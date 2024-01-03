package _1_60

import "sort"

/**
  @author: CodeWater
  @since: 2024/1/3
  @desc: 56. 合并区间
**/

func merge(a [][]int) [][]int {
	res := make([][]int, 0)
	if len(a) == 0 {
		return res
	}
	//先对左端点排序
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	//然后比较当前右端点和下一个区间的左端点是否重合，是的话更新右端点，不是的话就是一组解
	l, r := a[0][0], a[0][1]
	for i := 1; i < len(a); i++ {
		if a[i][0] > r {
			res = append(res, []int{l, r})
			l, r = a[i][0], a[i][1]
		} else {
			r = max(r, a[i][1])
		}
	}
	res = append(res, []int{l, r})
	return res
}
