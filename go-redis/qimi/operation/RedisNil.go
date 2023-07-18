package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	op "go-redis/qimi/init"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/7/18
  @desc:
	go-redis 库提供了一个 redis.Nil 错误来表示 Key 不存在的错误。
	因此在使用 go-redis 时需要注意对返回错误的判断。在某些场景下我们应该区别处理 redis.Nil 和其他不为 nil 的错误。
**/

func getValueFromRedis(key, defaultValue string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	val, err := op.Rdb.Get(ctx, key).Result()
	if err != nil {
		//如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			return defaultValue, nil
		}
		//出现其他错误
		return "other unknown error!!!!", err
	}
	return val, err
}

func main() {
	str, err := getValueFromRedis("key", "get existed value")
	fmt.Println(str, err)

	fmt.Println("==========================================================")

	str2, err := getValueFromRedis("hhh", "get none value")
	fmt.Println(str2, err)
}
