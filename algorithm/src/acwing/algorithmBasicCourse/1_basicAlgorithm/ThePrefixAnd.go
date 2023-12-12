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
  @since: 2023/8/25
  @desc: 前缀和
**/

/*
const N int = 100010

var (
	n int
	m int
	a [N]int
	s [N]int
)

func main() {

	fmt.Scan(&n, &m)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a[i])
	}
	for i := 1; i <= n; i++ {
		s[i] = a[i] + s[i-1]
	}

	var l, r int
	for m > 0 {
		fmt.Scan(&l, &r)
		fmt.Println(s[r] - s[l-1])
		m--
	}
}
*/

// ==================== 优化io ====================

const N int = 100010

var (
	n, m   int
	a, s   = make([]int, N), make([]int, N)
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func doScan(reader *bufio.Reader) []string {
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	return strings.Split(str, " ")
}

func main() {
	data := doScan(reader)
	n, _ = strconv.Atoi(data[0])
	m, _ = strconv.Atoi(data[1])
	data = doScan(reader)
	for i := 1; i <= n; i++ {
		a[i], _ = strconv.Atoi(data[i-1])
		s[i] = s[i-1] + a[i]
	}

	for m > 0 {
		data = doScan(reader)
		l, _ := strconv.Atoi(data[0])
		r, _ := strconv.Atoi(data[1])
		fmt.Fprintf(writer, "%d\n", s[r]-s[l-1])
		m--
	}
	defer writer.Flush()
}
