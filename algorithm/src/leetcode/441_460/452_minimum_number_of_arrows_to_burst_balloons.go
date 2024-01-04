package _41_460

import "sort"

/**
  @author: CodeWater
  @since: 2024/1/4
  @desc: 452. 用最少数量的箭引爆气球
**/

func findMinArrowShots(a [][]int) int {
	res := 1
	if len(a) == 1 {
		return res
	}
	//把所有区间的右端点排序，然后看当前区间的右端点是否在下一个区间内，
	//是的话不用更新r（可以用一个箭射穿这两个重复的），不是更新r
	sort.Slice(a, func(i, j int) bool {
		return a[i][0] < a[j][0]
	})
	r := a[0][1]
	for i := 1; i < len(a); i++ {
		if r < a[i][0] {
			res++
			r = a[i][1]
		}
	}
	return res
}
