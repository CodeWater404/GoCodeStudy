package __datastructure

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2024/1/24
  @desc: 240. 食物链
**/

const N = 50010

var (
	n, m           int
	p, d           [N]int
	reader, writer = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
)

func find(x int) int {
	if p[x] != x {
		t := find(p[x])
		d[x] += d[p[x]]
		p[x] = t
	}
	return p[x]
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &n, &m)
	for i := 1; i <= n; i++ {
		p[i] = i
	}
	res := 0
	for ; m > 0; m-- {
		var t, x, y int
		fmt.Fscan(reader, &t, &x, &y)
		if x > n || y > n {
			res++
		} else {
			px, py := find(x), find(y)
			if t == 1 {
				if px == py && (d[x]-d[y])%3 != 0 {
					res++
				} else if px != py {
					//同类，但是不同集合，把x合并到y集合中，并更新x的根到y的根的距离
					p[px] = py
					/*推导：xy是同类，x这个集合合并到y集合中时，x到本集合根的距离加上本集合到y集合根的距离之后的总和
					模3之后是0。也就是d【x】+ ？模3的余数跟d【y】模3的余数相等，进一步优化就是（d【x】+？-d【y】）
					模3的值为0，在进一步优化就是？ = d【y】-d【x】。？就是x集合根到y集合根的距离，px。
					*/
					d[px] = d[y] - d[x]
				}
			} else {
				//x吃y，说明y+1和x是同类
				if px == py && (d[x]-d[y]-1)%3 != 0 {
					res++
				} else if px != py {
					p[px] = py
					d[px] = d[y] + 1 - d[x]
				}
			}
		}
	}
	fmt.Fprintln(writer, res)
}
