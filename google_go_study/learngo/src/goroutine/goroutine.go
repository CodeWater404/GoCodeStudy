package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/4/23
  @desc: goroutine并发编程:用go关键字
	复习： 匿名函数：func(参数列表){}
	1.go run -race goroutine.go：监测数据访问冲突的命令，并发的时候用
	2.任何函数只需加上go就能送给调度器运行
	3.不需要在定义时区分是否是异步函数
	4.调度器在合适的点进行切换。一些可能切换的点如：
		a.I/O,select
		b.函数调用（有时）
		c.channel
		d.runtime.Gosched()
		e.等待锁
		f.PS:只是参考，不能保证切换，不能保证在其他地方不切换
	5.使用-race来检测数据访问冲突
	6.协程会被调度器分配到不同的线程里面去，可能有的线程里面只有1个协程，也有可能有的线程里面有好多个协程在运行
**/

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		//协程coroutine：
		//1.非抢占式，由协程主动交出控制权。（线程是抢占式的）
		//2.编译器/解释器/虚拟机层面的多任务
		//3.多个协程可能在一个或多个线程上运行
		go func(i int) {
			for {
				//一。创建演示
				//fmt.Printf("Hello from goroutine %d\n", i)

				//二。协程手动交出控制权:runtime.Gosched().如果不写会在这个for循环里面一直卡着
				a[i]++
				runtime.Gosched()
			}
		}(i)

	}

	//防止main线程循环执行快导致go开的协程还没运行起来就被杀掉，这里睡眠一下就能看到效果
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
