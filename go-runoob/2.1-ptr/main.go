package main

import "fmt"

func main() {
	//copyArray()
	testPPtr()
}

func test() {
	var a int = 100
	var aPtr *int
	aPtr = &a
	fmt.Println(aPtr)
	fmt.Println(*aPtr)
}

func copyArray() {
	a := [3]int{1, 2, 3}
	var ptr [3]* int

	for i := 0; i < len(a); i++ {
		ptr[i] = &a[i]
	}

	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
 }

func testPPtr() {
	a := 10
	ptr := &a
	pptr := &ptr

	fmt.Printf("变量 a = %d\n", a )
	fmt.Printf("指针变量 *ptr = %d\n", *ptr )
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}