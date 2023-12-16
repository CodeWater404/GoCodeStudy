package __basicAlgorithm

import "fmt"

/**
  @author: CodeWater
  @since: 2023/9/9
  @desc: 最长连续不重复子序列
**/

const N int = 100010

var (
	// n个数
	n int
	//q数组，s存储q数组里面每个元素出现多少次（下标是元素值，数组值是次数;如果题目中数组过大，那么可以开哈希表来进行存储）
	q, s [N]int
)

func main() {
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}

	res := 0
	// i是右端点，j是左端点（意思是：最长不重复序列离i的距离）；j————i
	for i, j := 0, 0; i < n; i++ {
		// q[i]元素出现次数加1
		s[q[i]]++
		// 判断是否有元素出现大于1次，只判断q[i]元素是因为i每次往右走的时候，出现的重复元素只会是q[i]
		for j < i && s[q[i]] > 1 {
			// 上面是右端点出现重复，这里收缩左端点j是因为j是头部，还需要往后继续扫描是否出现最长的不重复的子序列，所以不能在s[q[i]]--,
			s[q[j]]--
			j++
		}
		// 更新最长距离
		res = max(res, i-j+1)
	}
	fmt.Println(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
