package main

import "fmt"

func main() {
	createMap();
}

func createMap() {
	var stringMap = make(map[string]string)

	stringMap["A"] = "AAA"
	stringMap["B"] = "BBB"
	stringMap["C"] = "CCC"

	fmt.Printf("%v\n", stringMap)

	delete(stringMap, "A")

	var intMap = map[int]string{
		1 : "one",
		2 : "two",
		3 : "three",
	};
	fmt.Println(intMap)

	value, ok := intMap[1]
	if ok {
		fmt.Printf("key = 1; Value = %v", value)
	}
}