package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/5/4
  @desc: use select to schedule(使用select来调度)
	1. use nil channel in select
**/

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func test() {
	var c1, c2 = generator(), generator()
	w := createWorker(0)
	for {
		//whoever sends the data received it first
		select {
		case n := <-c1:
			//fmt.Println("Received from c1:", n)
			w <- n
		case n := <-c2:
			//fmt.Println("Received from c2:", n)
			w <- n
		}
	}
}

/**test2
** @Description:
	shortcomings:
	if the data is produced too quickly , then n will only save the latest data, and some data will be lost before it can be consumed
**/
func test2() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	n := 0
	hasValue := false
	for {
		var activWorker chan<- int
		if hasValue {
			activWorker = worker
		}

		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activWorker <- n:
			hasValue = false
		}
	}
}

/**test3
** @Description: fix bugs in test2
	exit the program after a period of time
**/
func test3() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		//the time spent on each select
		//that is , print timeout every time the time difference between generating data exceeds 800 milliseconds(也就是每两次生成数据之间的时间差超过800毫秒就打印timeout)
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			//periodically output the backlog data
			fmt.Println("queue len= ", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}

func main() {
	//fmt.Println("=============================write method one=============================")
	//test()

	//fmt.Println("=============================write method two=============================")
	//test2()

	fmt.Println("=============================write method three=============================")
	test3()

}
