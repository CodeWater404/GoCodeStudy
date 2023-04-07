package main

import (
	"fmt"
	"unicode/utf8"
)

/**
  @author: CodeWater
  @since: 2023/4/7
  @desc: string
**/

func main() {
	s := "Yes哈哈哈啊哈!" //utf-8
	fmt.Println(s)

	for _, b := range []byte(s) {
		fmt.Printf("%x , ", b)
	}
	fmt.Println("==========================================================")

	for i, ch := range s { //ch is a rune
		fmt.Printf("%d , %x ", i, ch)
	}

	fmt.Println("==========================================================")

	fmt.Println("Rune count:", uft8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Print("%x ", ch)
	}

	fmt.Println("==========================================================")

	for i, ch := range []rune(s) {
		fmt.Printf("%d , %x ", i, ch)
	}
	
}
