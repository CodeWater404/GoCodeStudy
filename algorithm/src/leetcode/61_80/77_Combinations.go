package _1_80

/**
  @author: CodeWater
  @since: 2024/2/11
  @desc: 77. 组合
**/

var (
	res  [][]int
	path []int
)

func combine(n int, k int) [][]int {
	res, path = make([][]int, 0), make([]int, 0)
	dfs(n, k, 1)
	return res
}

func dfs(n, k, start int) {
	if k == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		dfs(n, k-1, i+1)
		path = path[:len(path)-1]
	}
}
