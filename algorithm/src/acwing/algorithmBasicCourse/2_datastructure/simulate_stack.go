package __datastructure

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/12/19
  @desc: 模拟栈
**/

const N int = 100010

var (
	m, tt          int
	stk            [N]int
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
			stk[tt] = x
		} else if op == "pop" {
			tt--
		} else if op == "empty" {
			if tt > 0 {
				fmt.Fprintln(writer, "NO")
			} else {
				fmt.Fprintln(writer, "YES")
			}
		} else {
			fmt.Fprintln(writer, stk[tt])
		}
	}
}
