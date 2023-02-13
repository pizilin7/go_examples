package main

import (
	"fmt"
	"unsafe"
)

const (
	a = "hello"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	// 常量声明和初始化
	const stringConstant = "string"
	const boolConstant bool = true
	const intConstant1, intConstant2 = 1, 2

	fmt.Println(stringConstant, boolConstant, intConstant1, intConstant2)

	// const stringConstant100, missing constant value
	// stringConstant100 = "hello",初始化要一起

	// 多个常量
	const intConstant100, intConstant200, boolConstant300 = 1, 2, false
	fmt.Println(intConstant100, intConstant200, boolConstant300)

	// 简写方式
	const (
		company string = "hello"
		salary         = 1000.00
	)
	fmt.Println(company, salary)

	// 打印全局常量
	fmt.Println(a, b, c)
}
