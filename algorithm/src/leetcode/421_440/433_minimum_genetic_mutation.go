package _21_440

/**
  @author: CodeWater
  @since: 2024/2/6
  @desc: 最小基因变化
**/

// 图论 bfs 一个单词为一个点，到下一个点边权为1，下一个点从bank中拿
func minMutation(start string, end string, bank []string) int {
	S := make(map[string]bool)
	for _, v := range bank {
		S[v] = true
	}
	dist, q := make(map[string]int), make([]string, 0)
	q, dist[start] = append(q, start), 0
	chrs := []byte{'A', 'T', 'C', 'G'}
	for len(q) > 0 {
		t := q[0]
		q = q[1:]
		for i := 0; i < len(t); i++ {
			s := []byte(t)           // 转成byte数组，方便下面修改，string不能直接修改
			for _, v := range chrs { // 遍历t单词所有变化的可能
				s[i] = v
				str := string(s)
				// 如果存在于bank中并且到这个点单词为0，没有走过，更新
				if _, ok := dist[str]; !ok && S[str] { // if内部变量声明语句，且必须在最前
					dist[str] = dist[t] + 1
					if str == end { // 找到解
						return dist[str]
					}
					q = append(q, str)
				}
			}
		}
	}
	return -1
}
