package main

import "fmt"

func main() {
	//initArray()
	//initDoubleDimensionalArray()
	testAverage()
}

func initArray() {
	var a[10] int
	var i ,j int
	//var b = [10] int {1, 2, 10}
	//var bb = [...] float32 {1.0, 2.0}
	for i = 0; i < 10; i ++ {
		a[i] = 100 + i
	}

	for j = 0; j < 10; j ++ {
		fmt.Printf("x: ", a[j])
	}
}

func initDoubleDimensionalArray() {
	var a[10][10] int
	var i, j = 10, 10

	for i = 0; i < 10; i ++ {
		for j = 0; j < 10; j++ {
			a[i][j] = (i + 1) * (j + 1)
		}
	}

	for i = 0; i < 10; i ++ {
		for j = 0; j < 10; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

func getAverage(arr [5]int, size int) float32 {
	var i int
	var avg float32
	var sum int

	for i = 0; i < size;  i ++{
		sum += arr[i]
	}

	avg = float32(sum / size)
	return avg
}

func testAverage() {
	var balance = [5]int {1, 2, 3, 4, 11}
	var avg float32

	avg = getAverage(balance, len(balance))
	fmt.Println(avg)
}