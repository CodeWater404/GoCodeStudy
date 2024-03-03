package main

import "fmt"
import "time"

/**
  @author: CodeWater
  @since: 2024/2/19
  @desc: $
**/

func Getconf() *Baseconfig {
	return &conf
}

var conf Baseconfig

type Baseconfig struct {
	currentLoad uint64
}

func main() {
	go func() {
		for {
			Getconf().currentLoad = 1
		}
	}()
	go func() {
		for {
			fmt.Println(Getconf().currentLoad)
			time.Sleep(1 * time.Second)
			//todo：后面记录为一个博客
			// 如果不用sleep，go run -race 的结果会被冲掉所以看不出来(错误的说法，其实是sleep之后加大了出现程序竞争的概率，所以才能看到数据竞争的结果，
			// 因为数据竞争是在运行时检测的，而不是编译时)
			// chatgpt：，数据竞争并不总是出现，它可能取决于许多因素，包括操作系统、硬件、Go 版本等。有时候，在某些情况下，数据竞争可能不会被 go run -race 检测到。
		}
	}()
	select {}
}
