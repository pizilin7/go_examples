package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	printSlice(numbers[1:4])
	printSlice(numbers[0:4])
	printSlice(numbers[6:])

	fmt.Println("===========>append,copy")

	var numbers2 []int
	printSlice(numbers2)

	numbers2 = append(numbers2, 0, 1, 2, 3)
	printSlice(numbers2)

	var numbers3 []int = make([]int, 3, 10)
	copy(numbers3, numbers2)
	printSlice(numbers3)
}

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d, value = %v\n", len(s), cap(s), s)
}
