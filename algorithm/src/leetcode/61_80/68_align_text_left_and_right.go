package _1_80

import "strings"

/*
*

	@author: CodeWater
	@since: 2023/12/26
	@desc: 68. 文本左右对齐

*
*/
func fullJustify(words []string, maxWidth int) []string {
	var res []string
	for i := 0; i < len(words); i++ {
		j, size := i+1, len(words[i])
		// 下一个单词和当前单词加上空格不超过maxWidth，更新j和size
		for j < len(words) && size+1+len(words[j]) <= maxWidth {
			size += 1 + len(words[j])
			j++
		}
		// 开始放入一行结果中
		var line string
		// 左对齐的情况：一行只能放下一个单词；到最后一行了
		if j == len(words) || j == i+1 {
			//先加入当前单词，然后扩展后续可以加入的单词
			line += string(words[i])
			for k := i + 1; k < j; k++ {
				line += " " + words[k]
			}
			for len(line) < maxWidth {
				line += " "
			}
		} else { // 左右对齐情况
			cnt := j - i - 1           // 空隙个数（比如3个单词有2个间隙），j是终止单词下标
			r := maxWidth - size + cnt //剩余可以加入的空格个数
			line += words[i]
			k := 0 //指向单词的下标
			for k < r%cnt {
				//左半边加入的空格要多，所以加1
				line += strings.Repeat(" ", r/cnt+1) + words[i+k+1]
				k++
			}
			for k < cnt {
				line += strings.Repeat(" ", r/cnt) + words[i+k+1]
				k++
			}
		}
		// 加入一行结果，更新i位置
		res = append(res, line)
		i = j - 1
	}
	return res
}
