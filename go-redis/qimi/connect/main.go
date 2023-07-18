package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

/**
  @author: CodeWater
  @since: 2023/7/18
  @desc: redis连接
**/

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 20, //连接池数量
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong, err)
}
