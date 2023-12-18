package __datastructure

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/12/19
  @desc: 单调栈
**/

const N int = 100010

var (
	n, tt          int
	stk            [N]int
	reader, writer = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()

	fmt.Fscan(reader, &n)
	for ; n > 0; n-- {
		var x int
		fmt.Fscan(reader, &x)
		for tt > 0 && stk[tt] >= x {
			tt--
		}
		if tt > 0 {
			fmt.Fprintf(writer, "%d ", stk[tt])
		} else {
			fmt.Fprintf(writer, "-1 ")
		}
		tt++
		stk[tt] = x
	}
}
