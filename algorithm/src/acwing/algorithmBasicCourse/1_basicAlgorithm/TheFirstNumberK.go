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
  @since: 2023/5/28
  @desc: $
**/

const N = int(1e5) + 10

/*PS
1. var q = [N]int{}数组类型
2，var q = make([]int , N)切片类型,指定了切片的长度为 N。

//q []int切片类型。如果声明的是数组类型，需要把数组转换成切片，才能正确的传参！
//qSlice := q[:]
func quickSort(q []int , l , r , k int) int {}
*/
// var q = [N]int{}
var q = make([]int, N)

func quickSort(q []int, l, r, k int) int {
	if l >= r {
		return q[l]
	}
	x, i, j := q[(l+r)>>1], l-1, r+1
	for i < j {
		for {
			i++
			if q[i] >= x {
				break
			}
		}
		for {
			j--
			if q[j] <= x {
				break
			}
		}
		if i < j {
			temp := q[i]
			q[i] = q[j]
			q[j] = temp
		}
	}

	if k <= j-l+1 {
		return quickSort(q, l, j, k)
	} else {
		return quickSort(q, j+1, r, k-(j-l+1))
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	data := strings.Split(str, " ")
	n, _ := strconv.Atoi(data[0])
	k, _ := strconv.Atoi(data[1])

	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	data = strings.Split(str, " ")
	for i := 0; i < n; i++ {
		q[i], _ = strconv.Atoi(data[i])
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintf(writer, "%d", quickSort(q, 0, n-1, k))

}
