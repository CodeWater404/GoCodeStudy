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
			fmt.Printf("============>one line:%d %d %d \n", i, j, maze[i][j])
		}
	}
	return maze
}

func main() {
	maze := readMaze("maze/maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
