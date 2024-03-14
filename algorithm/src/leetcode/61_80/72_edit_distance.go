package _1_80

/**
  @author: CodeWater
  @since: 2024/3/14
  @desc: 72. 编辑距离
**/

func minDistance(a string, b string) int {
	n, m := len(a), len(b)
	a, b = " "+a, " "+b
	//f[i][j]:将a[1-i]变成b[1-j]所有按顺序操作的最小方案
	//(只需要考虑从前往后操作，不需要考虑操作的顺序，也就是先加一个字符和后加一个字符)
	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		//a字符串不空，b字符串空的情况
		f[i][0] = i
	}
	for i := 1; i <= m; i++ {
		//a字符串空，b字符串不空的情况
		f[0][i] = i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			//a删和a加的状态
			f[i][j] = min(f[i-1][j], f[i][j-1]) + 1
			t := 0
			if a[i] != b[j] {
				t = 1
			}
			//a修改的状态
			f[i][j] = min(f[i][j], f[i-1][j-1]+t)
		}
	}
	return f[n][m]
}
