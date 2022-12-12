package main

import "fmt"

func swap3(x, y *int) {
	fmt.Printf("before, x = %d, y = %d\n", *x, *y)
	*x, *y = *y, *x // 地址交换
	fmt.Printf("after, x = %d, y = %d\n", *x, *y)
}

func main() {
	a, b := 100, 200
	fmt.Printf("swap before, a = %d, b = %d\n", a, b)
	swap3(&a, &b)
	fmt.Printf("swap after, a = %d, b = %d\n", a, b)

	// 由上可知，swap函数中是引用，实参和形参互相影响
}
