package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/4/5
  @desc: 切片的操作
**/

/**sliceInit
** @Description: slice不用arr去初始化，也是可以赋值的

**/
func sliceInit2() {
	var s []int // zero value for slive is nil
	for i := 0; i < 100; i++ {
		//打印出每次slice的扩容:0 , 1 , 2 , 4 , 8 , 16...128
		printSliceLen(s)
		s = append(s, 2*i+1)
	}

	fmt.Println(s)

}

func printSliceLen(slice []int) {
	fmt.Printf("len=%d , cap=%d \n", len(slice), cap(slice))
}

func sliceOption() {
	s1 := []int{1, 2, 3} //固定赋值
	printSliceLen2(s1)

	s2 := make([]int, 10) //开一个长度len是10的slice
	printSliceLen2(s2)

	s3 := make([]int, 10, 30) //开一个长度len是10的slice，但是cap为30
	printSliceLen2(s3)

	fmt.Println("=============================copy slice=============================")
	copy(s2, s1) //复制之后，s2前三个值就是s1的值
	printSliceLen2(s2)

	fmt.Println("=============================delete element from slice=============================")
	s2 = append(s2[:2], s2[3:]...) //用切片操作删除第三个下标的元素，s2[3:]...是表示下标3到最后的元素，这个可变参数的写法
	printSliceLen2(s2)

	fmt.Println("=============================popping from front and tail =============================")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front, s2)

	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail, s2)

}

func printSliceLen2(slice []int) {
	fmt.Printf("value:%v , len=%d , cap=%d \n", slice, len(slice), cap(slice))
}

func main() {
	//sliceInit2()

	sliceOption()
}
