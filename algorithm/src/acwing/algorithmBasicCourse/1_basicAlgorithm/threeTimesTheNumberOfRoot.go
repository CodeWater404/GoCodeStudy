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
  @since: 2023/7/24
  @desc: 数的三次方根
**/

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func doScan(read *bufio.Reader) string {
	str, _ := read.ReadString('\n')
	str = strings.TrimSpace(str)
	return str
}

func main() {
	n, _ := strconv.ParseFloat(doScan(reader), 64)
	var (
		l float64
		r float64
	)
	l, r = -10000, 10000
	for r-l > 1e-8 {
		mid := (l + r) / 2
		if mid*mid*mid >= n {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Fprintf(writer, "%.6f", l)
	defer writer.Flush()
}
