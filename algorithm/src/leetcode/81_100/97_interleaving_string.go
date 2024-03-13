package _1_100

/**
  @author: CodeWater
  @since: 2024/3/13
  @desc: 97. 交错字符串
**/

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m := len(s1), len(s2)
	if len(s3) != n+m {
		return false
	}
	//f[i][j]:能否从s1(1,i)和s2(1,j)中选出匹配上s3(1,i+j-1)的方案
	f := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]bool, m+1)
	}
	s1, s2, s3 = " "+s1, " "+s2, " "+s3
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i == 0 && j == 0 {
				f[i][j] = true
			} else {
				if i > 0 && s1[i] == s3[i+j] {
					f[i][j] = f[i-1][j]
				}
				if j > 0 && s2[j] == s3[i+j] {
					f[i][j] = f[i][j] || f[i][j-1]
				}
			}
		}
	}
	return f[n][m]
}
