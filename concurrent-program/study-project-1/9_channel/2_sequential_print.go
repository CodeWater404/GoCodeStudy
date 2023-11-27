package main

import (
	"fmt"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/27
  @desc: 练习题
	有四个goroutine，编号为 1、2、3、4。每秒钟会有一个 goroutine 打印出它自己的编号，
	要求你编写一个程序，让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。
**/

type Token struct{}

func NewWorker(id int, ch chan Token, nextCh chan Token) {
	for {
		token := <-ch         //取的令牌
		fmt.Println((id + 1)) // id从1开始
		time.Sleep(time.Second)
		nextCh <- token // 传递令牌
	}
}

func main() {
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}

	// 创建4个worker
	for i := 0; i < 4; i++ {
		go NewWorker(i, chs[i], chs[(i+1)%4])
	}

	// 首先把令牌给第一个worker
	chs[0] <- struct{}{}

	select {}
}
