package main

import (
	"fmt"
	"sync"
)

/**
  @author: CodeWater
  @since: 2024/1/17
  @desc: map并发操作
	线程安全的
	适用场景：map并发读写
**/

func main() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
