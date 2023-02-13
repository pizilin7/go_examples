package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("条件语句：if")
	fmt.Println("%t\n", os.Args)
	if len(os.Args) > 1 && os.Args[1] == "Hello" {
		fmt.Println("Hello World")
	} else {
		fmt.Println("Hack the Planet")
	}
}
