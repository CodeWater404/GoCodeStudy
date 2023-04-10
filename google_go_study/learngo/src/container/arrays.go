package main

import "fmt"

/**arrayInit
** @Description: 数组的初始化几种方式
**/
func arrayInit() {
	//==========================array init=====================
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{4, 5, 6}
	fmt.Println(arr1, arr2, arr3)
	fmt.Println("========================")

	//==========================二维=====================
	var grid [4][5]int
	fmt.Println(grid)

	fmt.Println("==========================go through the groups================================")

	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}

	//print index and value
	//for i := range arr3 {
	//	fmt.Println(arr3[i])
	//}

	//for k, v := range arr3 {
	//	fmt.Println(k, v)
	//}

	//only print value
	for _, v := range arr3 {
		fmt.Println(v)
	}

}

/**printArray
** @Description: 数组作为函数参数的时候，会拷贝一份传过去。函数内部的改变不会影响函数外部数组的改变。（别的语言会改变，比如java）
** @param arr
**/
func printArray(arr [3]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrayRef(arr *[3]int) {
	arr[0] = 100
	for k, v := range arr {
		fmt.Println(k, v)
	}
}

func main() {
	//arrayInit()

	//[3]int and [5]int 不是同一种类型
	arr1 := [3]int{1, 2, 3}
	//arr2 := [5]int{4, 5, 6, 7, 8}
	printArray(arr1)
	//printArray(arr2) //error

	fmt.Println("============================数组参数值传递==============================")
	//printArray(arr1)
	//for _, i := range arr1 {
	//	fmt.Println(i)
	//}
	//改为引用
	printArrayRef(&arr1)
	for k, v := range arr1 { //重新输出，值发生变化
		fmt.Println(k, v)
	}

}
