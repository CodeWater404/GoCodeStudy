package main

import (
	"sync"
	"sync/atomic"
)

/**
  @author: CodeWater
  @since: 2023/11/24
  @desc: 自己实现once，do方法正确完成不再执行，错误时可以再次执行
**/

// Once 功能更强大的once
type Once struct {
	m    sync.Mutex
	done uint32
}

// Do 执行函数
// 传入的函数f有返回值error，如果初始化失败，需要返回失败的error
// Do方法会把这个error返回给调用者
func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.slowDo(f)
}

// slowDo  如果还没有初始化
func (o *Once) slowDo(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 { // 双重检查
		err = f()
		if err == nil { // 初始化成功,才设置done为1
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}

func main() {

}
