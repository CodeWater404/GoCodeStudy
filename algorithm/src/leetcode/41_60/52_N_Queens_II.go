package _1_60

/**
  @author: CodeWater
  @since: 2024/2/12
  @desc: 52. N皇后 II
**/

var (
	n            int
	col, dg, udg []bool // col列，dg对角线，udg反对角线上是否存在皇后
)

func totalNQueens(_n int) int {
	n = _n
	col, dg, udg = make([]bool, n), make([]bool, n*2), make([]bool, n*2)
	return dfs(0)
}

func dfs(u int) int {
	if u == n {
		return 1
	}
	res := 0
	// 遍历每一行上的每一个位置，i就相当于x，u就相当于y ，n相当于截距b， 然后从第一行往下遍历，所以
	// 对角线dg计算就是u=i-n,n=u-i(防止出现符数，两边同时加上n)，2n=u - i + n。反对角线计算同理
	for i := 0; i < n; i++ {
		if !col[i] && !dg[u-i+n] && !udg[u+i] {
			col[i], dg[u-i+n], udg[u+i] = true, true, true
			res += dfs(u + 1)
			col[i], dg[u-i+n], udg[u+i] = false, false, false
		}
	}
	return res
}
