package main

import "fmt"

// 定义结构体
type Circle struct {
	radius float64
}

// 该方法包含Circle类型对象
func (c Circle) getArea() float64 {
	return c.radius * c.radius * 3.14
}

func main() {
	var c1 Circle = Circle{
		radius: 10.00,
	}
	fmt.Println("area =", c1.getArea())
}
