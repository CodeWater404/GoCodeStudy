package __datastructure

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/12/25
  @desc: 合并集合
**/

const N int = 100010

var (
	p              [N]int
	n, m           int
	reader, writer = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
)

func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x])
	}
	return p[x]
}

func main() {
	defer writer.Flush()

	fmt.Fscan(reader, &n, &m)
	for i := 1; i <= n; i++ {
		p[i] = i
	}

	for ; m > 0; m-- {
		var op string
		var a, b int
		fmt.Fscan(reader, &op, &a, &b)
		if op == "M" {
			p[find(a)] = find(b)
		} else {
			if find(a) == find(b) {
				fmt.Fprintln(writer, "Yes")
			} else {
				fmt.Fprintln(writer, "No")
			}
		}
	}
}
