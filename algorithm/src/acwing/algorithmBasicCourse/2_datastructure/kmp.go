package __datastructure

import (
	"bufio"
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/12/18
  @desc: KMP算法
**/

const (
	N int = 100010
	M int = 1000010
)

var (
	n, m           int
	p              [N]byte
	s              [M]byte
	ne             [N]int
	reader, writer = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
)

func main() {
	// 处理输入到byte数组
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 1; i <= n; i++ {
		reader.Read(p[i : i+1])
	}
	//todo: 上面不用reader读取，直接用fmt.Scanf("%s\n", &p[1:])这里是不是就可以不用读取换行了？
	reader.ReadString('\n') // 第二行到第三行有个换行符需要
	fmt.Fscanf(reader, "%d\n", &m)
	for i := 1; i <= m; i++ {
		reader.Read(s[i : i+1])
	}

	//求next数组:i从2，j从0开始（1是开始存储字符的位置，i需要和j+1位置比较，所以
	//i从第二个和j+1从第一个开始比较）
	for i, j := 2, 0; i <= n; i++ {
		for j > 0 && p[i] != p[j+1] {
			j = ne[j]
		}
		// 退出上面循环情况：j无法后退；i和j+1位置处匹配上（考虑这种情况）
		if p[i] == p[j+1] {
			//j往后移，继续下一个字符的匹配
			j++
		}
		//记录下匹配的最长前后缀到next数组中
		ne[i] = j
	}

	for i, j := 1, 0; i <= m; i++ {
		for j > 0 && s[i] != p[j+1] {
			j = ne[j]
		}
		if s[i] == p[j+1] {
			j++
		}
		if j == n {
			// fmt.Printf("%d " , i - n)
			fmt.Fprintf(writer, "%d ", i-n)
			// 这里是s和p匹配，所以匹配上之后，j在回退继续准备下一轮的匹配
			j = ne[j]
		}
	}
	writer.Flush()
}
