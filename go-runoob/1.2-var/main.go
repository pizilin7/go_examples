package main

import "fmt"

var (
	aa int
	bb bool
	ss string
)
var cc , dd int = 1, 2
var ee, ff = 1 ,"hello"

func main() {
	var a string = "lqw"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var test_name string
	var test_int uint8
	var test_bool bool
	fmt.Println(test_bool, test_int, test_name)
}
