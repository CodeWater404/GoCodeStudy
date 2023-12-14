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
  @desc: 数的范围
	bufio比fmt快十倍
**/

const N = 1e5 + 10

var (
	n      int
	m      int
	q      = make([]int, N)
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func doScann(read *bufio.Reader) string {
	str, _ := read.ReadString('\n')
	str = strings.TrimSpace(str)
	return str
}

func main() {
	data := strings.Split(doScann(reader), " ")
	n, _ := strconv.Atoi(data[0])
	m, _ := strconv.Atoi(data[1])
	// fmt.Scanf("%d%d" , &n , &m)

	data = strings.Split(doScann(reader), " ")
	for i := 0; i < n; i++ {
		q[i], _ = strconv.Atoi(data[i])
		// fmt.Scanf("%d" , &q[i])
	}

	for m > 0 {
		m--
		x, _ := strconv.Atoi(doScann(reader))
		// var x int
		// fmt.Scanf("%d" , &x)

		l, r := 0, n-1
		for l < r {
			mid := (l + r) >> 1
			if q[mid] >= x {
				r = mid
			} else {
				l = mid + 1
			}

		}
		if q[l] != x {
			fmt.Fprintln(writer, "-1 -1")
			writer.Flush()

			// fmt.Println("-1 -1")
		} else {
			fmt.Fprint(writer, l, " ")
			writer.Flush()
			// fmt.Print(l , " ")
			ll, rr := 0, n-1
			for ll < rr {
				mid := (ll + rr + 1) >> 1
				if q[mid] <= x {
					ll = mid
				} else {
					rr = mid - 1
				}
			}
			fmt.Fprintln(writer, rr)
			writer.Flush()
			// fmt.Println(rr)
		}
	}
}

// ==================== 优化buff读取输出 ====================
const N int = 100010

var (
	n, q   int
	a      [N]int
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func doScan(reader *bufio.Reader) (intArr []int) {
	str, _ := reader.ReadString('\n')
	strArr := strings.Split(strings.TrimSpace(str), " ")
	for _, v := range strArr {
		value, _ := strconv.Atoi(v)
		intArr = append(intArr, value)
	}
	return
}

func main() {
	data := doScan(reader)
	n, q = data[0], data[1]
	data = doScan(reader)
	for i := 0; i < n; i++ {
		a[i] = data[i]
	}

	for ; q > 0; q-- {
		data = doScan(reader)
		k, l, r := data[0], 0, n-1
		for l < r {
			mid := (l + r) >> 1
			if a[mid] >= k {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if a[l] != k {
			fmt.Fprintf(writer, "-1 -1\n")
		} else {
			fmt.Fprintf(writer, "%d ", l)
			ll, rr := 0, n-1
			for ll < rr {
				mid := (ll + rr + 1) >> 1
				if a[mid] <= k {
					ll = mid
				} else {
					rr = mid - 1
				}
			}
			fmt.Fprintf(writer, "%d\n", ll)
		}
	}
	defer writer.Flush()
}
