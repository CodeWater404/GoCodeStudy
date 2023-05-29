package _461_2480

import (
	"container/heap"
	"sort"
)

/**
  @author: CodeWater
  @since: 2023/5/29
  @desc:
**/
func totalCost(costs []int, k, candidates int) int64 {
	ans := 0
	//如果前后两个candidates区间不重合，那就用堆模拟；重合的话就直接排序，用前k个最小的返回代价总和
	if n := len(costs); candidates*2 < n {
		pre := hp{costs[:candidates]}
		heap.Init(&pre) // 原地建堆
		suf := hp{costs[n-candidates:]}
		heap.Init(&suf)
		for i, j := candidates, n-1-candidates; k > 0 && i <= j; k-- {
			if pre.IntSlice[0] <= suf.IntSlice[0] {
				ans += pre.IntSlice[0]
				pre.IntSlice[0] = costs[i]
				heap.Fix(&pre, 0)
				i++
			} else {
				ans += suf.IntSlice[0]
				suf.IntSlice[0] = costs[j]
				heap.Fix(&suf, 0)
				j--
			}
		}
		//三个点 ... 表示切片的展开操作符（slice spread）。它用于将一个切片或数组展开成独立的元素，可以作为函数参数传递或在切片拼接时使用。
		//pre和suf不会有重复的元素，因为上面for在遍历两个区间重合的时候，就退出了，所以不会有
		costs = append(pre.IntSlice, suf.IntSlice...)
	}
	sort.Ints(costs)
	//这里再加上还剩k轮的代价，上面代码中k是更新的
	for _, c := range costs[:k] { // 也可以用快速选择算法求前 k 小
		ans += c
	}
	return int64(ans)
}

type hp struct{ sort.IntSlice }

func (hp) Push(interface{})     {} // 没有用到，留空即可
func (hp) Pop() (_ interface{}) { return }
