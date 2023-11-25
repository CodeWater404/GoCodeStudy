package main

import (
	"context"
	"fmt"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/25
  @desc: 用context来取消goroutine

**/

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer func() {
			fmt.Println("goroutine exit...")
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second * 2)
}
