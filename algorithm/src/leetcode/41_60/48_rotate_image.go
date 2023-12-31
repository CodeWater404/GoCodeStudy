package _1_60

/**
  @author: CodeWater
  @since: 2023/12/31
  @desc: 48. 旋转图像
**/
//先按照对角线反转，然后按照中轴反转即可得到顺时针九十度的反转
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		//j中轴左边，k中轴右边
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}
