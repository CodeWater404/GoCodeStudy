package __20

/*
*

	@author: CodeWater
	@since: 2023/12/26
	@desc: N字型变换

*
*/
func convert(s string, n int) string {
	res := ""
	//特殊情况，直接返回，否则下面会进入死循环
	if n == 1 {
		return s
	}

	for i := 0; i < n; i++ {
		//首先是第一行和最后一行，只有一个等差公式即可计算出下一个位置
		if i == 0 || i == n-1 {
			for j := i; j < len(s); j += 2*n - 2 {
				res += string(s[j])
			}
		} else {
			// 中间的行，有两个分开计算的等差公式算出分开两个位置，位于竖线上和斜线上的
			for j, k := i, 2*n-2-i; j < len(s) || k < len(s); j, k = j+2*n-2, k+2*n-2 {
				if j < len(s) {
					res += string(s[j])
				}
				if k < len(s) {
					res += string(s[k])
				}
			}
		}
	}
	return res
}
