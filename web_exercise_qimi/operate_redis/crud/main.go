package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

/**
  @author: CodeWater
  @since: 2023/11/8
  @desc: go-redis基本使用
**/

var rdb *redis.Client

//initClient 初始化
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

//initClientSentry 连接redis哨兵模式
func initClientSentry() (err error) {
	rdb = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

//initClientCluster 连接redis集群
func initClientCluster() (err error) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

//redisExample set/get案例
func redisExample() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed , err:%v\n", err)
		return
	}
	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed , err:%v\n", err)
		return
	}

	fmt.Printf("score:%v\n", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Printf("name does not exist..")
	} else if err != nil {
		fmt.Printf("get name failed , err:%v\n", err)
		return
	} else {
		fmt.Printf("name:%s\n", val2)
	}
}

//redisExample2 zset案列
func redisExample2() {
	zsetKey := "lang_rank"
	languages := []redis.Z{
		redis.Z{Score: 90.0, Member: "golang"},
		redis.Z{Score: 98.0, Member: "java"},
		redis.Z{Score: 95.0, Member: "python"},
		redis.Z{Score: 97.0, Member: "javaScript"},
		redis.Z{Score: 99.0, Member: "c/cpp"},
	}
	num, err := rdb.ZAdd(zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed , err:%v\n", err)
		return
	}
	fmt.Printf("zadd %d succ\n", num)

	newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed , err:%v\n", err)
		return
	}
	fmt.Printf("golang's score is %f now \n", newScore)

	//取分数最高的3个
	ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed , err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Printf("===>%s , %f\n", z.Member, z.Score)
	}

	//取95-100的
	op := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed , err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Printf("===>2 %s , %f\n", z.Member, z.Score)
	}

}

func main() {
	if err := initClient(); err != nil {
		fmt.Printf("init redis failed , err:%v\n", err)
	}
	fmt.Printf("init redis success!\n")

	//redisExample()
	redisExample2()
}
