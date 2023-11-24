package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

/**
  @author: CodeWater
  @since: 2023/11/24
  @desc: 扩展once，如果执行过once返回true，否则返回false
**/

// Once 功能更强大的once
type Once struct {
	sync.Once
}

// Done 判断once是否执行过
func (o *Once) Done() bool {
	return atomic.LoadUint32((*uint32)(unsafe.Pointer(&o.Once))) == 1
}

func main() {
	var flag Once
	fmt.Println("===>1 : ", flag.Done()) // false

	flag.Do(func() {
		time.Sleep(time.Second)
	})

	fmt.Println("===>2 : ", flag.Done()) // true
}
