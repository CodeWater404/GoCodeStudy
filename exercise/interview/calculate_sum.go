package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2024/2/21
  @desc: 两个协程分别计算1-100和101-200的和，然后将两个和相加，最后打印出来
	go1:1-100之和  sum1
	go2:100-200之和 sum2
**/

func main() {
	method2()
}

func method1() {
	var (
		sum1 int = 0
		sum2 int = 0
		sum  int = 0
		wg   sync.WaitGroup
	)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			/*
				单纯sum1 += i 不会影响外面的sum1
			*/
			*(&sum1) += i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 101; i <= 200; i++ {
			*(&sum2) += i
		}
	}()

	wg.Wait() // 一定要用这个让主goroutine等待子goroutine已经执行完毕，否则主goroutine可能会直接退出
	sum = sum1 + sum2
	fmt.Printf("sum1:%v , sum2:%v , sum:%v\n", sum1, sum2, sum)
	// 保持主线程不退出
	//for {
	//}
	//select {} // 会死锁
	/*
		`for {}` 和 `select {}` 都可以用来阻塞主 goroutine，但是它们的工作方式有所不同。

		`for {}` 是一个无限循环，它会一直运行，不断消耗 CPU 资源，但是不会引发死锁。因为它并不依赖于其他 goroutine 的状态，所以即使其他 goroutine 都已经进入休眠状态，`for {}` 也可以继续运行。

		`select {}` 则是在等待一个永远不会发生的事件。`select` 语句用于在多个通道操作中进行选择，但是如果没有提供任何 case，那么它就会永远阻塞。如果你的程序中的其他 goroutine 在完成任务后都已经退出了，那么就没有其他的 goroutine 可以发送或接收数据，导致主 goroutine 无法从阻塞状态中恢复，从而引发了死锁。

		所以，如果你想要在程序的最后阻塞主 goroutine，但是不想引发死锁，你可以使用 `for {}`。但是请注意，`for {}` 会一直消耗 CPU 资源，所以如果你的程序需要长时间运行，那么这可能会导致 CPU 使用率过高。
	*/
}

func dosum1(start, end int, wg *sync.WaitGroup) {
	defer wg.Done()
	localSum := 0
	for i := start; i <= end; i++ {
		localSum += i
	}
	responseChannel <- localSum
}

var responseChannel = make(chan int, 2)

func resp() {
	for rc := range responseChannel {
		sum = append(sum, rc)
	}
}

var sum = make([]int, 0)

func method2() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go dosum1(1, 100, wg)
	go dosum1(101, 200, wg)

	// 等待所以协程执行完毕
	wg.Wait()
	close(responseChannel)
	go resp()
	// sleep让主goroutine等待resp执行完毕
	time.Sleep(time.Second)
	s := sum[0] + sum[1]
	fmt.Printf("sum1:%v , sum2:%v , sum:%v\n", sum[0], sum[1], s)

}
