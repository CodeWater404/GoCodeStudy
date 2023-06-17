package offer

/**
  @author: CodeWater
  @since: 2023/6/17
  @desc: 二维数组中的查找
**/

func findNumberIn2DArray(matrix [][]int, target int) bool {
	//1.暴力
	// for _ , i := range matrix {
	//     for _ , j := range i {
	//         if j == target {
	//             return true
	//         }
	//     }
	// }
	// return false

	//2.优化:先从最后一行组内搜索，如果不在这一组内往上一行搜索
	i, j := len(matrix)-1, 0
	for i >= 0 && j < len(matrix[0]) {
		//先判断行减小搜索范围，在判断列
		if target < matrix[i][j] {
			i--
		} else if target > matrix[i][j] {
			j++
		} else {
			return true
		}
	}
	return false

}
