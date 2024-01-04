package _1_60

/**
  @author: CodeWater
  @since: 2024/1/4
  @desc: 57. 插入区间
**/

func insert(a [][]int, b []int) [][]int {
	res := make([][]int, 0)
	//新区间把老区间分为三部分：左边没有交集的，中间有交集的，右边无交集的
	//没有交集的部分直接复制到res中，有交集的合并即可
	k := 0
	for k < len(a) && a[k][1] < b[0] { //左边无交集的：a右端点小于新区间的左端点
		res = append(res, a[k])
		k++
	}
	//k还没有超过区间范围
	if k < len(a) {
		b[0] = min(a[k][0], b[0])
		//中间有交集的处理
		for k < len(a) && a[k][0] <= b[1] { //要确保a区间的左端点小于b区间的右端点这样才能合并
			b[1] = max(b[1], a[k][1]) //更新b的右端点
			k++
		}
	}
	res = append(res, b)
	//右边无交集的
	for k < len(a) {
		res = append(res, a[k])
		k++
	}
	return res
}
