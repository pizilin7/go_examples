// iota 每新增一次常量声明计数增加一次
package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

const (
	a1 = iota
	b1
	c1
)

// iota 如果声明，在const中还是在自动增加
const (
	a100 = iota
	a101
	a102
	a103 = "hello"
	a104
	a105 = iota
	a106 = 100
	a107 = iota
)

func main() {
	fmt.Println("const iota")
	fmt.Println(a, b, c)
	fmt.Println(a1, b1, c1)
	fmt.Println(a100, a101, a102, a103, a104, a105, a106, a107)

	fmt.Println("const iota 位运算")
	const (
		i = 2 << iota
		j = 2 << iota
		k
		l
	)

	fmt.Println(i, j, k, l)
}
