package _01_120

/**
  @author: CodeWater
  @since: 2024/3/8
  @desc: 120. 三角形最小路径和
**/

// f[i][j]：表示从下面往上走，走到（i，j）这个位置的最小和
func minimumTotal(f [][]int) int {
	//从最下层往上遍历
	for i := len(f) - 2; i >= 0; i-- {
		//当前层往后遍历每一位数字
		for j := 0; j <= i; j++ {
			f[i][j] += min(f[i+1][j], f[i+1][j+1])
		}
	}
	return f[0][0]
}
