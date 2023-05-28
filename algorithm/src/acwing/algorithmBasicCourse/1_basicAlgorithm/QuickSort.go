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
  @desc: 快速排序

**/

var N = 100010
var n int
var q = make([]int, N)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	// 去除首位空白字符，当然也可以使用strings.TrimRight(line, "\n")
	str = strings.TrimSpace(str)
	data := strings.Split(str, " ")
	// 字符串转换为整数
	n, _ := strconv.Atoi(data[0])

	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	data = strings.Split(str, " ")
	for i := 0; i < n; i++ {
		q[i], _ = strconv.Atoi(data[i])
	}
	quickSort(q, 0, n-1)

	//用bufio由800提升到188ms
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%d ", q[i])
	}
}

func quickSort(q []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l-1, r+1
	x := q[(i+j)>>1]
	for i < j {
		/*捕获异常，debug
		  defer func() {
		      if r := recover(); r != nil {
		           fmt.Printf("Recovered:%s , %d , %d\n", r , i , j)
		      }
		  }()*/
		//cpp: do while转化
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
	quickSort(q, l, j)
	quickSort(q, j+1, r)
}
