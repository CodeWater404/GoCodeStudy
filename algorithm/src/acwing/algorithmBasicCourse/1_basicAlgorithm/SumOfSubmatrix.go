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
  @since: 2023/8/27
  @desc: 子矩阵的和
**/

const N int = 1010

var (
	n      int
	m      int
	q      int
	s      [N][N]int
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func doScan(reader *bufio.Reader) []string {
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	data := strings.Split(str, " ")
	return data
}

func main() {
	// 超时
	// fmt.Scan(&n , &m , &q)
	data := doScan(reader)
	n, _ := strconv.Atoi(data[0])
	m, _ := strconv.Atoi(data[1])
	q, _ := strconv.Atoi(data[2])

	for i := 1; i <= n; i++ {
		data = doScan(reader)
		for j := 1; j <= m; j++ {
			s[i][j], _ = strconv.Atoi(data[j-1])
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			s[i][j] += s[i-1][j] + s[i][j-1] - s[i-1][j-1]
		}
	}

	for q > 0 {
		data = doScan(reader)
		x1, _ := strconv.Atoi(data[0])
		y1, _ := strconv.Atoi(data[1])
		x2, _ := strconv.Atoi(data[2])
		y2, _ := strconv.Atoi(data[3])
		fmt.Fprintln(writer, s[x2][y2]-s[x1-1][y2]-s[x2][y1-1]+s[x1-1][y1-1])
		q--
	}
	defer writer.Flush()

}
