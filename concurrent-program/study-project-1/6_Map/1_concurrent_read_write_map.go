package main

import (
	"fmt"
	"sync"
)

/**
  @author: CodeWater
  @since: 2023/11/24
  @desc: 并发读写map导致panic
**/

func main() {
	var m = make(map[int]int, 10)
	go func() {
		for {
			m[1] = 1 // 设置key
		}
	}()

	go func() {
		for {
			_ = m[2] // 读取key
		}
	}()

	fmt.Println("====1===")
	select {}
	sync.Map{}

}
