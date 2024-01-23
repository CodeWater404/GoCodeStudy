package __datastructure

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2024/1/23
  @desc: 837. 连通块中点的数量
**/

const (
	N = 100010
	M = 100010
)

var (
	n, m   int
	p, cnt [N]int
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
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
		cnt[i] = 1
	}

	for ; m > 0; m-- {
		var (
			op   string
			a, b int
		)
		fmt.Fscan(reader, &op)
		if op == "C" {
			fmt.Fscan(reader, &a, &b)
			a, b = find(a), find(b)
			if a != b {
				p[a] = b
				cnt[b] += cnt[a]
			}
		} else if op == "Q1" {
			fmt.Fscan(reader, &a, &b)
			if find(a) == find(b) {
				fmt.Fprintln(writer, "Yes")
			} else {
				fmt.Fprintln(writer, "No")
			}
		} else {
			fmt.Fscan(reader, &a)
			fmt.Fprintln(writer, cnt[find(a)])
		}
	}
}
