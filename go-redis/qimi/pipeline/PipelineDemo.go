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
  @since: 2023/7/23
  @desc:	Redis Pipeline 允许通过使用单个 client-server-client 往返执行多个命令来提高性能。
	区别于一个接一个地执行100个命令，你可以将这些命令放入 pipeline 中，然后使用1次读写操作像执行单个命令一样执行它们。
	这样做的好处是节省了执行命令的网络往返时间（RTT）。

	在那些我们需要一次性执行多个命令的场景下，就可以考虑使用 pipeline 来优化。
**/

var ctx = context.Background()

/*下面的代码相当于将以下两个命令一次发给 Redis Server 端执行，与不使用 Pipeline 相比能减少一次RTT。
INCR pipeline_counter
EXPIRE pipeline_counts 3600
*/
func pipelineDemo() {

	//1。
	//pipe := op.Rdb.Pipeline()
	//incr := pipe.Incr(ctx, "pipeline_counter")
	//pipe.Expire(ctx, "pipeline_counter", time.Hour)
	//cmds, err := pipe.Exec(ctx)

	//2.或者，你也可以使用Pipelined 方法，它会在函数退出时调用 Exec。
	var incr *redis.IntCmd

	cmds, err := op.Rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		incr = pipe.Incr(ctx, "pipelined_counter")
		pipe.Expire(ctx, "pipeline_counter", time.Hour)
		return nil
	})

	if err != nil {
		panic(err)
	}

	for k, v := range cmds {
		fmt.Println("cmds k:", k, " , v:", v)
	}

	// 在执行pipe.Exec之后才能获取到结果
	fmt.Println(incr.Val())
}

/*
下方的示例代码中使用pipiline一次执行了100个 Get 命令，在pipeline 执行后遍历取出100个命令的执行结果。
*/
func pipelineDemo2() {
	cmds, err := op.Rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 100; i++ {
			pipe.Get(ctx, fmt.Sprintf("key%d", i))
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, cmd := range cmds {
		fmt.Println(cmd.(*redis.StringCmd).Val())
	}
}

func main() {
	//pipelineDemo()

	pipelineDemo2()
}
