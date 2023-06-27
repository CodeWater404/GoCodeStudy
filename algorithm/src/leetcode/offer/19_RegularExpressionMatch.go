package offer

/**
  @author: CodeWater
  @since: 2023/6/27
  @desc: 正则表达式匹配
**/
func isMatch(s string, p string) bool {
	m, n := len(s)+1, len(p)+1
	dp := make([][]bool, m)
	//必须要这样初始化，不然dp【0】【0】索引越界
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	//初始化首行
	for j := 2; j < n; j += 2 {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}
	//状态转移
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if p[j-1] == '*' {
				if dp[i][j-2] {
					dp[i][j] = true
				} else if dp[i-1][j] && s[i-1] == p[j-2] {
					dp[i][j] = true
				} else if dp[i-1][j] && p[j-2] == '.' {
					dp[i][j] = true
				}
			} else {
				if dp[i-1][j-1] && s[i-1] == p[j-1] {
					dp[i][j] = true
				} else if dp[i-1][j-1] && p[j-1] == '.' {
					dp[i][j] = true
				}
			}

		}

	}
	return dp[m-1][n-1]
}
