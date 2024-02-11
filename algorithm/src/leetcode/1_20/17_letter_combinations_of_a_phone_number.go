package __20

/**
  @author: CodeWater
  @since: 2024/2/11
  @desc: 17. 电话号码的字母组合
**/

var (
	res  []string
	strs = []string{
		"", "", "abc", "def", "ghi", "jkl", "mno",
		"pqrs", "tuv", "wxyz",
	}
)

func letterCombinations(digits string) []string {
	res = make([]string, 0)
	if len(digits) == 0 {
		return res
	}
	dfs(digits, 0, "")
	return res
}

func dfs(digits string, u int, path string) {
	if u == len(digits) {
		res = append(res, path)
	} else {
		for _, c := range strs[digits[u]-'0'] {
			dfs(digits, u+1, path+string(c))
		}
	}
}
