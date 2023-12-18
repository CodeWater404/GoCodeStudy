package main

import (
	"fmt"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/12/18
  @desc: trace的编程过程:这个方式是通过命令行来查看gmp的信息
	go build trace2.go
	// mac下的执行
	GODEBUG=schedtrace=1000 ./trace2.exe   	===> 会显示gmp的信息
	// win下的执行, 1000毫秒
	$env:GODEBUG="schedtrace=1000" ; .\trace2.exe
	// 或者也可以：set GODEBUG=schedtrace=1000 ，然后  .\trace2.exe

**/

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello gmp")
	}

}
