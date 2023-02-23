package main

import "fmt"

func main() {
	x, y := 10, 199
	fmt.Printf("%d, %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("%d, %d\n", x, y)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
