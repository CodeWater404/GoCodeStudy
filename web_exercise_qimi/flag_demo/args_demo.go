package main

import (
	"fmt"
	"os"
)

/**
  @author: CodeWater
  @since: 2023/11/11
  @desc: 获取命令行参数
**/

func main() {
	fmt.Printf("os.Args=%v\n", os.Args)
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}
