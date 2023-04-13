package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/**
  @author: CodeWater
  @since: 2023/4/13
  @desc: 函数式编程
**/

/**fibonacci
** @Description: 斐波那契数列。1 ， 1， 2 ， 3 ， 5 ， 8.。。
** @return func() int
**/
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	//todo：p太小，如果数字太大放不下会有问题
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

//读取文件
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := fibonacci()
	fmt.Printf("f: %T\n", f)
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("=============================用printFileContents读取fib=============================")
	printFileContents(f)
}
