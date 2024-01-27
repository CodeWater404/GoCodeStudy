package __datastructure

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2024/1/27
  @desc: 堆排序
**/

const N = 100010

var (
	n, m, cnt      int
	h              [N]int
	reader, writer = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
)

func down(u int) {
	t := u
	if u*2 <= cnt && h[u*2] < h[t] {
		t = u * 2
	}
	if u*2+1 <= cnt && h[u*2+1] < h[t] {
		t = u*2 + 1
	}
	if u != t {
		h[u], h[t] = h[t], h[u]
		down(t)
	}
}

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &n, &m)
	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &h[i])
	}
	cnt = n
	for i := n / 2; i > 0; i-- {
		down(i)
	}

	for ; m > 0; m-- {
		fmt.Fprintf(writer, "%d ", h[1])
		h[1] = h[cnt]
		cnt--
		down(1)
	}
	fmt.Fprintln(writer)
}
