package main

import "fmt"

func main() {
	max := max(10, 20)
	fmt.Println("最大的数值：%d", max)

	a, b := swap("world", "hello")
	fmt.Println(a, b)
}

func max(num1 int, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x, y string) (string, string) {
	return y, x
}