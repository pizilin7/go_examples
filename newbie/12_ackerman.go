package main

import "fmt"

func ack(n, m int64) int64 {
	for n != 0 {
		if m == 0 {
			m = 1
		} else {
			m = ack(n, m-1)
		}
		n = n - 1
	}
	return m + 1
}

func main() {
	fmt.Println(ack(2, 2))
}