package main

import "fmt"

func main() {
	fmt.Println("===>strs")
	var strs = []string{"hello", "world"}
	for index, data := range strs {
		fmt.Println(index, data)
	}

	fmt.Println("===>numbers")
	var numbers = []int{100, 200, 300, 400}
	for index, data := range numbers {
		fmt.Println(index, data)
	}
}
