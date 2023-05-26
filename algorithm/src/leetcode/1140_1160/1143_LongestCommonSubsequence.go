package _140_1160

/**
  @author: CodeWater
  @since: 2023/5/26
  @desc: $
**/
func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			f[i][j] = max(f[i][j-1], f[i-1][j])
			if text1[i-1] == text2[j-1] {
				f[i][j] = max(f[i][j], f[i-1][j-1]+1)
			}
		}
	}

	return f[n][m]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
