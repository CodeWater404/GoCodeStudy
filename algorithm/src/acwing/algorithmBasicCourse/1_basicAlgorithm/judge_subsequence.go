package __basicAlgorithm

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/12/18
  @desc:判断子序列
**/

const N int = 100010

var (
	n, m           int
	a, b           [N]int
	reader, writer = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &n, &m)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}

	i, j := 0, 0
	for i < n && j < m {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	if i == n {
		fmt.Fprintln(writer, "Yes")
	} else {
		fmt.Fprintln(writer, "No")
	}
}
