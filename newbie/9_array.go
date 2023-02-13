package main

import "fmt"

func main() {
	var strs = []string{"hello", "world", "a", "b", "c"}
	var ints = []int{1, 3, 5, 7, 9}
	var maps = map[int]string{}
	var maps2 = map[string]interface{}{}

	for i := 0; i != 5; i++ {
		fmt.Println(ints[i], "\t", strs[i])
		maps[ints[i]] = strs[i]
		maps2[strs[i]] = maps
	}

	fmt.Println(maps)
	fmt.Println(maps2)
}
