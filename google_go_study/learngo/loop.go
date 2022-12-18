package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//转二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result //每次取模就是高位的，所以直接放在前面
	}
	return result
}

//用for一行一行的读取文件
func printFile(filename string) {
	file, err := os.Open(filename)
	//err不等于空，也就是有报错信息，
	if err != nil {
		panic(err)
	}

	//scanner一行一行的读
	scanner := bufio.NewScanner(file)
	//go没有while ， 这种就相当于while了
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}

//死循环
func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	//fmt.Println(
	//	convertToBin(5),
	//	convertToBin(13),
	//	convertToBin(543645),
	//)

	//printFile("abc.txt")

	forever()
}
