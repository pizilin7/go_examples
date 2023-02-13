package main

import "fmt"

func fibonacci(n uint64) {
	var a uint64 = 0
	var b uint64 = 1
	var i uint64
	var sum uint64
	for i = 0; i < n; i++ {
		fmt.Println(a)
		sum = a + b
		a = b
		b = sum
	}
}

func main() {
	a := uint64(10)
	fmt.Println(fibonacci(a))
}
