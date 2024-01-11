package __search_and_graph_theory

import "fmt"

/**
  @author: CodeWater
  @since: 2023/12/22
  @desc: n皇后
**/

// ================== 第一种搜索方式：逐行搜索 ==================
const N int = 10

var (
	n        int
	row, col [N]bool     // 标记某行和某列是否遍历过
	dg, udg  [2 * N]bool // 标记某个对角线和反对角线是否遍历过
	g        [N][N]byte  // 存储皇后放置的位置
)

// dfs (x,y)当前遍历到的点的坐标 s遍历到第几个
func dfs(x, y, s int) {
	if s > n {
		return
	}
	// 遍历到一行的末尾处，转化xy下标到下一行
	if y == n {
		x, y = x+1, 0
	}
	// 遍历到最后一行
	if x == n {
		// 找到一组解
		if s == n {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					fmt.Printf("%c", g[i][j])
				}
				fmt.Println()
			}
			fmt.Println()
		}
		return
	}

	// 初始化当前位置，标记为未放置皇后
	g[x][y] = '.'
	// 往下一行搜（如果到了一行末尾，上面会处理到下一行）
	dfs(x, y+1, s) //此时还没有放置皇后，所以s不用加1

	// 如果当前行、列、对角线、反对角线都没有放置过皇后，那么标记该点合适
	if !row[x] && !col[y] && !dg[y-x+n] && !udg[y+x] {
		row[x], col[y], dg[y-x+n], udg[y+x] = true, true, true, true
		g[x][y] = 'Q'
		// 继续下一行，放置皇后s+1
		dfs(x, y+1, s+1)
		// 回溯，恢复现场
		g[x][y] = '.'
		row[x], col[y], dg[y-x+n], udg[y+x] = false, false, false, false
	}
}

func main() {
	fmt.Scan(&n)
	dfs(0, 0, 0)
}

/* 第二种搜索方式：搜索每一行的每一列

package main

import (
    "fmt"
)

const N int = 10

var (
    n int
    g [N][N]byte
    col [N]bool
    dg , udg [N * 2]bool
)

// dfs u相当于当前行或者是放置到的第几个皇后
func dfs(u int) {
    //找到一组解
    if u == n {
        for i := 0 ; i < n ; i++ {
            for j := 0 ; j < n ; j++ {
                fmt.Printf("%c" , g[i][j])
            }
            fmt.Println()
        }
        fmt.Println()
        return
    }

    // 这里遍历的是当前行的每一列，判断这列是否可以放置皇后
    for i := 0 ; i < n ; i++ {
        //对角线计算可以画个坐标系，然后计算下截距b即可（+n是因为防止变成负数）
        //u相当于y，i相当于x，y=x+b相当于u=i+b,所以b=u-i
        if !col[i] && !dg[u - i + n] && !udg[u + i] {
            g[u][i] = 'Q'
            col[i] , dg[u - i + n] , udg[u + i] = true , true , true
            dfs(u + 1)
            col[i] , dg[u - i + n] , udg[u + i] = false , false , false
            g[u][i] = '.'
        }
    }
}


func main() {
    fmt.Scan(&n)
    for i := 0 ; i < n ; i++ {
        for j := 0 ; j < n ; j++ {
            g[i][j] = '.'
        }
    }

    dfs(0)
}
*/
