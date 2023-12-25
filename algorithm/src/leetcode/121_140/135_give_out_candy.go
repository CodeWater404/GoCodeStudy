package _21_140

/*
*

	@author: CodeWater
	@since: 2023/12/25
	@desc: 分发糖果

*
*/
var (
	f, w []int //f[i]表示i这个位置给的糖果数；w就是每个孩子的评分
	n    int
)

func candy(ratings []int) int {
	w, n = ratings, len(ratings)
	f = make([]int, n)

	// 初始化
	for i := 0; i < n; i++ {
		f[i] = -1
	}

	res := 0
	for i := 0; i < n; i++ {
		//把每个孩子给的糖果数相加就是总和
		res += dp(i)
	}
	return res
}

// dp 求出x这个位置的孩子给的糖果数
func dp(x int) int {
	// f[x]不为-1，说明这个位置已经求过了，直接返回
	if f[x] != -1 {
		return f[x]
	}
	// 每个位置起码给1个糖果
	f[x] = 1
	//x比x-1位置相比评分高，高的赋值给当前f[x] ; 下同
	if x > 0 && w[x-1] < w[x] {
		f[x] = max(f[x], dp(x-1)+1)
	}
	if x+1 < n && w[x+1] < w[x] {
		f[x] = max(f[x], dp(x+1)+1)
	}
	return f[x]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
