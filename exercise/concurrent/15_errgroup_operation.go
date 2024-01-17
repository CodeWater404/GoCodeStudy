package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
)

/**
  @author: CodeWater
  @since: 2024/1/17
  @desc: errgroup操作
	errgroup可以在一组Goroutine中提供了同步、错误传播以及上下文取消的功能
	使用场景：只要一个goroutine出错我们就不再等其他goroutine了，减少资源浪费，并
	且返回错误
**/

func main() {
	var g errgroup.Group
	var urls = []string{
		"http://www.baidu.com",
		"http://www.google.com",
		"http://www.qq.com",
	}
	for i := range urls {
		url := urls[i]
		g.Go(func() error {
			resp, err := http.Get(url)
			defer resp.Body.Close()
			if err == nil {
				all, _ := io.ReadAll(resp.Body)
				fmt.Println("url:", url, "resp:", string(all))
			}
			return err
		})
	}
	err := g.Wait()
	if err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println("Failed to fetch URL:", err.Error())
	}

}
