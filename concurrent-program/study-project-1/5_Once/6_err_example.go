package main

import (
	"fmt"
	"sync"
)

/**
  @author: CodeWater
  @since: 2023/11/24
  @desc: once错误示例
	想要重置once导致的错误示例
**/

type MyOnce struct {
	sync.RWMutex
	sync.Once
	flag int32
}

func (mo *MyOnce) Reset() {
	mo.Lock()
	defer mo.Unlock()
	mo.Once = sync.Once{}
	mo.flag = 0
}

func (mo *MyOnce) setFlag(i int) int32 {
	mo.RLock()
	if i%2 == 0 {
		//fatal error: sync: unlock of unlocked mutex
		// 因为下面once重置成一个新的空实例了，所以这里do函数里面的锁是未加锁的，unlock才会失败
		defer mo.Do(mo.Reset)
	}
	mo.flag = 1
	mo.RUnlock()
	return mo.flag
}
func main() {
	fmt.Println("==================================")
	mo := new(MyOnce)
	fmt.Println(mo.setFlag(4))
	fmt.Println(mo.setFlag(2))
}
