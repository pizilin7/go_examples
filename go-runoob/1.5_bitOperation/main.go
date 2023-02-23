package main

import "fmt"

func main() {
	var a uint = 14 /* 14 = 1110 */
	var b uint = 111 /* 111 = 1101111 */
	var c uint = 0

	c = a & b
	fmt.Println(c)
	c = a | b
	fmt.Println(c)
	c = a ^ b
	fmt.Println(c)
	c = a << 2
	fmt.Println(c)
	c = a >> 2
	fmt.Println(c)
}