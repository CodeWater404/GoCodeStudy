package __datastructure

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2024/1/12
  @desc: 143 最大异或对
**/

const N, M int = 100010, 3100010

var (
	n, idx         int
	a              [N]int
	son            [M][2]int
	reader, writer = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
)

func insert(x int) {
	p := 0
	for i := 30; i >= 0; i-- {
		s := &son[p][x>>i&1]
		if *s == 0 {
			idx++
			*s = idx
		}
		p = *s
	}

}

func search(x int) int {
	p, res := 0, 0
	for i := 30; i >= 0; i-- {
		s := x >> i & 1
		// if son[p][!s] > 0 {
		if son[p][s^1] > 0 { //s取反
			res += 1 << i
			p = son[p][s^1]
		} else {
			p = son[p][s]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		insert(a[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		res = max(res, search(a[i]))
	}

	fmt.Fprintf(writer, "%d\n", res)
	writer.Flush()
}
