package _21_140

/**
  @author: CodeWater
  @since: 2023/10/27
  @desc:加油站
**/
// 枚举+优化
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	/*
		    i的含义：i表示当前尝试的起始加油站的位置。
		j的含义：j表示从起始加油站i开始，汽车已经经过的加油站数量。例如，当j=2时，汽车已经从加油站i行驶到加油站i+2。
	*/
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
		// 以i为起点j为终点的不能环行，下一次直接从j+1点开始进行。这里这么写是因为：结合上面ij含义，就是
		//从i点开始经过j+1个点，所以图上的“j+1”点真实的索引位置是i+j+1。
		i = i + j + 1
	}
	// 所有加油站都遍历完，不存在解
	return -1
}
