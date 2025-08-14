package main

import (
	"fmt"
	"slices"
)

// input []int
// output
func MinMaxSum(arr []int) {
	var listSum []int32
	for i := 0; i < 5; i++ {
		sum := 0
		for j := 0; j < 5; j++ {
			if i != j {
				sum += arr[j]
			}
		}
		listSum = append(listSum, int32(sum))
	}
	fmt.Println("output : ", slices.Min(listSum), slices.Max(listSum))
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	fmt.Println("input : ", input)
	MinMaxSum(input)
}
