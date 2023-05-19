package main

import (
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/5/14
  @desc: bfs广搜，走迷宫
	6行5列的迷宫，0可以通过，1不能通过。起点左上方，终点右下方.示例在maze.in文件中
**/

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	//这里初始化的是row行，也就是第一个【】 ， 类型是【】int的切片（注意不是数组）
	maze := make([][]int, row)
	for i := range maze {
		//这里初始化的是col列，也就是第二个【】 ， 类型是int的切片（注意不是数组）
		maze[i] = make([]int, col)

		for j := range maze[i] {
			//给二维数组中的每一个元素赋值，（从file中读取）
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

//定义要走的方向:上左下右
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

/**at
** @Description: 判断时候在地图中越界
** @receiver p
** @param grid
** @return int: 返回路径数组中记录的值
** @return bool： false越界， true没有
**/
func (p point) at(grid [][]int) (int, bool) {
	//行范围
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	//列范围
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	//steps记录走过的路径
	steps := make([][]int, len(maze))
	//初始化
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	//要遍历的队列,start初始节点入队
	Q := []point{start}
	//遍历队列。（两种情况结束：1.队列为空  2.遍历到）
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		//走到终点
		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			//可以走的情况：
			//maze at next is 0 , and steps at next is 0 , and next != start
			//下面排除的是不可以走的情况
			val, ok := next.at(maze)
			if !ok || val == 1 { //1表示遇到墙，不能走
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}
			curSteps, _ := cur.at(steps)
			//符和的情况，路径加1
			steps[next.i][next.j] = curSteps + 1
			//更新队列，点加入到队列中
			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	//maze := readMaze("maze/maze.in")
	maze := readMaze("maze/maze_temp.in")

	//1.验证maze读取时候正确
	//for _, row := range maze {
	//	for _, val := range row {
	//		fmt.Printf("%d ", val)
	//	}
	//	fmt.Println()
	//}

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	//2.验证steps走过的路径是否对
	for _, row := range steps {
		for _, val := range row {
			//3位对齐
			fmt.Printf("%3d ", val)
		}
		fmt.Println()
	}
}
