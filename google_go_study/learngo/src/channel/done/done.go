package main

import (
	"fmt"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/5/2
  @desc: channel通道
	使用Channel来等待goroutine结束
	以及WaitGroup的使用

**/

func channelDemo() {
	//1. create a channel
	//var c chan int

	//2, create a channel
	c := make(chan int)

	//3.1create a gorouting :receive data from channel
	//a channel is the communication coroutines and coroutines
	//go func() {
	//	for {
	//		n := <-c
	//		fmt.Println(n)
	//	}
	//}()

	//3.2create a gorouting :receive data from channel
	//go worker(c)

	//3.3 another way to writer it
	go worker1(0, c)

	//send data to channel
	c <- 1
	c <- 2
	//add a time to delay: prevent 2 from printing too late
	time.Sleep(time.Millisecond)
}

/**worker
** @Description: channel as a parameter
** @param c
**/
func workered(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}

/**worker1
** @Description: channel as a parameter
** @param id
** @param c
**/
func worker1(id int, c chan int) {
	for {
		fmt.Printf("worker %d received %d\n", id, <-c)
	}
}

func worker2(id int, c chan int) {
	//1.
	//for {
	////ok check whether the channel is finished
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker2 %d received %d\n", id, n)
	//}

	//2. another way to write it
	//use range to automatically determine the end of sending data
	for n := range c {
		fmt.Printf("worker2 %d received %d\n", id, n)
	}
}

func worker3(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
	}
}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c \n", id, n)
		w.done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

/**createWorker
** @Description: channel as a return value
	1. 从channel中发送数据写法： chan<-
	2. 从channel中收数据写法： <-chan
** @param id
** @return chan
**/
func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w)
	return w
}

func channelArray() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}

	for _, worker := range workers {
		<-worker.done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	//wait for all of them
	for _, worker := range workers {
		<-worker.done
	}

}

type workerType struct {
	in chan int
	//1.
	//wg *sync.WaitGroup

	//2. another way to write it
	done func()
}

//func doWork2(id int, c chan int, wg *sync.WaitGroup) {
func doWork2(id int, w workerType) {
	//for n := range c {
	for n := range w.in {
		fmt.Printf("worker %d received %c \n", id, n)

		//1.end of mark
		//wg.Done()
		//2.another way to write it
		w.done()
	}
}

func createWorker2(id int, wg *sync.WaitGroup) workerType {
	w := workerType{
		in: make(chan int),
		//1.
		//wg: wg,
		//2.another way to write it
		done: func() {
			wg.Done()
		},
	}
	//1.
	//go doWork2(id, w.in, wg)
	//2.another way to write it
	go doWork2(id, w)
	return w
}

func channelArray2() {
	var wg sync.WaitGroup
	var workers [10]workerType

	for i := 0; i < 10; i++ {
		workers[i] = createWorker2(i, &wg)
	}
	//add 20 tasks
	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	//wait for the coroutine to end
	wg.Wait()
}

func main() {
	channelArray()

	fmt.Println("=============================use built-in library=============================")
	channelArray2()
}
