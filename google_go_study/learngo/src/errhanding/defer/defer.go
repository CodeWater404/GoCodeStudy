package main

import (
	"bufio"
	"errors"
	"fmt"
	"learngo/functional/fib"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/4/14
  @desc: 错误处理和资源管理
	defer调用：
		1.确保调用在函数结束时发生
		2.参数在defer语句时计算
		3.defer列表为后进先出
	何时使用deferi调用：
		1.Open/Close
		2.Lock/Unlock
		3.PrintHeader/PrintFooter
**/

func tryDefer() {
	//加了defer，会在函数结束后执行
	defer fmt.Println(1)
	defer fmt.Println(2)
	//defer使用的是栈，先进后出
	fmt.Println(3)

	//有了defer。函数内部发生错误或者中间return，都可以执行
	panic("error occurred")
	fmt.Println(4)

}

/**tryDefer2
** @Description: 2.参数在defer语句时计算.会打印从30到1的，不会打印30行30
**/
func tryDefer2() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many.....")
		}
	}
}

func writeFile(filename string) {
	//file, err := os.Create(filename)
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	//还可以自定义错误
	err = errors.New("this is a custom error!")

	//错误处理
	if err != nil {
		//已知错误
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else { //未知错误
			fmt.Printf("%s , %s , %s \n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer() //1
	//tryDefer2()//2
	writeFile("fib.txt")
}
