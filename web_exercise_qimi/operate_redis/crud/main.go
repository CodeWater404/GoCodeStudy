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

func main() {
	if err := initClient(); err != nil {
		fmt.Printf("init redis failed , err:%v\n", err)
	}
	fmt.Printf("init redis success!")
}
