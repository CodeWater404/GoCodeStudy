package _41_160

/**
  @author: CodeWater
  @since: 2024/3/3
  @desc: 149. 直线上最多的点数
**/

// 枚举所有的直线
// 判断这个直线上有多少个点
// 斜率k表示一条直线，垂线斜率不存在0，中心点（0，0）是所有直线都会经过，
func maxPoints(points [][]int) int {
	res := 0
	for _, p := range points {
		//特殊情况记录： ss过中心点（0，0）上的数量，vs垂线上点的数量
		ss, vs := 0, 0
		cnt := make(map[float64]int)
		for _, q := range points {
			// p第一个点，q第二个点，通过pq可以求出斜率
			if p[0] == q[0] && p[1] == q[1] { // 两个点一摸一样，说明斜率不存在
				ss++
			} else if p[0] == q[0] { // 两个点的x坐标相同，说明过垂线
				vs++
			} else { // 正常情况，求出斜率，并统计斜率表示的这条线上点的数量
				k := float64(q[1]-p[1]) / float64(q[0]-p[0])
				cnt[k]++
			}
		}
		c := vs
		for _, t := range cnt {
			c = max(c, t)
		}
		// ss是落到原点上的，基本上每个直线都会经过，所以每个直线都可以加一下
		res = max(res, c+ss)
	}
	return res
}
