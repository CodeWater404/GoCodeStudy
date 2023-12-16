package __basicAlgorithm

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/**
  @author: CodeWater
  @since: 2023/12/16
  @desc: 区间合并
**/

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type pair struct {
	first  int
	second int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(segs []pair) []pair {
	var res []pair

	sort.Slice(segs, func(i, j int) bool {
		return segs[i].first < segs[j].first
	})
	// 初始左右边界为无穷大
	st, ed := int(-2e9), int(-2e9)
	for _, v := range segs {
		// 当前区间的右端点和遍历到的区间左端点不重合
		if ed < v.first {
			// 当前区间左端点不为无穷大，是一种解
			if st != int(-2e9) {
				res = append(res, pair{st, ed})
			}
			// 更新当前区间的左右端点
			st, ed = v.first, v.second
		} else {
			// // 当前区间的右端点和遍历到的区间左端点重合,更新当前区间的右端点
			ed = max(ed, v.second)
			// fmt.Printf("k:%d %T, v:%d %T" , first , first , second , second)
		}
	}
	// 最后一种情况加入，防止不为空
	if st != int(-2e9) {
		res = append(res, pair{st, ed})
	}
	return res
}

func main() {
	var n int
	fmt.Fscan(reader, &n)
	var segs []pair
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		segs = append(segs, pair{l, r})
	}

	segs = merge(segs)

	fmt.Fprintln(writer, len(segs))
	writer.Flush()
}
