package main

import "fmt"

func swap2(x, y int) {
	fmt.Printf("before, x = %d, y = %d\n", x, y)
	x, y = y, x
	fmt.Printf("after, x = %d, y = %d\n", x, y)
}

func main() {
	a, b := 100, 200
	fmt.Printf("swap before, a = %d, b = %d\n", a, b)
	swap2(a, b)
	fmt.Printf("swap after, a = %d, b = %d\n", a, b)

	// 由上可知，swap函数中是值传递，实参和形参互不影响
}
