package main

import (
	"fmt"
	"reflect"
)

/**
  @author: CodeWater
  @since: 2023/11/27
  @desc: $
**/

func createCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建recv case
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	// 创建send case
	for i, ch := range chs {
		v := reflect.ValueOf(i)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: v,
		})
	}

	return cases
}

func main() {
	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)

	// 创建select case
	var cases = createCases(ch1, ch2)

	// 执行10次select
	for i := 0; i < 10; i++ {
		// 返回三个值：被选中的case的索引，接收到的值，以及一个表示是否成功接收或发送的布尔值。
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() { // recv.IsValid()返回true，则表示这是一个接收操作，否则是一个发送操作。
			fmt.Println("recv: ", cases[chosen].Dir, recv, ok)
		} else { // send case
			fmt.Println("send: ", cases[chosen].Dir, ok)
		}
	}
}
