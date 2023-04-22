package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//函数外的变量定义（不是全局，go没有这个概念，作用域是在当前包下）
var a = 1
var b = 2

//b := 3 //函数外不能用:

//简洁写法
var (
	d = 4
	f = 5
	q = "ds"
)

//默认值
func variableZeroValue() {
	var a int
	var b string
	//%q把空串可视化打出来（一般是%s）
	fmt.Printf("%d %q\n", a, b)
}

//赋初值
func variableInitialValue() {
	var a, b int = 3, 4
	var s = "abc"
	fmt.Println(a, b, s)
}

//可从后推断类型，省略类型
func variableTypeDeduction() {
	var a, b = 1, 2
	var c, d = 3, true //类型还可以不一样

	fmt.Println(a, b, c, d)
}

//连续赋值的简洁写法
func variableShorter() {
	a, b, c, d := 1, true, "123", 1.2 //":"相当于定义变量

	fmt.Println(a, b, c, d)
	//再次赋值就不能用:
	b = false
	fmt.Println(b)
}

//欧拉公式
func euler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))

	//公式: e^iΠ + 1 = 0
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

//类型转换是强制的
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func triangleTest() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

//const
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//enum
func enums() {
	//普通枚举
	const (
		cpp = iota //iota自增
		java
		_ //“_”跳过
		python
		javascript
	)

	//自增枚举
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, java, python)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

//没有char，只有rune

func main() {
	//fmt.Println("Hello world!!!")
	//
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeduction()
	//variableShorter()
	//
	////函数外变量打印
	//fmt.Println(a, b, d, f, q)

	//euler()

	//triangle()

	//consts()

	enums()
}
