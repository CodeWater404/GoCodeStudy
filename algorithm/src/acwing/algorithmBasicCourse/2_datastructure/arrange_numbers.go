package __datastructure

import "fmt"

/**
  @author: CodeWater
  @since: 2023/12/22
  @desc: 排列数字
**/

const N int = 10

var (
	n    int
	path [N]int  // 记录一组解
	st   [N]bool //  标记某个数是否被用到过
)

func dfs(u int) {
	if u == n {
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", path[i])
		}
		fmt.Println()
		return
	}

	for i := 0; i < n; i++ {
		// 该数没有用过
		if !st[i] {
			// 数字是从1开始的，所以这里加1
			path[u] = i + 1
			st[i] = true
			dfs(u + 1)
			//回溯，恢复现场
			st[i] = false
		}
	}
}

func main() {
	fmt.Scan(&n)
	dfs(0)
}
