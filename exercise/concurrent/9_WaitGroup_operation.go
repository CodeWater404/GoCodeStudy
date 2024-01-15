package main

import "sync"

/**
  @author: CodeWater
  @since: 2024/1/15
  @desc: WaitGroup基本操作
	sync.WaitGroup可以等待一组Goroutine的返回
	使用场景：并发等待，任务编排，一个比较常见的使用场景是批量发出RPC或者HTTP请求
**/

func main() {
	requests := []*Request{}
	wg := &sync.WaitGroup{}
	wg.Add(len(requests))

	for _, request := range requests {
		go func(r *Request) {
			defer wg.Done()
			//res , err := service.call(r)
		}(request)
	}
	wg.Wait()
}
