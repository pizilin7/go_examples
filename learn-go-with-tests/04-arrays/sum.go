package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(list ...[]int) (sums []int) {
	c := len(list)
	sums = make([]int, c)
	for i, numbers := range list {
		sums[i] = Sum(numbers)
	}
	return
}

func SumAll1(list ...[]int) (sums []int) {
	for _, numbers := range list {
		sums = append(sums, Sum(numbers))
	}
	return
}
