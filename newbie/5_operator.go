package main

import "fmt"

func main() {
	fmt.Println("算术运算符")
	a := 100
	b := 200
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)
	a++
	fmt.Println("a ++ =", a)
	a--
	fmt.Println("a -- =", a)

	fmt.Println("关系运算符")
	var c bool
	c = a == b
	fmt.Println("a == b,", c)
	fmt.Println("a != b,", c)
	fmt.Println("a > b,", a > b)
	fmt.Println("a < b,", a < b)
	fmt.Println("a >= b,", a >= b)
	fmt.Println("a <= b,", a <= b)

	fmt.Println("逻辑运算符")
	fmt.Println("true && false,", true && false)
	fmt.Println("false && false,", false && false)
	fmt.Println("true && true", true && true)
	fmt.Println("true || false", true || false)
	fmt.Println("true || true", true || true)
	fmt.Println("false || false", false || false)
	fmt.Println("!false,", !false)
	fmt.Println("!true,", !true)

	fmt.Println("位运算")
	bit1 := 0b00010101 // 21
	bit2 := 0b00001001 // 9

	fmt.Println("bit1 & bit2 =", bit1&bit2) // 0b00000001
	fmt.Println("bit1 | bit2 =", bit1|bit2) // 0b00011101
	fmt.Println("bit1 ^ bit2 =", bit1^bit2) // 0b00011100
	fmt.Println("bit1 << 3 =", bit1<<3)     // 0b10101000
	fmt.Println("bit2 >> 2 =", bit2 >> 2)	// ob10

	fmt.Println("赋值运算")
	c100 := a + b
	fmt.Printf("%d = a + b\n", c100)
}
