package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/25
  @desc: 用atomic包下的value类型来实现配置变更的场景应用
**/

type Config struct {
	NodeName string
	Addr     string
	Count    int32
}

func loadNewConfig() Config {
	return Config{
		NodeName: "node-1",
		Addr:     "10.10.11.23",
		Count:    rand.Int31(),
	}
}

func main() {
	var config atomic.Value
	config.Store(loadNewConfig())
	var cond = sync.NewCond(&sync.Mutex{})

	// 设置新的config
	go func() {
		for {
			time.Sleep(time.Duration(5+rand.Int63n(5)) * time.Second)
			config.Store(loadNewConfig())
			cond.Broadcast() // 通知等待着的配置已改变
		}
	}()

	go func() {
		for {
			cond.L.Lock()
			cond.Wait()
			c := config.Load().(Config)        // 等待变更信号
			fmt.Printf("new config: %+v\n", c) // 读取新的配置
			cond.L.Unlock()
		}
	}()

	select {}
}
