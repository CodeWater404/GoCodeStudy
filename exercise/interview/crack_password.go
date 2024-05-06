package main

import "fmt"

/*
*

	@author: CodeWater
	@since: 2024/5/6
	@desc: 破解密码

*
*/

// 正确的密码
var correctPassword = "A3k5d9"

// 递归生成密码组合
func generatePassword(prefix string, length int, characters string) {
	if length == 0 {
		if prefix == correctPassword {
			fmt.Println("找到正确密码:", prefix)
			return
		}
		return
	}

	for _, char := range characters {
		generatePassword(prefix+string(char), length-1, characters)
	}
}

func main() {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	passwordLength := 6

	// 递归生成密码组合
	generatePassword("", passwordLength, characters)

	fmt.Println("未找到正确密码")
}
