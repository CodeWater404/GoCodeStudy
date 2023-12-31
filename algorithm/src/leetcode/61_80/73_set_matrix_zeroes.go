package _1_80

/*
*

	@author: CodeWater
	@since: 2023/12/31
	@desc: 73. 矩阵置零

*
*/
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	n, m := len(matrix), len(matrix[0])
	r0, c0 := 1, 1           //表示0行和0列位置是否有元素为0
	for i := 0; i < m; i++ { //遍历0行每一列的元素
		if matrix[0][i] == 0 {
			r0 = 0
		}
	}
	for j := 0; j < n; j++ { //遍历0列的每一行第一个元素
		if matrix[j][0] == 0 {
			c0 = 0
		}
	}

	//用第0行和第0列位置，标记其余位置是否有元素为0
	//先遍历列方向上的元素
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[j][i] == 0 {
				matrix[0][i] = 0 //标记到第0行对应列上位置
			}
		}
	}
	//遍历行方向上的元素
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0 //如果有0，第0列对应行上置为0
			}
		}
	}
	//把第0列上标记出来的0全部置为0
	for i := 1; i < m; i++ {
		if matrix[0][i] == 0 {
			for j := 0; j < n; j++ { //j这一列全部为0
				matrix[j][i] = 0
			}
		}
	}
	//把第0行上标记出来的0全部置为0
	for i := 1; i < n; i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < m; j++ { // j这一行全部为0
				matrix[i][j] = 0
			}
		}
	}
	//第0行位置上有元素为0，这一行全部为0
	if r0 == 0 {
		for i := 0; i < m; i++ {
			matrix[0][i] = 0
		}
	}
	//第0列位置上有元素为0，这一列全部为0
	if c0 == 0 {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}

}
