// 数据类型
package main

import (
	"fmt"
)

func main() {
	// bool
	fmt.Println("===>bool")
	fmt.Printf("%t\n", true)
	fmt.Printf("%t\n", false)

	// number
	// 数字类型：uint,uint8,uint16,uint32,uint64,int,int8,int16,int32,int64
	fmt.Println("===>uint")
	var uint_n uint = 18446744073709551616 - 1
	fmt.Printf("uint_n = %d\n", uint_n)

	var uint8_n uint8 = 255
	fmt.Printf("uint8_n = %d\n", uint8_n)

	var uint16_n uint16 = 0xffff
	fmt.Printf("uint16_n = %d\n", uint16_n)

	var uint32_n uint32 = 0xffffffff
	fmt.Printf("uint32_n = %d\n", uint32_n)

	var uint64_n uint64 = 0xffffffffffffffff
	fmt.Printf("uint64_n = %d\n", uint64_n)

	var int_n int = 0xfffffffffffffff
	fmt.Printf("int_n = %d\n", int_n)

	fmt.Println("===>int")

	var int8_n int8 = 127
	fmt.Printf("int8_n = %d\n", int8_n)

	var int16_n int16 = 0xfff
	fmt.Printf("int16_n = %d\n", int16_n)

	var int32_n int32 = 0xfffffff
	fmt.Printf("int32_n = %d\n", int32_n)

	var int64_n int64 = 0xfffffffffffffff
	fmt.Printf("int64_n = %d\n", int64_n)

	// float
	// 浮点类型：float32、float64、complex64、complex128
	fmt.Println("===>float")
	var float32_n float32 = 0xffffffff
	fmt.Printf("float32_n = %e\n", float32_n)
	fmt.Printf("float32_n = %f\n", float32_n)
	fmt.Printf("float32_n = %G\n", float32_n)

	var float64_n float32 = 0xffffffffffffffff
	fmt.Printf("float64_n = %e\n", float64_n)
	fmt.Printf("float64_n = %f\n", float64_n)
	fmt.Printf("float64_n = %G\n", float64_n)

	// string
	// 字符串类型：byte、rune、string
	fmt.Println("===>string")
	var str string = "hello world"
	fmt.Printf("str = %s\n", str)
	fmt.Printf("str = %q\n", str)
	fmt.Printf("str = %x\n", str)

	var byte_n byte = 101 // unit8别名，表示字符
	fmt.Printf("byte_n = %c\n", byte_n)
	fmt.Printf("byte_n = %d\n", byte_n)

	fmt.Printf("rune_n = %U\n", []rune("Go语言"))
}

// 数据类型：https://www.runoob.com/go/go-data-types.html
// Go 之 格式化输出 《Go语言基础》：https://juejin.cn/post/7042636455239221256
// 详解 Go 中的 rune 类：https://www.cnblogs.com/cheyunhua/p/16007219.html
