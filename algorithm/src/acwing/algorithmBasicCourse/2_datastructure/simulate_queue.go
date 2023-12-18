package __datastructure

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/12/19
  @desc: 模拟队列
**/

const N int = 100010

var (
	m, hh          int
	tt             = -1
	q              [N]int
	reader, writer = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()

	fmt.Fscan(reader, &m)
	for ; m > 0; m-- {
		var op string
		var x int
		fmt.Fscan(reader, &op)
		if op == "push" {
			fmt.Fscan(reader, &x)
			tt++
			q[tt] = x
		} else if op == "pop" {
			hh++
		} else if op == "query" {
			fmt.Fprintln(writer, q[hh])
		} else {
			if hh <= tt {
				fmt.Fprintln(writer, "NO")
			} else {
				fmt.Fprintln(writer, "YES")
			}
		}
	}
}
