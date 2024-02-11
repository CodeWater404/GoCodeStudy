package _1_60

/**
  @author: CodeWater
  @since: 2024/2/11
  @desc: 46. 全排列
**/

var (
	res  [][]int
	path []int
	st   []bool
)

func permute(nums []int) [][]int {
	res, path, st = make([][]int, 0), make([]int, len(nums)), make([]bool, len(nums))
	if len(nums) == 0 {
		return res
	}
	dfs(nums, 0)
	return res
}

func dfs(nums []int, u int) {
	if u == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if st[i] == false {
			path[u] = nums[i]
			st[i] = true
			dfs(nums, u+1)
			st[i] = false
			//path不用管，后面的值会覆盖
		}
	}
}
