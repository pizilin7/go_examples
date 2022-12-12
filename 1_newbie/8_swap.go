package main

import "fmt"

func swap(x, y string) (string, string) {
	// 返回多个值
	return y, x
}

func main() {
	a, b := "world", "hello"
	fmt.Println("a =", a, " ,b =", b)
	a, b = swap(a, b)
	fmt.Println("swap, a =", a, " ,b =", b)
}