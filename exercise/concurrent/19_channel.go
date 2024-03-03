package main

import "fmt"

/**
  @author: CodeWater
  @since: 2024/3/2
  @desc: 当通过通道发送有限的数据时，我们可以通过close函数关闭通道来告知从该通道接收值的goroutine停止等待。
	当通道被关闭时，往该通道发送值会引发panic，从该通道里接收的值一直都是类型零值。

	通道关闭的原则：
	In Go, it is not always necessary to close a channel. Whether or not to close a channel depends on the specific use case and the design of the program. Here are some guidelines for when to close a channel in Go:

	Closing for Signaling:

	Close a channel to signal that no more values will be sent. This is useful when the receiver needs to be notified that all values have been sent and it should terminate its operations, such as in a range loop.
	Use of range:

	If the receiver is using a range loop to iterate over the values sent on a channel, it's important to close the channel when all the values have been sent. This allows the receiver to exit the loop when the channel is closed.
	Resource Cleanup:

	If the sender knows that the receiver is no longer interested in receiving values from the channel, it can close the channel to free up resources and allow for clean shutdown.
	Avoiding Deadlocks:

	In some cases, leaving a channel open when it's no longer needed can lead to deadlocks, especially in scenarios where the receiver is waiting indefinitely for more values.
	Multiple Senders:

	If a channel has multiple concurrent senders, it's generally not necessary to close the channel, as it can be challenging to coordinate the closing of the channel among multiple senders.
	It's important to note that Go is garbage collected, so leaving a channel open without closing it will not lead to memory leaks. The channel will be eventually garbage collected once it's no longer in use.

	In summary, while it's not always mandatory to close a channel in Go, there are specific scenarios, such as signaling completion or resource cleanup, where closing a channel is beneficial for the overall design and behavior of the program.


**/

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环!!!!!所以通道用完之后一定要关闭，否则会引起死锁
		fmt.Println(i)
	}
}
