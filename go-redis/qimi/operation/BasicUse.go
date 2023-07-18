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
  @desc: go-redis基本操作
**/

func doCommand() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	//执行命令获取结果
	val, err := op.Rdb.Get(ctx, "key").Result()
	fmt.Println("===1. ", val, " , ", err)

	//先获取到命令对象再执行命令
	cmder := op.Rdb.Get(ctx, "key")
	fmt.Println("====2. ", cmder.Val())
	fmt.Println(cmder.Err())

	//直接执行命令获取错误
	err = op.Rdb.Set(ctx, "key", 10, time.Hour).Err()
	fmt.Println("====3. ", err)

	//直接执行命令获取数据
	value := op.Rdb.Get(ctx, "key").Val()
	fmt.Println("====4. ", value)
}

func main() {
	doCommand()
}
