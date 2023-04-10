package main

import "fmt"

/**
  @author: CodeWater
  @since: 2023/4/5
  @desc: map基本概念和操作
		 1. 获取mao中的值时，也是返回多个参数的，第一个是值，第二个是bool。
			ageExist, ok := m["agee"]
**/

func mapInit() {
	m := map[string]string{
		"name": "code",
		"age":  "18",
		"city": "san francosco",
	}

	m2 := make(map[string]string) //m2 is empty map

	var m3 map[string]string //m3 is nil map
	/*
		var m4 = map[int]int{
			1 : 1 ,
			2 : 2,
		}
	*/

	fmt.Println(m, m2, m3)

	fmt.Println("=============================traversing map=============================")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("=============================print map's value=============================")
	name := m["name"]
	fmt.Println(name)
	ageError := m["agee"]
	fmt.Println(ageError) //print empty string

	fmt.Println("=============================optimize to print=============================")
	if ageExist, ok := m["agee"]; ok { //用ok去做if的判断,int不能直接转换为bool
		fmt.Printf("agee key is exist:%s\n", ageExist)
	} else {
		fmt.Println("agee key is not exist")
	}

	fmt.Println("=============================delete map's value=============================")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}

func main() {
	mapInit()

}
