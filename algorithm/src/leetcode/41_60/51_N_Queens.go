package _1_60

/**
  @author: CodeWater
  @since: 2024/2/12
  @desc: 51. N 皇后
**/

var (
	n            int
	col, dg, udg []bool // col列，dg对角线，udg反对角线上是否存在皇后
	path         []string
	ans          [][]string
)

func solveNQueens(_n int) [][]string {
	n = _n
	col, dg, udg = make([]bool, n), make([]bool, n*2), make([]bool, n*2)
	ans, path = make([][]string, 0), make([]string, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			path[i] += "."
		}
	}
	dfs(0)
	return ans
}

func dfs(u int) {
	if u == n {
		temp := make([]string, len(path))
		copy(temp, path)
		ans = append(ans, temp)
		return
	}
	// 遍历每一行上的每一个位置，i就相当于x，u就相当于y ，n相当于截距b， 然后从第一行往下遍历，所以
	// 对角线dg计算就是u=i-n,n=u-i(防止出现符数，两边同时加上n)，2n=u - i + n。反对角线计算同理
	for i := 0; i < n; i++ {
		if !col[i] && !dg[u-i+n] && !udg[u+i] {
			col[i], dg[u-i+n], udg[u+i] = true, true, true
			path[u] = path[u][:i] + "Q" + path[u][i+1:]
			dfs(u + 1)
			path[u] = path[u][:i] + "." + path[u][i+1:]
			col[i], dg[u-i+n], udg[u+i] = false, false, false
		}
	}

}
