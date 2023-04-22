package main

import (
	"fmt"
	"unicode/utf8"
)

/**
  @author: CodeWater
  @since: 2023/4/7
  @desc: string
		1. Golang中没有专门的字符类型，如果要存储单个字符(字母)，一般使用byte来保存。byte等价于uint8
			eg: var ch byte = 's'
		2. rune等价于int32
		3. 不同：
			Golang中不能直接对字符串string进行下标操作，都是利用rune跟byte也进行字符串的操作。
			rune是用来区分字符值和整数值的
			rune 等同于int32，即4个字节长度,常用来处理unicode或utf-8字符。
			byte 等同于int8，即一个字节长度，常用来处理ascii字符
			中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。
			ASCII编码是1个字节，而UTF-8是可变长的编码
			当要表示中文等非ASCll编码的字符时，需要使用UTF-8编码来保证不会乱码。
			UTF8编码下一个中文汉字由3~4个字节组成，而字符串是由byte字节组成，所以长度也是byte字符长度，这样遍历时遇到中文就乱码了
			所谓对字符串的修改其实不是对字符串本身的修改，而是复制字符串，同时修改值，即重新分配来内存。
			在go中修改字符串，需要先将字符串转化成数组，[]byte 或 []rune，然后再转换成 string型。

**/

func main() {
	s := "Yes哈哈哈啊哈!" //utf-8,一个中文三个字节。字节数计算len是19，字符数计算len是9
	fmt.Println(s)
	fmt.Printf("len(byte(s)):%d\n", len([]byte(s))) // 返回len(str):19

	// 以字符数来计算长度
	fmt.Printf("len(rune(s)):%d\n", len([]rune(s))) // 返回len(rune(str)):9

	for _, b := range []byte(s) {
		fmt.Printf("%d , ", b)
	}
	fmt.Println()
	fmt.Println("==========================================================")

	for i, ch := range s { //ch is a rune
		fmt.Printf("(index:%d , value:%d) , ", i, ch)
	}

	fmt.Println("==========================================================")

	//统计有多少个字符（一个英文算一个，一个中文也算一个）
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		//DecodeRune从byte数组中拿到第一个utf8的字符进行解码，然后返回该字符rune和其size
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	fmt.Println("==========================================================")

	//把string分割成一个一个字符rune，i是下标ch是对应字符
	for i, ch := range []rune(s) {
		fmt.Printf("(%d  %c), ", i, ch)
	}

}
