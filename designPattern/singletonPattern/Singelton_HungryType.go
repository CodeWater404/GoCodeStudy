package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/7/4
  @desc: 单例模式： 饿汉式和懒汉式.
	三个要点：
		一是某个类只能有一个实例：
		二是官必须自行创建这个实例；
		三是它必须自行向整个系统提供这个实例。
	1. 饿汉式：不管系统是否运行直接new一个对象
	2. 懒汉式：是在系统中需要改对象的时候才会去创建这个对象
**/

type singelton struct{}

var instance *singelton = new(singelton)

func GetInstance() *singelton {
	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("单例的一个方法")
}

func main() {
	s1 := GetInstance()
	s1.SomeThing()

	s2 := GetInstance()
	if s1 == s2 {
		fmt.Println("s1 == s2")
	}
}
