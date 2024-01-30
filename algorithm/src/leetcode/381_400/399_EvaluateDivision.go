package _81_400

/*
*

	@author: CodeWater
	@since: 2023/6/5
	@desc: 除法求值

*
*/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// vers存储所有的点， d[a][b]=val存储a点到b点的权值为val
	vers, d := make(map[string]bool), make(map[string]map[string]float64)
	for i := 0; i < len(equations); i++ {
		a, b := equations[i][0], equations[i][1]
		c := values[i]
		if _, ok := d[a]; !ok { //首次遍历到a点，分配空间（go特有）
			d[a] = make(map[string]float64)
		}
		if _, ok := d[b]; !ok {
			d[b] = make(map[string]float64)
		}
		// a到b是c，那么b到a就是1/c
		d[a][b], d[b][a] = c, 1/c
		//加入到点集
		vers[a], vers[b] = true, true
	}

	// floyd（类似于无向图）
	for k := range vers {
		for i := range vers {
			for j := range vers {
				_, ok1 := d[i][k]
				_, ok2 := d[j][k]
				//存在i到k和j到k，那么i就可以到j
				if ok1 && ok2 {
					d[i][j] = d[i][k] * d[k][j]
				}
			}
		}
	}

	res := make([]float64, 0)
	for _, q := range queries {
		a, b := q[0], q[1]
		if _, ok := d[a][b]; ok {
			res = append(res, d[a][b])
		} else {
			res = append(res, -1)
		}
	}
	return res
}
