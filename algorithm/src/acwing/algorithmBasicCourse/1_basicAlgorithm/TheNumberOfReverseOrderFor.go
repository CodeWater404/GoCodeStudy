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
  @since: 2023/5/29
  @desc: 逆序对的数量
**/

const N = int(1e5) + 10

var q, temp = make([]int, N), make([]int, N)

func mergeSort(q []int, l, r int) int {
	if l >= r {
		return 0
	}
	mid := (l + r) >> 1
	res := mergeSort(q, l, mid) + mergeSort(q, mid+1, r)

	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if q[i] <= q[j] {
			temp[k] = q[i]
			i++
		} else {
			temp[k] = q[j]
			j++
			res = res + mid - i + 1
		}
		k++
	}

	for i <= mid {
		temp[k] = q[i]
		i++
		k++
	}
	for j <= r {
		temp[k] = q[j]
		j++
		k++
	}

	for i, j = l, 0; i <= r; i, j = i+1, j+1 {
		q[i] = temp[j]
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	data := strings.Split(str, " ")
	n, _ := strconv.Atoi(data[0])

	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	data = strings.Split(str, " ")
	for i := 0; i < n; i++ {
		q[i], _ = strconv.Atoi(data[i])
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintln(writer, mergeSort(q, 0, n-1))
}
