package main

import (
	"fmt"
	"math"
	"sort"
)

func closestNumbers(numbers []int32) {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	minDiff := int32(math.MaxInt32)
	found := make(map[[2]int32]bool)

	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]
		if diff < minDiff {
			minDiff = diff
		}
	}

	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]
		if diff == minDiff {
			pair := [2]int32{numbers[i], numbers[i+1]}
			if !found[pair] {
				found[pair] = true
				fmt.Println(numbers[i], numbers[i+1])
			}
		}
	}
}

func main() {

}
