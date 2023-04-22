package main

import "testing"

/**
  @author: CodeWater
  @since: 2023/4/21
  @desc: 最长连续不重复子串测试类
		PS:注意一个包只能有一个main
**/

//代码覆盖率测试：go tool cover -html=c.out
func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}

}

/**BenchmarkSubstr
** @Description: 性能测试 (当前目录下)：go test -bench .
	性能优化:
		1. go test -bench . -cpuprofile cpu.out(生成二进制的分析文件cpu.out)
		2. go tool pprof cpu.out(交互式查看文件，输入web，通过浏览器查看
** @param b
**/
func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for intput %s , expected %d", actual, s, ans)
		}
	}
}
