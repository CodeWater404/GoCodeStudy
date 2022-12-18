package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		//return a / b
		q, _ := div2(a, b) //_接收多余的返回值，但不需要用到
		return q
	default:
		panic("unsupported operation: " + op)
	}
}

/**eval2
** @Description: go可以有多个返回值，并且可以起名字！
** @param a
** @param b
** @param op
** @return int
** @return error
**/
func eval2(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		q, _ := div2(a, b) //_接收多余的返回值，但不需要用到
		return q, nil
	default:
		//panic("unsupported operation: " + op) //中断报错不好看
		return 0, fmt.Errorf("unsupported operation: %s ", op)
	}
}

//返回多个值
func div(a, b int) (int, int) {
	return a / b, a % b
}

//多返回值指定名字
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return //这里直接返回就行
}

/**apply
** @Description: 函数式编程.a , b 会传给op函数的中的参数
** @param op
** @param a
** @param b
** @param int
** @return int
**/
func apply(op func(int, int) int, a, b int) int {
	//反射，拿到函数名
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()

	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n ", opName, a, b)
	return op(a, b)
}

//重写库函数，参数和返回值转int
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

/**sum
** @Description: 求和
** @param numbers可变参数
** @return int返回值
**/
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	//fmt.Println(eval(1, 2, "-"))
	//fmt.Println(div(13, 3))
	//
	//q, r := div2(15, 4)
	//fmt.Println(q, r)

	//fmt.Println(eval(3, 5, "/"))

	//if conents, err := eval2(1, 2, "x"); err != nil {
	//	fmt.Println("Error : ", err)
	//} else {
	//	fmt.Println(conents)
	//}

	//fmt.Println(apply(pow, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}
