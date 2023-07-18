package main

import (
	"context"
	"fmt"
	op "go-redis/qimi/init"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/7/18
  @desc:go-redis 还提供了一个执行任意命令或自定义命令的 Do 方法，特别是一些 go-redis 库暂时不支持的命令都可以使用该方法执行。
	具体使用方法如下。
**/

func doDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	//直接执行命令获取错误("EX" 表示下一个参数是过期时间的设置方式,3600秒)
	//SET key2 20 EX 6
	err := op.Rdb.Do(ctx, "set", "key2", 20, "EX", 6).Err()
	fmt.Println("err: ", err)

	//执行命令获取结果
	val, err := op.Rdb.Do(ctx, "get", "key2").Result()
	fmt.Println("do result: ", val, err)
}

func main() {
	doDemo()
}
