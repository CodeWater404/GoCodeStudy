package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/4/12
  @desc: 函数式编程
		“正统”函数式编程（相当于一种规范）：
		1.不可变性：不能有状态，只有常量和函数
		2.函数只能有一个参数
**/

/**adder
** @Description: 闭包：累加器。参数、返回值是一个函数
** @return func(int) int
**/
func adder() func(int) int {
	sum := 0
	//这里最终返回的是：函数、sum的引用
	return func(v int) int {
		//sum在这里是自由变量，v是局部变量
		sum += v
		return sum
	}
}

//type定义函数类型，返回值两个：int和iAdder自己
type iAdder func(int) (int, iAdder)

/**adder2
** @Description: 正统的函数式编程：没有变量，只有参数和常量
** @param base
** @return iAdder
**/
func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...+%d = %d\n", i, a(i))
	}

	fmt.Println("=============================正统函数式编程=============================")
	b := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		//s =》 int ， b =》 iAdder
		s, b = b(i)
		fmt.Printf("0+1+...+%d = %d \n", i, s)
	}
}
