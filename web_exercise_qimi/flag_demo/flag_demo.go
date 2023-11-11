package main

import (
	"flag"
	"fmt"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/11/11
  @desc: 解析命令行参数demo
**/

func main() {
	//定义命令房参数
	var name string
	var age int
	var married bool
	var delay time.Duration
	//name是命令行的参数名，张三是默认值，姓名是说明
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")

	//解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	//返回命令行后的未定义的其他参数
	fmt.Println(flag.Args())
	//返回命令行后的未定义的其他参数个数
	fmt.Println(flag.NArg())
	//返回使用的命令行参数个数
	fmt.Println(flag.NFlag())

}
