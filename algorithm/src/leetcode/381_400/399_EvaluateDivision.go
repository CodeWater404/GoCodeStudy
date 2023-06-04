package _81_400

/**
  @author: CodeWater
  @since: 2023/6/5
  @desc: 除法求值
**/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	//顶点数组
	vers := make(map[string]bool)
	//两点之间的距离
	d := make(map[string]map[string]float64)

	for i := 0; i < len(equations); i++ {
		a, b := equations[i][0], equations[i][1]
		c := values[i]

		if _, ok := d[a]; !ok {
			d[a] = make(map[string]float64)
		}
		if _, ok := d[b]; !ok {
			d[b] = make(map[string]float64)
		}
		//建边上的权值
		d[a][b], d[b][a] = c, 1/c
		vers[a], vers[b] = true, true
	}

	//Floyd算法
	for k := range vers {
		for i := range vers {
			for j := range vers {
				if d[i][k] != 0 && d[k][j] != 0 {
					d[i][j] = d[i][k] * d[k][j]
				}
			}
		}
	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		a, b := q[0], q[1]
		//边上的权值存在，说明存在答案
		if val, ok := d[a][b]; ok {
			res[i] = val
		} else {
			res[i] = -1
		}
	}

	return res
}
