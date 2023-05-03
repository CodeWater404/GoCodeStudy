package main

import (
	"fmt"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/5/2
  @desc: channel通道
	1. 往channel中发送数据的话，一定要有人来接收数据

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
func worker(c chan int) {
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
		fmt.Printf("worker %d receive %c\n", id, <-c)
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

func channelClose() {
	c := make(chan int)
	go worker2(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	//close() is used only by the sender
	close(c)
	time.Sleep(time.Millisecond)

}

/**createWorker
** @Description: channel as a return value
	1. 从channel中发送数据写法： chan<-
	2. 从channel中收数据写法： <-chan
** @param id
** @return chan
**/
func createWorker(id int) chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %c \n", id, <-c)
		}
	}()
	return c
}

func channelArray() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		//1,
		//channels[i] = make(chan int)
		//2.
		channels[i] = createWorker(i)
		go worker1(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	//缓冲区大小为3
	c := make(chan int, 3)
	go worker1(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'

	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()

	fmt.Println("=============================chan array=============================")
	channelArray()

	fmt.Println("=============================buffered channel=============================")
	bufferedChannel()

	fmt.Println("=============================channel close=============================")
	channelClose()
}
