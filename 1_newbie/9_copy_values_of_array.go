package main

import "fmt"

func main() {
	my_arr1 := [5]string{"a", "b", "c", "e", "f"}
	my_arr2 := my_arr1 // 拷贝了一份
	fmt.Println("my_arr1 =", my_arr1)
	fmt.Println("my_arr2 =", my_arr2)

	my_arr1[0] = "a100" // 访问数组元素
	fmt.Println("my_arr1 =", my_arr1)
	fmt.Println("my_arr2 =", my_arr2)
}
