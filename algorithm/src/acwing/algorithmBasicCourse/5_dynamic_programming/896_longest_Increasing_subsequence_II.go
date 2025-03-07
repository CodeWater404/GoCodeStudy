package __dynamic_programming

/*
  @author: CodeWater
  @since: 2025/03/06
  @desc: 896. 最长上升子序列 II
*/

import "fmt"

const N = 100010

var (
	n int
	// a存储序列
	/* q[i]: 长度为i的子序列最后一个数是q[i]存储的值，始终保持严格单调递增。注意：
	q中存储的序列并不是题目答案的序列，而是一个辅助数组，用于存储长度为i的子序列最后一个数的值！注意区分！
	q的实际存储值的长度是最长上升子序列的长度，末尾不断更新是为了保持往后拓展有更多的
	可能性，因为以一个较小的值明显比一个较大的值为结尾有更多的可能性，也就是拓展长度。

	如果后来的值比q末尾最后一个值小，那么会更新到0-length之间的某个位置，这样就保证了
	q中存储的序列是严格单调递增的，同时也保证了q中存储的序列是最长上升子序列的长度。同样，
	这也是为什么q中存储的序列不是最长上升子序列的原因。
	如果后来的值比q末尾最后一个值大，那么就会添加到q的末尾去，相当于长度增加了。
	*/
	a, q [N]int
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	length := 0
	for i := 0; i < n; i++ {
		// 二分查找当前数a[i]在q中的位置
		l, r := 0, length
		for l < r {
			// 注意加括号，加法优先级比位移低
			mid := (l + r + 1) >> 1
			if q[mid] < a[i] {
				l = mid
			} else {
				r = mid - 1
			}
		}
		// 找到之后，按照r来更新length，因为r是q中实际存储的长度
		length = max(length, r+1)
		// 更新长度是r+1的子序列最后一个结尾值
		q[r+1] = a[i]
	}

	fmt.Println(length)
}
