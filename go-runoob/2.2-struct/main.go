package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	fmt.Println(Books{"go语言", "lqw", "go 语言教程", 123456})
	fmt.Println(Books{title: "go语言"})
}
