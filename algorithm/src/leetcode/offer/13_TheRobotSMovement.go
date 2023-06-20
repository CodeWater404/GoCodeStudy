package offer

/**
  @author: CodeWater
  @since: 2023/6/20
  @desc: 机器人的运动范围
**/
var (
	m, n, k int
	visited [][]bool
)

func movingCount(mm int, nn int, kk int) int {
	m, n, k = mm, nn, kk
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	return dfs2(0, 0, 0, 0)
}

func dfs2(i, j, si, sj int) int {
	//(i,j)当前遍历到的坐标   si，sj数位之和
	//遍历的位置不在范围内，或者数位之和大于k，或者该位置已经被遍历过
	if i >= m || j >= n || k < si+sj || visited[i][j] {
		return 0
	}
	//标记遍历过
	visited[i][j] = true
	siNext, sjNext := si, sj
	if (i+1)%10 == 0 {
		siNext = si - 8
	} else {
		siNext = si + 1
	}
	if (j+1)%10 == 0 {
		sjNext = sj - 8
	} else {
		sjNext = sj + 1
	}
	//机器人每次只能移动一格（即只能从 x 运动至 x±1），因此每次只需计算 x 到 x±1的数位和增量。
	//向下移动sj是不变的，只有行数在变化；向右移动同理
	return 1 + dfs2(i+1, j, siNext, sj) + dfs2(i, j+1, si, sjNext)
}
