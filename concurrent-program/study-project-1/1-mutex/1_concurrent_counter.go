package main

import (
	"fmt"
	"sync"
)

/**
  @author: CodeWater
  @since: 2023/10/13
  @desc: 不使用和使用mutex的并发计算案例：
		开十个协程，每个协程对count累加10万次===》count最后正确结果应该是100万
**/

/*counter：不使用mutex
相关命令：go run -race .\1_concurrent_counter.go
*/
func counter() {
	var count = 0
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				/*count++ 不是一个原子操作，它至少包含几个步骤，比如:
				读取变量count 的当前值，
				对这个值加 1，
				把结果再保存到 count 中。
				因为不是原子操作，就可能有并发的问题。
				Go 提供了一个检测并发访问共享资源是否有问题的工具：race detector，它可以帮助我们自动发现程序有没有 data race 的问题*/
				count++
			}
		}()
	}
	//等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

/*counterByMutex:使用mutex
运行：go run -race .\1_concurrent_counter.go 无警告信息
*/
func counterByMutex() {
	var count = 0
	var wg sync.WaitGroup
	/*这里有一点需要注意：Mutex 的零值是还没有 goroutine 等待的未加锁的状态，所以你不
	需要额外的初始化，直接声明变量（如 var mu sync.Mutex）即可*/
	var mg sync.Mutex
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				//使用mutex加互斥锁保护计数器
				mg.Lock()
				count++
				mg.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func main() {
	//counter()
	counterByMutex()

}
