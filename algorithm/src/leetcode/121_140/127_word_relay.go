package _21_140

/**
  @author: CodeWater
  @since: 2024/2/8
  @desc: 单词接龙
**/

// ladderLength 和433题一样
func ladderLength(beginWord string, endWord string, wordList []string) int {
	S := make(map[string]bool) // 记录合法的中间单词变化状态
	for _, v := range wordList {
		S[v] = true
	}
	// dist[a]=1记录单词变化到当前a单词所需的变化1次，q放入每一个可能单词变化的状态，bfs
	dist, q := make(map[string]int), make([]string, 0)
	dist[beginWord], q = 0, append(q, beginWord)

	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		r := t
		for i := 0; i < len(t); i++ {
			rTmp := []rune(r)
			for j := 'a'; j <= 'z'; j++ {
				rTmp[i] = j
				//dist中不存在，但是在S中存在，说明当前rTmp可以走到下一个更新状态，更新dist
				if _, ok := dist[string(rTmp)]; !ok && S[string(rTmp)] {
					dist[string(rTmp)] = dist[r] + 1
					if string(rTmp) == endWord {
						return dist[string(rTmp)] + 1
					}
					q = append(q, string(rTmp))
				}
			}
		}
	}
	return 0
}
