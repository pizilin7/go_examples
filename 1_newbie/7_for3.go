package main

import "fmt"

func main() {
	i := 1
	for {
		i += i
		fmt.Println(i)
		if i > 10 {
			break
		}
	}
}
