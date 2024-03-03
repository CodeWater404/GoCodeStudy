package main

import (
	"fmt"
	"sync"
)

/**
  @author: CodeWater
  @since: 2024/3/1
  @desc:
	使用Go语言编写一个程序，实现两个goroutine之间使用channel进行通信，要求一个goroutine生成1到100的整数，另一个goroutine接收并计算这些整数的平方和。
**/

func main() {
	//method1()
	//method2()
	method3()

}

// 一个通道+wg
func method1() {
	num := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			num <- i
		}
		// 一定要关闭，否则会导致下面的goroutine死锁，因为range会一直等待无法结束循环，直到channel关闭。
		// 确保第二个 goroutine 在读取完所有数据后可以正常退出循环，避免了死锁问题。
		close(num)
	}()
	sum := 0

	go func() {
		defer wg.Done()
		for x := range num {
			*(&sum) += x * x
		}
	}()
	wg.Wait()
	fmt.Println(sum)
}

// 两个通道（一个通道缓冲100）+wg
func method2() {
	num := make(chan int, 100)
	sumChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			num <- i
		}
		close(num)
	}()

	go func() {
		defer wg.Done()
		localSum := 0
		for x := range num {
			localSum += x * x
		}
		sumChan <- localSum
	}()
	sum := <-sumChan //必须要在wait之前消费掉，因为wait会等待所有goroutine执行完毕，如果不消费掉第二个goroutine会一直阻塞在最后发数据那，
	//因为sumchan无缓冲 ，会导致死锁
	wg.Wait()
	fmt.Println("method2:", sum)
}

// 两个通道（一个通道缓冲只有1）+wg
func method3() {
	num := make(chan int, 1)
	sumChan := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			num <- i
		}
		close(num)
	}()

	go func() {
		defer wg.Done()
		localSum := 0
		for x := range num {
			localSum += x * x
		}
		sumChan <- localSum
	}()
	sum := <-sumChan
	wg.Wait()
	fmt.Println(sum)
}

//ot
