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
	主要作用于某一段区间加上一个值
**/

const N int = 100010

var (
	n, m   int
	a, b   [N]int // a数组表示原始数组，b数组表示差分数组；也就是说，a数组是b数组的前缀和数组
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
	data := doScan(reader)
	n, _ = strconv.Atoi(data[0])
	m, _ = strconv.Atoi(data[1])
	data = doScan(reader)
	for i := 1; i <= n; i++ {
		a[i], _ = strconv.Atoi(data[i-1])
		// 这里构造差分数组
		insert(i, i, a[i])
	}

	for ; m > 0; m-- {
		data = doScan(reader)
		l, _ := strconv.Atoi(data[0])
		r, _ := strconv.Atoi(data[1])
		c, _ := strconv.Atoi(data[2])
		insert(l, r, c)
	}

	for i := 1; i <= n; i++ {
		b[i] += b[i-1]
		fmt.Fprintf(writer, "%d ", b[i])
	}
	defer writer.Flush()

}

// ==================== 优化string转化int ====================
//package main
//
//import (
//"fmt"
//"bufio"
//"os"
//"strings"
//"strconv"
//)

const N int = 100010

var (
	n, m   int
	a, b   [N]int
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func doScan(reader *bufio.Reader) (intArr []int) {
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	// 修改原来返回string切片改为int切片
	strArr := strings.Split(str, " ")
	for _, v := range strArr {
		value, _ := strconv.Atoi(v)
		intArr = append(intArr, value)
	}
	return
}

func insert(l, r, c int) {
	b[l] += c
	b[r+1] -= c
}

func main() {
	data := doScan(reader)
	n, m = data[0], data[1]
	data = doScan(reader)
	for i := 1; i <= n; i++ {
		a[i] = data[i-1]
		insert(i, i, a[i])
	}

	for ; m > 0; m-- {
		data = doScan(reader)
		l, r, c := data[0], data[1], data[2]
		insert(l, r, c)
	}

	for i := 1; i <= n; i++ {
		b[i] += b[i-1]
		fmt.Fprintf(writer, "%d ", b[i])
	}
	defer writer.Flush()

}

// =================================================
