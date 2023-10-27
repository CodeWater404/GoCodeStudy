package _21_140

/**
  @author: CodeWater
  @since: 2023/10/27
  @desc:加油站
**/
// 枚举+优化
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	// i枚举每一个加油站
	for i, j := 0, 0; i < n; {
		// 所剩的油量
		left := 0
		for j = 0; j < n; j++ {
			// 每一轮起点
			k := (i + j) % n
			// 加上当前加油站的油量，减掉去下一个加油站需要的油量
			left += gas[k] - cost[k]
			if left < 0 {
				// 无法到达下一站，直接退出本轮以i为起点的加油站
				break
			}
		}
		// j完整的遍历完所有的加油站，说明刚刚以i为起点的加油站可以环行
		if j == n {
			return i
		}
		// 以i为起点j为终点的不能环行，下一次直接从j+1点开始进行
		i = i + j + 1
	}
	// 所有加油站都遍历完，不存在解
	return -1
}
