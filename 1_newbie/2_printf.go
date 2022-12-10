package main

import "fmt"

type data struct {
	a, b int
	c, d string
	e, f bool
	g, h float64
}

func printx(str string, d data) {
	fmt.Printf("%"+str+":\t"+str+"\n", d)
}

func main() {
	x := data{0, 1, "a", "b", true, false, 3.1415926535, 2.357111317192329}

	// printx("%v", x)
	// printx("%+v", x)
	// printx("%#v", x)
	// printx("%t", x)
	// printx("%T", x)
	// printx("%d", x)
	// printx("%b", x)
	// printx("%c", x)
	// printx("%x", x)
	// printx("%f", x)
	// printx("%e", x)
	// printx("%E", x)
	// printx("%s", x)
	// printx("%q", x)
	// printx("%p", x)
}
