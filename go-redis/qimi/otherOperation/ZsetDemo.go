package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	op "go-redis/qimi/init"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/7/19
  @desc: go-redis库操作zset
**/

func zsetDemo() {
	//key
	zsetKey := "Language_Rank"
	//value
	languages := []*redis.Z{
		{Score: 90.0, Member: "golang"},
		{Score: 99.1, Member: "java"},
		{Score: 91, Member: "pythone"},
		{Score: 89, Member: "cpp"},
		{Score: 83, Member: "JavaScript"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	//add zset
	err := op.Rdb.ZAdd(ctx, zsetKey, languages...).Err()
	if err != nil {
		fmt.Println("zadd failed, err is :", err)
		return
	}
	fmt.Println("=====zadd success=====")

	//golang分数加10(ZIncrBy方法如果找不到对应的键值，会新建一个键)
	newScore, err := op.Rdb.ZIncrBy(ctx, zsetKey, 10, "golang").Result()
	if err != nil {
		fmt.Println("zset modift failed , err is : ", err)
		return
	}
	fmt.Println("golang's score is ", newScore, "now!!!!")

	//获取分数最小的
	ret := op.Rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
		//取95---100分的
		oper := &redis.ZRangeBy{
			Min: "95",
			Max: "100",
		}
		ret, err := op.Rdb.ZRangeByScoreWithScores(ctx, zsetKey, oper).Result()
		if err != nil {
			fmt.Println("zrangeByScore failed , err id :", err)
			return
		}
		for _, z := range ret {
			fmt.Println("======get zset with condition success!!!!=====")
			fmt.Println(z.Member, z.Score)
		}
	}
}

func main() {
	zsetDemo()
}
