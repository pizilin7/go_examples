package main

import "fmt"

func max(num1, num2 int) int {		// func 函数名称(参数列表) [返回类型]
	if num1 > num2 {
		return num1
	}

	return num2
}

func main() {
	fmt.Println(max(100, 200))
}
