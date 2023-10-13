package main

import (
	"fmt"
	"sync"
)

/**
  @author: CodeWater
  @since: 2023/10/14
  @desc: mutex的其他使用场景
		1. 嵌入到其它 struct 中使用
		2. 采用嵌入字段的方式
**/

/*Counter1
嵌入到其它 struct 中使用：在初始化嵌入的 struct 时，也不必初始化这个 Mutex 字段，不会因为没有初始化出现空指针或者是无法获取到锁的情况.
*/
type Counter1 struct {
	mu    sync.Mutex
	Count uint64
}

/*Counter2
采用嵌入字段的方式：通过嵌入字段，你可以在这个 struct 上直接调用 Lock/Unlock 方法.如果嵌入的 struct 有多个字段，我们一般会把 Mutex 放在要控制的字段上面，然后使用
空格把字段分隔开来。即使你不这样做，代码也可以正常编译，只不过，用这种风格去写的话，逻辑会更清晰，也更易于维护。
*/
type Counter2 struct {
	sync.Mutex
	Count uint64
}

func (c2 *Counter2) calculator2() {
	var counter2 Counter2
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter2.Lock()
				counter2.Count++
				counter2.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter2.Count)
}

/*Counter3 线程安全的计数器类型
counter2进化版，把获取锁、释放锁、计数加一的逻辑封装成一个方法，对外不需要暴露锁等逻辑。
*/
type Counter3 struct {
	CounterType int
	Name        string

	mu    sync.Mutex
	count uint64
}

/*Incr
加1的方法，内部使用互斥锁保护
*/
func (c3 *Counter3) Incr() {
	c3.mu.Lock()
	c3.count++
	c3.mu.Unlock()
}

/*GetCount3
得到计数器的值，也需要锁保护
*/
func (c3 *Counter3) GetCount3() uint64 {
	c3.mu.Lock()
	defer c3.mu.Unlock()
	return c3.count
}

func (c3 *Counter3) calculator3() {
	//封装好的计数器
	var counter3 Counter3
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				//收到锁保护的方法
				counter3.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter3.GetCount3())
}

func main() {
	//counter2 := &Counter2{Count: 0}
	//counter2.calculator2()

	counter3 := &Counter3{
		count: 0,
	}
	counter3.calculator3()
}
