package main

import (
	"fmt"
	"reflect"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/27
  @desc: 任务编排——or-done channel
**/

// orRecursion 递归实现or-done
func orRecursion(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有零个或者一个channel
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2: // 两个channel的情况也是特殊
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default: // 超过2个，二分递归处理
			m := len(channels) / 2
			select {
			case <-orRecursion(channels[:m]...):
			case <-orRecursion(channels[m:]...):

			}
		}
	}()
	return orDone
}

// orReflect 反射实现or-done
func orReflect(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有零个或者一个channel
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		// 利用反射构建selectcase
		var cases []reflect.SelectCase
		for _, ch := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			})
		}
		// 随机选择一个可用的case
		reflect.Select(cases)
	}()
	return orDone
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()

	<-orRecursion(
		sig(10*time.Second),
		sig(20*time.Second),
		sig(30*time.Second),
		sig(40*time.Second),
		sig(50*time.Second),
		sig(01*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}
