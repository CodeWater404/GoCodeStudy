package main

import (
	"fmt"
	"io/ioutil"
)

/**switch:
1. 不用每次写break
2。 switch后可以没有表达式
*/
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintln("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"

	}
	return g
}

func main() {
	const filename = "abc.txt"
	//go可以返回两个变量
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("%s\n", contents)
	//}

	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println("can not print file contents: ", err)
	} else {
		fmt.Println(string(contents))
	}

	fmt.Println(
		grade(59),
		grade(78),
		grade(97),
		grade(100),
		grade(-1),
	)

}
