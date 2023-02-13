package main

import "fmt"

func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)
	return avg
}

func getAverage2(arr *[5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)
	return avg
}

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	var avg = getAverage(arr, 5)

	fmt.Printf("avg = %f\n", avg)

	var avg2 = getAverage2(&arr, 5)
	fmt.Printf("avg2 = %f\n", avg2)
}