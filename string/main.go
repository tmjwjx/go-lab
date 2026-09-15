package main

import (
	"fmt"
	"strconv"
)

func main() {

	// "1.234"字符串 转换为64位浮点数
	float, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(float)

	// "111"字符串 10进制 转换为64位整数
	n, _ := strconv.ParseInt("111", 10, 64)
	fmt.Println(n)

	// "0x1000"字符串 16进制 0是自己推断进制的意思 转换为64位整数
	n, _ = strconv.ParseInt("0x1000", 0, 64)
	fmt.Println(n)

	// "123"整数字符串 转换为整数
	n2, _ := strconv.Atoi("123")
	fmt.Println(n2)

	// "AAA"字符串 转换为整数 语法错误 返回 strconv.Atoi: parsing "AAA": invalid syntax
	n2, err := strconv.Atoi("AAA")
	fmt.Println(n2, err)
}
