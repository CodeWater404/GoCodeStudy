package main

import (
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
)

/**
  @author: CodeWater
  @since: 2024/3/4
  @desc: 控制协程的并发数量
	在使用协程并发处理某些任务时, 其并发数量往往因为各种因素的限制不能无限的增大. 例如网络请求、数据库查询等等。

	从运行效率角度考虑，在相关服务可以负载的前提下（限制最大并发数），尽可能高的并发。

	在Go语言中，可以使用一些方法来控制协程（goroutine）的并发数量，以防止并发过多导致资源耗尽或性能下降。以下是一些常见的方法：
**/

// 1.使用信号量:可以使用 Go 语言中的 channel 来实现简单的信号量机制，限制并发数量。
func semaphoreMethod() {
	concurrency := 3
	sem := make(chan struct{}, concurrency) // 限制并发数量最多3个
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, sem)
		}(i)
	}
	wg.Wait()
	close(sem)
	fmt.Println("semaphoreMethod done...")
}

func worker(id int, sem chan struct{}) {
	sem <- struct{}{} // 占用一个信号量
	defer func() {
		<-sem // 释放一个信号量
	}()

	// do something about work
	fmt.Println("worker ", id, " start...")
}

// 2. 协程池：可以创建一个固定数量的协程池，将任务分发给这些协程执行。
func poolMethod() {
	const numJobs = 5
	const numWorkers = 3
	// jobs 通道用于存储任务，results 通道用于存储处理结果。通过创建固定数量的工作协程，可以有效地控制并发数量。
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup
	// 启动协程池
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			work1(id, jobs, results)
		}(i)
	}

	// 提交任务
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 等待所有任务完成
	go func() {
		wg.Wait()
		close(results)
	}()

	// 处理结果
	for r := range results {
		fmt.Println("result: ", r)
	}
}

func work1(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker ", id, ": working on job ", j)
		results <- j * 2
	}
}

// 3. 使用 golang.org/x/sync/semaphore 包：Go 语言的官方包 golang.org/x/sync/semaphore 提供了一种更加灵活的信号量机制，可以用来控制并发数量。
func semaphoreMethod2() {
	concurrency := 3
	// todo： go get "golang.org/x/sync/semaphore" 没成功。。。
	sem := semaphore.NewWeighted(int64(concurrency)) // 限制并发数量最多3个
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			work2(id, sem)
		}(i)
	}
	wg.Wait()
}

func work2(id int, sem *semaphore.Weighted) {
	//首先调用 sem.Acquire 方法来获取信号量，这将阻塞直到有足够的资源可用。然后，在函数结束时调用 defer sem.Release(1) 来释放信号量资源。
	sem.Acquire(semaphore.WithWeight(1))
	defer sem.Release(1)

	// do something about work
	fmt.Println("worker ", id, " start...")
}
func main() {
	//semaphoreMethod()
	//poolMethod()
	semaphoreMethod2()
}
