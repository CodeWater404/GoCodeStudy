package main

import (
	"context"
	"fmt"
	op "go-redis/qimi/init"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/7/21
  @desc: 遍历所有的key
	你可以使用KEYS prefix:* 命令按前缀获取所有 key。
	但是如果需要扫描数百万的 key ，那速度就会比较慢。这种场景下你可以使用Scan 命令来遍历所有符合要求的 key。
**/

func scanKeysDemo1() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var cursor uint64
	for {
		var keys []string
		var err error
		//按前缀扫描key：match匹配模式，count没懂具体是怎么作用。。。(可以使用KEYS prefix:* 命令按前缀获取所有 key)
		keys, cursor, err = op.Rdb.Scan(ctx, cursor, "t*", 1).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key:", key)
		}

		//没有key了
		if cursor == 0 {
			break
		}
	}
}

func scanKeysDemo2() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	//按前缀扫描key
	iter := op.Rdb.Scan(ctx, 0, "*", 5).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys: ", iter.Val())
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}

}

func delKeyByMatch(match string, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	iter := op.Rdb.Scan(ctx, 0, match, 0).Iterator()
	for iter.Next(ctx) {
		err := op.Rdb.Del(ctx, iter.Val()).Err()
		if err != nil {
			panic(err)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

func main() {
	scanKeysDemo1()
	fmt.Println("==========================================================")
	scanKeysDemo2()
	fmt.Println("==========================================================")
	delKeyByMatch("t*", 500*time.Millisecond)
}
