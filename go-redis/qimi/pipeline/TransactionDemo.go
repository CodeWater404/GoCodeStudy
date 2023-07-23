package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	op "go-redis/qimi/init"
	"log"
	"sync"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/7/24
  @desc: 事务
	Redis 是单线程执行命令的，因此单个命令始终是原子的，但是来自不同客户端的两个给定命令可以依次执行，例如在它们之间交替执行。
	但是，Multi/exec能够确保在multi/exec两个语句之间的命令之间没有其他客户端正在执行命令。
**/

//以包级运行的时候，可以注释掉这行
var ctx = context.Background()

/*
下面代码相当于在一个RTT下执行了下面的redis命令：
MULTI
INCR pipeline_counter
EXPIRE pipeline_counts 3600
EXEC
*/
func transaction() {
	//TxPipeline
	pipe := op.Rdb.TxPipeline()
	incr := pipe.Incr(ctx, "tx_pipeline_counter")
	pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
	_, err := pipe.Exec(ctx)
	fmt.Println(incr.Val(), err)

	fmt.Println("==========================================================")

	//TxPipelined
	var incr2 *redis.IntCmd
	_, err = op.Rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		incr2 = pipe.Incr(ctx, "tx_pipelined_counter")
		pipe.Expire(ctx, "tx_pipelined_counter", time.Hour)
		return nil
	})
	fmt.Println(incr2.Val(), err)

}

/*
我们通常搭配 WATCH命令来执行事务操作。从使用WATCH命令监视某个 key 开始，直到执行EXEC命令的这段时间里，
如果有其他用户抢先对被监视的 key 进行了替换、更新、删除等操作，那么当用户尝试执行EXEC的时候，事务将失败并返回一个错误，
用户可以根据这个错误选择重试事务或者放弃事务。
*/
func watchDemo(ctx context.Context, key string) error {
	return op.Rdb.Watch(ctx, func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 假设操作耗时10秒
		// 10秒内我们通过其他的客户端修改key，当前事务就会失败;如果不动，当前key就会自动加1
		time.Sleep(10 * time.Second)
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n+1, time.Hour)
			return nil
		})
		return err
	}, key)

}

// go-redis 官方文档中使用 GET 、SET和WATCH命令实现一个 INCR 命令的完整示例。
const routineCount = 100

func watchDemo2() {
	increment := func(key string) error {
		txf := func(tx *redis.Tx) error {
			n, err := tx.Get(ctx, key).Int()
			if err != nil && err != redis.Nil {
				//return err
				log.Fatalf("watchdemo2 increment txf get error: %v", err)
				return err
			}
			n++

			_, err = tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
				pipe.Set(ctx, key, n, 0)
				log.Fatalf("watchdemo2 increment txf pipelined error: %v", err)
				return nil
			})
			log.Printf("watchdemo2 increment txf finished: %v", err)
			return err
		}
		for retries := routineCount; retries > 0; retries-- {
			err := op.Rdb.Watch(ctx, txf, key)
			if err != redis.TxFailedErr {
				return err
			}
		}
		return errors.New("increment reached maximum number of retries...")
	}
	var wg sync.WaitGroup
	wg.Add(routineCount)

	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()
			if err := increment("counter0"); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := op.Rdb.Get(ctx, "counter0").Int()
	fmt.Println("ended with , key: ", n, " , val: ", err)
}

func main() {
	//transaction()

	//err := watchDemo(ctx, "key0")
	//if err != nil {
	//	fmt.Println("watch err:", err)
	//}

	watchDemo2()
}
