package _01_220

/**
  @author: CodeWater
  @since: 2024/1/31
  @desc: 课程表
**/

// canFinish 存在拓扑排序就是有解，无环。入度为0加入到队列，然后不断遍历队列，加入
func canFinish(n int, edges [][]int) bool {
	// g存储图，g[a]=b表示a点可以走到b   d：存储每个点的入度，d[a]=1,a点的入度为1
	g, d := make([][]int, n), make([]int, n)
	for _, e := range edges {
		b, a := e[0], e[1]
		g[a] = append(g[a], b)
		d[b]++
	}

	q := make([]int, 0)
	for i := 0; i < n; i++ {
		if d[i] == 0 { //入度为0的点入队
			q = append(q, i)
		}
	}

	cnt := 0
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		cnt++
		for _, i := range g[t] {
			d[i]--
			if d[i] == 0 {
				q = append(q, i)
			}
		}
	}
	return cnt == n
}
