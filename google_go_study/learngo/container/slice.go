package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/4/5
  @desc: 切片
		1.Slice本身没有数据，是对底层array的一个view
		2.当向slice中append数据的时候，如果到了最大长度cap会替换掉最后一个；如果超出了最大长度，那么在底层会新开一个arr去存放。原来的arr如果还在使用就会保存，如果不用就会被垃圾回收掉。
**/

/**sliceInit
** @Description: 切片左闭右开
**/
func sliceInit() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	fmt.Println(s)
}

/**updateSlice
** @Description: 切片中的值改变了，传给函数参数的变量内外都会变
** @param arr
**/
func updateSlice(arr []int) {
	arr[0] = 100
}

func main() {
	//sliceInit()

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//fmt.Println("=============================before update=============================")
	//s := arr[0:4]
	//fmt.Println(s)
	//fmt.Println("=============================after update=============================")
	//updateSlice(s)
	//fmt.Println(s)
	//fmt.Println(arr)

	fmt.Println("=============================extend slice=============================")
	s1 := arr[2:6] //2 , 3 , 4, 5
	//fmt.Println(s1[4])  //直接取不到 runtime error: index out of range [4] with length 4
	s2 := s1[3:5] //5 , 6//但是对于s2切片来说，它知道是对于arr切的，所以这里还是能打印s1[4],但是不能向前扩展
	fmt.Println(s1, s2)

	fmt.Println("=============================实际的容量=============================")
	fmt.Printf("s2=%v , len(s2)=%d , cap(s2)=%d\n", s2, len(s2), cap(s2))

}
