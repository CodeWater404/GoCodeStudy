package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
  @author: CodeWater
  @since: 2023/7/5
  @desc: 懒汉式
**/

type lazy struct{}

var instance_lazy *lazy

//最普通的饿汉式：但是并发情况下有性能安全问题
func GetInstance_lazy() *lazy {
	//首次调用的时候才会生成对象
	if instance_lazy == nil {
		instance_lazy = new(lazy)
		return instance_lazy
	}
	return instance_lazy
}

var lock sync.Mutex

//为了性能安全，加一个锁
func GetInstance_lazy2() *lazy {
	lock.Lock()
	defer lock.Unlock()

	if instance_lazy == nil {
		instance_lazy = new(lazy)

	}
	return instance_lazy
}

//GetInstance_lazy2虽然性能安全了，但是每次调用的时候都会加锁导致慢.这里可以加个标志位来优化一下
var initialized uint32 //默认值0

func GetInstance_lazy3() *lazy {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance_lazy
	}
	lock.Lock()
	defer lock.Unlock()

	if instance_lazy == nil {
		instance_lazy = new(lazy)
		atomic.StoreUint32(&initialized, 1)
	}
	return instance_lazy
}

//go语言提供了相同的实现，下面是对GetInstance_lazy3的优化
var once sync.Once

func GetInstance_lazy4() *lazy {
	once.Do(func() {
		instance_lazy = new(lazy)
	})
	return instance_lazy
}

func main() {
	s1 := GetInstance_lazy()
	s2 := GetInstance_lazy2()
	s3 := GetInstance_lazy3()
	s4 := GetInstance_lazy4()
	if s1 == s2 {
		fmt.Printf("s1==s2 , address value: %v\n", s1)
	}
	if s3 == s4 {
		fmt.Printf("s3==s4 , address value: %v\n", s3)
	}
}
