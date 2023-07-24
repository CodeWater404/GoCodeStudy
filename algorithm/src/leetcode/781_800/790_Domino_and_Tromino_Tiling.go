package _81_800

/**
  @author: CodeWater
  @since: 2023/5/26
  @desc:
**/
func numTilings(n int) int {
	const MOD = int(1e9) + 7
	w := [][]int{
		{1, 1, 1, 1},
		{0, 0, 1, 1},
		{0, 1, 0, 1},
		{1, 0, 0, 0},
	}
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, 4)
	}
	f[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				f[i+1][k] = (f[i+1][k] + f[i][j]*w[j][k]) % MOD
			}
		}
	}
	return f[n][0]
}
