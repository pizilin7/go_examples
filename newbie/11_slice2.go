package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	numbers2 := []int{1, 2, 3}
	printSlice(numbers2)

	var numbers3 []int
	printSlice(numbers3)
	if numbers3 == nil {
		fmt.Println("numbers3 nil")
	}

	var numbers4 = []int{}
	printSlice(numbers4)
	if numbers4 == nil {
		fmt.Println("numbers4 nil")
	}
}

func printSlice(x []int) {
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(x), cap(x), x)
}
