package main

import (
	"fmt"
)

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice(intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice(float32Slice))

	var intSlice2 = []int{}
	fmt.Println(isEmpty(intSlice2))

	var float32Slice2 = []float32{1, 2, 3}
	fmt.Println(isEmpty(float32Slice2))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}
