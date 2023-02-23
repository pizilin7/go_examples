package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d = "word"
		e
		f = iota
		g
		h = 1234
		i
		j
		k = iota
	)
	fmt.Println(a, b, c, d, e, f, g, h, i, j ,k)

	const (
		l = 1 << iota
		m = 2 << iota
		n
		o
	)
	fmt.Println(l, m, n, o)
}