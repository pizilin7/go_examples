package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 10

		if want != got {
			t.Errorf("got %v, want %v, given %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 10

		if want != got {
			t.Errorf("got %v, want %v, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("SumAll", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 6})
		want := []int{3, 9, 7}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("SumAll1", func(t *testing.T) {
		got := SumAll1([]int{1, 2}, []int{3, 6})
		want := []int{3, 9, 7}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
