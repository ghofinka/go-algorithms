package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(s string) bool {
	// normalize string
	// set as lowercase, remove space, etc.
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}

	return true
}

func main() {
	words := []string{"katak", "level", "kasur rusak", "golang"}

	for _, w := range words {
		fmt.Printf("%s > %v\n", w, IsPalindrome(w))
	}
}
