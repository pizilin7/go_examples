package main

import (
	"fmt"
	"unsafe"
)

const (
	aa = "abc"
	bb = len(aa)
	cc = unsafe.Sizeof(aa)
)
func main() {
	const LENGTH int = 10
	const WIDTH int = 20
	var area int
	const a, b, c = 1, false , "str"

	area = LENGTH * WIDTH
	fmt.Println("面积位： %d", area)
	println(a, b, c)

	//变量做枚举
	const (
		Unknown = 0
		Female = 1
		Male = 2
	)

	fmt.Println(aa, bb, cc)
	fmt.Println(Unknown, Female, Male)
}