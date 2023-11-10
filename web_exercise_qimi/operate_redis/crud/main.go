package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/8
  @desc: go-redis基本使用
	https://www.liwenzhou.com/posts/Go/redis/
**/

var rdb *redis.Client

// initClient 初始化
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

// initClientSentry 连接redis哨兵模式
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

// initClientCluster 连接redis集群
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

// redisExample set/get案例
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

// redisExample2 zset案列
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

/*
pipelineDemo
主要是一种网络优化。它本质上意味着客户端缓冲一堆命令并一次性将它们发送到服务器。
这些命令不能保证在事务中执行。这样做的好处是节省了每个命令的网络往返时间（RTT）。
所以前后的命令有依赖的时候，就不能用这个了，因为没有保证事务性;
在某些场景下，当我们有多条命令要执行时，就可以考虑使用pipeline来优化。
*/
func pipelineDemo() {
	pipe := rdb.Pipeline()
	incr := pipe.Incr("pipeline_counter")
	pipe.Expire("pipeline_counter", time.Hour)
	_, err := pipe.Exec()
	if err != nil {
		fmt.Printf("pipeline failed , err:%v\n", err)
		return
	}
	fmt.Printf("pipeline succ , incr:%v\n", incr.Val())
}

// pipelinedDemo 与pipelineDemo类似，只是这个是用的函数
func pipelinedDemo() {
	var incr *redis.IntCmd
	_, err := rdb.Pipelined(func(pipe redis.Pipeliner) error {
		incr = pipe.Incr("pipelined_counter")
		pipe.Expire("pipelined_counter", time.Hour)
		return nil
	})
	if err != nil {
		fmt.Printf("pipelined failed , err:%v\n", err)
		return
	}
	fmt.Printf("pipelined succ , incr:%v\n", incr.Val())
}

/*
txPipelineDemo 事务
Redis是单线程的，因此单个命令始终是原子的，但是来自不同客户端的两个给定命令可以依次执行(也就是在程序运行时，另外一个客户端有命令执行，这两个不冲突不会err)，
watch会报err，
例如在它们之间交替执行。但是，Multi/exec能够确保在multi/exec两个语句之间的命令之间没有其他客户端正在执行命令。
在这种场景我们需要使用TxPipeline。TxPipeline总体上类似于上面的Pipeline，但是它内部会使用MULTI/EXEC包裹排队的命令。
*/
func txPipelineDemo() {
	pipe := rdb.TxPipeline()
	incr := pipe.Incr("tx_pipeline_counter")
	pipe.Expire("tx_pipeline_counter", time.Hour)
	_, err := pipe.Exec()
	if err != nil {
		fmt.Printf("tx_pipeline failed , err:%v\n", err)
		return
	}
	fmt.Printf("tx_pipeline succ , incr:%v\n", incr.Val())
}

// txPipelinedDemo 与txPipelineDemo类似，只是这个是用的函数
func txPipelinedDemo() {
	var incr *redis.IntCmd
	_, err := rdb.TxPipelined(func(pipe redis.Pipeliner) error {
		incr = pipe.Incr("tx_pipelined_counter")
		pipe.Expire("tx_pipelined_counter", time.Hour)
		return nil
	})
	fmt.Printf("tx_pipelined finished , incr:%v , err:%v\n", incr.Val(), err)
}

/*
watchDemo
在某些场景下，我们除了要使用MULTI/EXEC命令外，还需要配合使用WATCH命令。
在用户使用WATCH命令监视某个键之后，直到该用户执行EXEC命令的这段时间里，
如果有其他用户抢先对被监视的键进行了替换、更新、删除等操作，那么当用户尝试执行EXEC的时候
，事务将失败并返回一个错误，用户可以根据这个错误选择重试事务或者放弃事务。
*/
func watchDemo() {
	key := "watch_count"
	err := rdb.Watch(func(tx *redis.Tx) error {
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		_, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
			//具体的业务逻辑
			time.Sleep(time.Second * 6) //在这个期间修改会err
			pipe.Set(key, n+1, 0)
			return nil
		})
		return err
	}, key)
	fmt.Printf("watch finished , err:%v\n", err)
}

func main() {
	if err := initClient(); err != nil {
		fmt.Printf("init redis failed , err:%v\n", err)
	}
	fmt.Printf("init redis success!\n")

	//redisExample()
	//redisExample2()
	//pipelineDemo()
	//pipelinedDemo()
	//txPipelineDemo()
	//txPipelinedDemo()
	watchDemo()
}
