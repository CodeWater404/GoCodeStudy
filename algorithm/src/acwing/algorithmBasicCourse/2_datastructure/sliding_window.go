package __datastructure

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/12/21
  @desc: 滑动窗口
**/

const N int = 1000010

var (
	// a存储数组；q模拟队列
	a, q           [N]int
	reader, writer = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n, k int
	fmt.Fscan(reader, &n, &k)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	//hh队头；tt队尾
	hh, tt := 0, -1
	//先求窗口内的最小值
	for i := 0; i < n; i++ {
		// 队列存在值并且窗口长度大于对头存储的下标值，存储长度过长移除元素
		if hh <= tt && i-k+1 > q[hh] {
			hh++
		}
		// 队列存在值并且队尾元素大于等于当前的值，那么一直缩小队列
		for hh <= tt && a[q[tt]] >= a[i] {
			tt--
		}
		//存下当前值比前面小的数
		tt++
		q[tt] = i
		//如果窗口满足长度了（一开始窗口可能只有一个），输出当前窗口内最小的值队头
		//（单调队列，当前是递增的）
		if i >= k-1 {
			fmt.Fprintf(writer, "%d ", a[q[hh]])
		}
	}
	fmt.Fprintf(writer, "\n")

	//找窗口内最大的元素，这个时候就要保持队列呈现递减的趋势
	hh, tt = 0, -1
	for i := 0; i < n; i++ {
		if hh <= tt && i-k+1 > q[hh] {
			hh++
		}
		for hh <= tt && a[q[tt]] <= a[i] {
			tt--
		}
		tt++
		q[tt] = i
		if i >= k-1 {
			fmt.Fprintf(writer, "%d ", a[q[hh]])
		}
	}
	fmt.Fprintf(writer, "\n")
}
