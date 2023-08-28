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
  @since: 2023/8/28
  @desc: 差分
**/

const N int = 100010

var (
	n      int
	m      int
	a      [N]int
	b      [N]int
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func doScan(reader *bufio.Reader) []string {
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	return strings.Split(str, " ")
}

func insert(l, r, c int) {
	b[l] += c
	b[r+1] -= c
}

func main() {
	// 超时
	// fmt.Scan(&n , &m)
	data := doScan(reader)
	n, _ := strconv.Atoi(data[0])
	m, _ := strconv.Atoi(data[1])

	data = doScan(reader)
	for i := 1; i <= n; i++ {
		a[i], _ = strconv.Atoi(data[i-1])
	}

	for i := 1; i <= n; i++ {
		insert(i, i, a[i])
	}

	for m > 0 {
		data = doScan(reader)
		l, _ := strconv.Atoi(data[0])
		r, _ := strconv.Atoi(data[1])
		c, _ := strconv.Atoi(data[2])
		insert(l, r, c)
		m--
	}

	for i := 1; i <= n; i++ {
		b[i] += b[i-1]
	}

	for i := 1; i <= n; i++ {
		fmt.Fprintf(writer, "%d ", b[i])
	}
	defer writer.Flush()
}
