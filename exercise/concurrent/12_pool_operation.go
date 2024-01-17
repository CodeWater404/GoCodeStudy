package main

import (
	"fmt"
	"sync"
)

/**
  @author: CodeWater
  @since: 2024/1/17
  @desc: sync.Pool池操作
	sync.Pool可以将暂时将不用的对象缓存起来，待下次需要的时候直接使用，不用再次经过
	内存分配，复用对象的内存，减轻GC的压力，提升系统的性能（频鼕地分配、回收内存
	会给GC带来一定的负担，严重的时候会引起CPU的毛刺)
	使用场景：对象池化，TCP连接池、数据库连接池、worker Pool
**/

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	for i := 0; i < 10; i++ {
		//第一次取地0 ，默认用New方法创建。后面get到的都是put进去的
		v := pool.Get()
		fmt.Println(v) //取出来的值都是put进去的，对象复用；如果是新建对象，则取来的是0
		pool.Put(i)
	}
}
