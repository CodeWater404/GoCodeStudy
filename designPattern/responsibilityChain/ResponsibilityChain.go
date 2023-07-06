package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/6
  @desc: 责任链
	定义：将请求的发送和接收解耦，让多个接收对象都有机会处理这个请求。将这些接
		收对象串成一条链，并沿着这条链传递这个请求，直到链上的某个接收对象能
		够处理它为止。
	场景：过滤器、中间件
	实现方式： 链表、数组

	例子：有一个网站人们可以在上面随意评论，作为开发者需要对一些广告、辱骂、政治敏感词进行过滤。
**/

//敏感词过滤器：判断是否时敏感词
type SensitiveWordFilter interface {
	Filter(content string) bool
}

//职责链
type SensitiveWordFilterChain struct {
	filters []SensitiveWordFilter
}

//添加一个过滤器
func (c *SensitiveWordFilterChain) AddFilter(filter SensitiveWordFilter) {
	c.filters = append(c.filters, filter)
}

//执行过滤
func (c *SensitiveWordFilterChain) Filter(content string) bool {
	for _, filter := range c.filters {
		if filter.Filter(content) {
			return true
		}
	}
	return false
}

//广告敏感词
type AdSensitiveWordFilter struct{}

//filter广告过滤算法
func (f *AdSensitiveWordFilter) Filter(content string) bool {
	//具体的实现逻辑
	return false
}

//政治敏感词
type PoliticalWordFilter struct{}

//政治敏感过滤算法
func (f *PoliticalWordFilter) Filter(content string) bool {
	//具体的实现逻辑
	return true
}

//主体业务逻辑
func main() {
	chain := &SensitiveWordFilterChain{}
	chain.AddFilter(&AdSensitiveWordFilter{})
	if chain.Filter("test...") == true {
		fmt.Println("广告敏感词过滤成功")
	} else {
		fmt.Println("广告敏感词过滤失败。。。。。。。。。。")
	}

	chain.AddFilter(&PoliticalWordFilter{})
	if chain.Filter("hhhhh") == true {
		fmt.Println("政治敏感词过滤成功")
	} else {
		fmt.Println("政治敏感词过滤失败。。。。。。")
	}
}
