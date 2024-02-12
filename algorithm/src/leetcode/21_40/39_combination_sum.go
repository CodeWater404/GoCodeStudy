package _1_40

/**
  @author: CodeWater
  @since: 2024/2/12
  @desc: 39. 组合总和
**/

var (
	ans  [][]int
	path []int
)

func combinationSum(candidates []int, target int) [][]int {
	ans, path = make([][]int, 0), make([]int, 0)
	dfs(candidates, 0, target)
	return ans
}

func dfs(candidates []int, u, target int) {
	if 0 == target {
		temp := make([]int, len(path))
		copy(temp, path)
		ans = append(ans, temp)
		return
	}
	if u == len(candidates) {
		return
	}
	for i := 0; candidates[u]*i <= target; i++ {
		dfs(candidates, u+1, target-candidates[u]*i)
		path = append(path, candidates[u])
	}
	//回溯
	for i := 0; candidates[u]*i <= target; i++ {
		path = path[:len(path)-1]
	}
}
