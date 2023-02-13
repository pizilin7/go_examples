// 变量
package main

import "fmt"

var b bool = false
var b1 bool

// b2 := true,不能再函数外使用

func main() {
	var str string = "hello" // var 标识符 类型 = 赋值
	var str1 = "hello"       // var 标识符 = 赋值，自动推导类型
	str2 := "hello"          // 标识符 := 赋值，自动推导类型

	var str3 string // var 标识符 类型
	str3 = "hello"  // 赋值

	fmt.Println(str, str1, str2, str3)

	b1 = true
	fmt.Println(b, b1)
}
