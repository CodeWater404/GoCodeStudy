package main

import (
	"fmt"
	"learngo/object/queue"
)

/**
  @author: CodeWater
  @since: 2023/4/10
  @desc: $
**/

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println("=============================1.支持任何类型的queue=============================")
	q.Push("abc")
	fmt.Println(q.Pop())

}
