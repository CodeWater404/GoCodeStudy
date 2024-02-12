package _1_40

/**
  @author: CodeWater
  @since: 2024/2/12
  @desc: 22. 括号生成
**/

var ans []string

func generateParenthesis(n int) []string {
	ans = make([]string, 0)
	dfs(n, 0, 0, "")
	return ans
}

func dfs(n, lc, rc int, seq string) {
	if lc == n && rc == n {
		ans = append(ans, seq)
	} else {
		if lc < n { // 只要左括号有就填
			dfs(n, lc+1, rc, seq+"(")
		}
		if rc < n && lc > rc { // 右括号必须小于左括号数量的时候才可以填
			dfs(n, lc, rc+1, seq+")")
		}
	}
}
