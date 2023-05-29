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
  @desc: 归并排序
**/

const N = int(1e5) + 10

var q, temp = make([]int, N), make([]int, N)

func mergeSort(q []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(q, l, mid)
	mergeSort(q, mid+1, r)
	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if q[i] < q[j] {
			temp[k] = q[i]
			i++
		} else {
			temp[k] = q[j]
			j++
		}
		k++
	}
	for i <= mid {
		temp[k] = q[i]
		k, i = k+1, i+1
	}
	for j <= r {
		temp[k] = q[j]
		k, j = k+1, j+1
	}

	for i, j = l, 0; i <= r; i, j = i+1, j+1 {
		q[i] = temp[j]
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	//也可以直接写成： n , _ := strconv.Atoi(string(str[0]))
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

	mergeSort(q, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%d ", q[i])
	}

}
