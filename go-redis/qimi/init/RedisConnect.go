package init

import "github.com/go-redis/redis/v8"

/**
  @author: CodeWater
  @since: 2023/7/18
  @desc: redis包级变量初始化连接
**/
var Rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
	PoolSize: 20,
})
