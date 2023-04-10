package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/4/8
  @desc: 最长不重复子串
**/

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		//看map中是否已经出现过字符，如果出现过说明该字符重复；并且上一次出现的位置要大于当前start，这样才可以去更新start，重新找一段不重复的子串
		if lastI, ok := lastOccured[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		//i与start所在的长度是否超过之前记录的最大长度，超过则更新
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		//记录每个字符所在的位置
		lastOccured[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcdads"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("哈哈哈"))
	fmt.Println(lengthOfNonRepeatingSubStr("大厦和客户端"))
}
