package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

/**
  @author: CodeWater
  @since: 2023/4/10
  @desc: 接口
		1.接口的实现是隐式的
		2.只要实现接口里面的方法
		3.接口变量中有实现者的类型 、实现者的值（或者指针，指向实现者）
		4.接口变量自带指针
		5.接口变量同样采用值传递，几乎不需要使用接口的指针
		6.指针接收者实现只能以指针方式使用；值接收者都可
**/

const url = "https://www.baidu.com"

type Retriever interface {
	Get(url string) string //这里函数的实现是在mock/Retriever结构体中
}

type Poster interface {
	Post(url string, form map[string]string) string
}

//
//  RetrieverPoster
//  @Description: 接口的组合
//
type RetrieverPoster interface {
	Retriever
	Poster
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func Post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "code",
		"course": "goland",
	})
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked ibaidu.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T  , %v \n", r, r)
	fmt.Print(" > Type switch:")
	fmt.Printf("type: %T , value: %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("Contents:", v.UserAgent)
	}
}

func main() {
	var r Retriever
	//值接收者这里加&mock.Retriever也可以
	r = mock.Retriever{"this is a fake imooc.com"}
	//fmt.Printf("type：%T ， value：%v\n", r, r) //2.打印类型，值
	inspect(r) //3
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0 (Windows)",
		Timeout:   time.Minute,
	}
	//fmt.Printf("type：%T ， value：%v\n", r, r)
	inspect(r) //3

	//4.type assertion
	//mockRetriever := r.(*real.Retriever)
	//fmt.Println(mockRetriever)

	//5.type assertion and precess error。
	//强制类型转换r.(mock.Retriever)，r转成mock.Retriever类型
	if mockRetriever, ok := r.(mock.Retriever); ok {
		fmt.Println(mockRetriever)
	} else {
		fmt.Println("not a mock retriever")
	}

	//fmt.Println(download(r)) //1

	//6.
	fmt.Println("==========Try a session=============")
	var s RetrieverPoster
	s = &mock.Retriever{"hhhhhhhhh"}
	fmt.Println(session(s))
}
