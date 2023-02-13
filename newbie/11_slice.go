package main

import "fmt"

func main() {
	var str string = "Lorem ipsum dolor sit amet"
	fmt.Println(str[6:11])

	s := make([]string, 3)

	s = append(s, "abc")
	fmt.Println(s)
}