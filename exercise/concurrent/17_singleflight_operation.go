package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"sync/atomic"
	"time"
)

/**
  @author: CodeWater
  @since: 2024/1/17
  @desc: singleflight操作
	用于抑制对下游的重复清求
	使用场景：访问缓存、数据库等场景，缓存过期时只有一个请求去更新数据库
**/

var count int32

// 模拟查询数据库
func getArticle(id int) (article string, err error) {
	//假设这里会对数据库进行查询，模拟不同并发下耗时不同
	atomic.AddInt32(&count, 1)
	time.Sleep(time.Duration(count) * time.Millisecond)
	return fmt.Sprintf("article: %d", count), nil
}

// 模拟查询缓存, 有缓存返回缓存，没有缓存返回数据库查询结果,并且只有一个请求读取数据库，其他请求等待
func singleflightGetArticle(sg *singleflight.Group, id int) (string, error) {
	v, err, _ := sg.Do(fmt.Sprintf("%d", id), func() (interface{}, error) {
		return getArticle(id)
	})

	return v.(string), err
}

func main() {
	time.AfterFunc(time.Second*1, func() {
		atomic.AddInt32(&count, -count)
	})

	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 1000
		sg  = &singleflight.Group{}
	)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			res, _ := singleflightGetArticle(sg, 1)
			if res != "article: 1" {
				panic("======================>singleflight err.....")
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("同时并发%d个请求，耗时: %v\n", n, time.Since(now))
}
