package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/21
  @desc: cond例子
	10 个运动员进入赛场之后需要先做拉伸活动活动筋骨，向观众和粉丝
	招手致敬，在自己的赛道上做好准备；等所有的运动员都准备好之后，裁判员才会打响发
	令枪。每个运动员做好准备之后，将 ready 加一，表明自己做好准备了，同时调用 Broadcast 方
	法通知裁判员。因为裁判员只有一个，所以这里可以直接替换成 Signal 方法调用。调用
	Broadcast 方法的时候，我们并没有请求 c.L 锁，只是在更改等待变量的时候才使用到了
	锁。
	裁判员会等待运动员都准备好for ready != 10 。虽然每个运动员准备好之后都唤醒了裁判员，
	但是裁判员被唤醒之后需要检查等待条件是否满足（运动员都准备好了）。可以看到，裁
	判员被唤醒之后一定要检查等待条件，如果条件不满足还是要继续等待。
**/

// Cond 的使用其实没那么简单。它的复杂在于：一，这段代码有时候需要加锁，有
// 时候可以不加；二，Wait 唤醒后需要检查条件；三，条件变量的更改，其实是需
// 要原子操作或者互斥锁保护的。
func main() {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			// 加锁更改等待条件
			c.L.Lock()
			ready++
			c.L.Unlock()

			log.Printf("运动员#%d 已准备就绪\n", i)
			// 广播唤醒所有的等待者
			c.Broadcast()
		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		c.Wait()
		log.Println("裁判员被唤醒一次")
	}
	c.L.Unlock()

	// 所有的运动员都准备就绪，裁判员才开始发令枪
	log.Println("所有运动员都准备就绪，裁判员发令枪！")
}
