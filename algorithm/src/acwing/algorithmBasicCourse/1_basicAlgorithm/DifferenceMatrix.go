package __basicAlgorithm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
  @author: CodeWater
  @since: 2023/9/8
  @desc: 差分矩阵
**/

const N int = 1010

var (
	m, n, q int
	a, b    [N][N]int
	reader  = bufio.NewReader(os.Stdin)
	writer  = bufio.NewWriter(os.Stdout)
)

func doScan(reader *bufio.Reader) []string {
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	return strings.Split(str, " ")
}

func insert(x1, y1, x2, y2, c int) {
	b[x1][y1] += c
	b[x2+1][y1] -= c
	b[x1][y2+1] -= c
	b[x2+1][y2+1] += c
}

func main() {
	// fmt超时
	// fmt.Scanf("%d%d%d" , &n , &m , &q)
	data := doScan(reader)
	n, _ = strconv.Atoi(data[0])
	m, _ = strconv.Atoi(data[1])
	q, _ = strconv.Atoi(data[2])

	for i := 1; i <= n; i++ {
		data = doScan(reader)
		for j := 1; j <= m; j++ {
			a[i][j], _ = strconv.Atoi(data[j-1])
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			insert(i, j, i, j, a[i][j])
		}
	}

	for q > 0 {
		data = doScan(reader)
		x1, _ := strconv.Atoi(data[0])
		y1, _ := strconv.Atoi(data[1])
		x2, _ := strconv.Atoi(data[2])
		y2, _ := strconv.Atoi(data[3])
		c, _ := strconv.Atoi(data[4])
		insert(x1, y1, x2, y2, c)
		q--
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			b[i][j] += b[i-1][j] + b[i][j-1] - b[i-1][j-1]
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Fprintf(writer, "%d ", b[i][j])
		}
		fmt.Fprintf(writer, "\n")
	}
	defer writer.Flush()
}
