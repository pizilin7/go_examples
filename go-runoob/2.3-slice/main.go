package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	var a = []int{1, 2, 4, 5, 6}
	printSlice(a)
	printSlice(a[1: 5])
	printSlice(a[3:])
	printSlice(a[:4])

	a = append(a, 2, 3, 4 ,5)
	printSlice(a)

	aa := make([]int, len(numbers), (cap(numbers) * 2))
	copy(aa, a)
	printSlice(aa)
}

func printSlice(x []int) {
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(x), cap(x), x)
}