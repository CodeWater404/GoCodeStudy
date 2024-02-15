package _1_80

/**
  @author: CodeWater
  @since: 2024/2/15
  @desc: 74. 搜索二维矩阵
**/

// 二维的可以按照按行转化成一维的，然后用二分处理。其中涉及二位坐标的转换（mid/m , mid%m），m是列数
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	l, r := 0, n*m-1
	for l < r {
		mid := (l + r) >> 1
		if matrix[mid/m][mid%m] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return matrix[r/m][r%m] == target
}
