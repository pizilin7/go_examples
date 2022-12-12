package main

import "fmt"

type cb func(int) int

func main() {
	testCB(1, callBack)
}

func testCB(x int, f cb) {	// 函数f作为testCB函数的一个参数来使用
	f(x)
}

func callBack(x int) int {
	fmt.Println("x :", x)
	return x
}
