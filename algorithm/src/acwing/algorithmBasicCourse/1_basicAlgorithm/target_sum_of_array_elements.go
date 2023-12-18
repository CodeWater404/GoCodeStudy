package __basicAlgorithm

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/12/18
  @desc:数组元素的目标和
**/

const N int = 100010

var (
	n, m, x        int
	A, B           [N]int
	reader, writer = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	fmt.Fscan(reader, &n, &m, &x)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &A[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &B[i])
	}

	for i, j := 0, m-1; i < n; i++ {
		for j >= 0 && A[i]+B[j] > x {
			j--
		}
		if j >= 0 && A[i]+B[j] == x {
			fmt.Fprintf(writer, "%d %d\n", i, j)
			return
		}
	}
}
