package _1_40

/**
  @author: CodeWater
  @since: 2023/12/30
  @desc: 30. 串联所有单词的子串
**/

func findSubstring(s string, words []string) []int {
	res := []int{}
	if len(words) == 0 {
		return res
	}
	n, m, w := len(s), len(words), len(words[0])
	tot := make(map[string]int)
	for _, word := range words {
		tot[word]++
	}

	for i := 0; i < w; i++ {
		wd, cnt := make(map[string]int), 0
		for j := i; j+w <= n; j = j + w {
			if j >= i+m*w {
				word := s[j-m*w : j-(m-1)*w] //j-m*w后面w个长度
				wd[word]--
				if wd[word] < tot[word] {
					cnt--
				}
			}
			word := s[j : j+w]
			wd[word]++
			if wd[word] <= tot[word] {
				cnt++
			}
			if cnt == m {
				res = append(res, j-(m-1)*w)
			}
		}
	}
	return res
}
