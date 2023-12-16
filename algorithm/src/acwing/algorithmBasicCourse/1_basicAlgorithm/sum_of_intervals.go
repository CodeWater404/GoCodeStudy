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
  @desc: 区间和
	原理：离散化
**/

// n次插入最多1e5个数，m次查询有左右两个端点也都算入离散化的数组中2*1e5，所以量级是3*1e5
const N int = 300010

type pair struct {
	first  int
	second int
}

var (
	n, m int
	a, s [N]int // a存储散列化后的数，最多只会存储N个;s是a的前缀和数组
	// alls存储插入数值的下标和查询操作端点下标，所以里面的元素下标有可能重复===》所以需要去重
	alls []int
	// add存储插入的操作； query存储查询的操作
	add, query []pair
	reader     = bufio.NewReader(os.Stdin)
	writer     = bufio.NewWriter(os.Stdout)
)

// find 二分查找一个数,返回的是下标
func find(x int) int {
	l, r := 0, len(alls)-1
	for l < r {
		mid := (l + r) >> 1
		if alls[mid] >= x {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// +1是因为要求前缀和，前缀和数组下标从1开始比较好算
	return r + 1
}

// unique 对数组中重复的数值去重
func unique(a []int) []int {
	j := 0
	for i := 0; i < len(a); i++ {
		if i == 0 || a[i] != a[i-1] {
			a[j] = a[i]
			j++
		}
	}
	//返回从头到j位置处的元素；j后面的元素都是重复的元素
	return a[:j]
}

func main() {
	// fmt.Scanf("%d%d" , &n , &m)// 超时
	fmt.Fscanf(reader, "%d%d\n", &n, &m) // 直接一点可以fmt.Fscan(&n , &m)
	// 处理输入的操作并存入对应数组
	for i := 0; i < n; i++ {
		var x, c int
		fmt.Fscanf(reader, "%d%d\n", &x, &c)
		add = append(add, pair{x, c})
		alls = append(alls, x)
	}
	for i := 0; i < m; i++ {
		var l, r int
		fmt.Fscanf(reader, "%d%d\n", &l, &r)
		query = append(query, pair{l, r})
		alls = append(alls, l, r)
	}
	// 排序后去重
	sort.Ints(alls)
	alls = unique(alls)

	// 处理插入
	for _, v := range add {
		// 对于指定first位置上加上second的操作，先在alls中找到其对应的下标然后直接作为a的下标
		x := find(v.first)
		a[x] += v.second
	}

	// 处理前缀和（这里遍历长度是去重后的alls，不是遍历a的长度，因为len（a）不一定
	//真有这么多个元素）
	for i := 1; i <= len(alls); i++ {
		s[i] = s[i-1] + a[i]
	}

	// 处理询问
	for _, v := range query {
		// 找到在alls中的对应的下标，然后就可以拿这个下标用前缀和处理了
		l, r := find(v.first), find(v.second)
		// fmt.Printf("%d\n" , s[r] - s[l - 1])
		fmt.Fprintf(writer, "%d\n", s[r]-s[l-1])
	}
	writer.Flush()
}
