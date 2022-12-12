package main

import "fmt"

func main() {
	map1 := make(map[int]float32)
	for i := 0; i < 5; i++ {
		map1[i] = float32(i*10) + float32(0.01)
	}

	for index, data := range map1 {
		fmt.Printf("key = %d, value = %f\n", index, data)
	}

	for key, _ := range map1 {
		fmt.Printf("key is %v\n", key)
	}

	for _, value := range map1 {
		fmt.Printf("value is %f\n", value)
	}
}
