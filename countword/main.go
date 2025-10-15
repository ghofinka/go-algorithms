package main

import (
	"fmt"
	"strings"
)

func CountVowels(s string) int {
	vowels := "aiueo"
	count := 0

	for _, c := range strings.ToLower(s) {
		if strings.ContainsRune(vowels, c) {
			count++
		}
	}

	return count
}

func CountWords(s string) int {
	s = strings.ToLower(s)
	words := strings.Fields(s)
	return len(words)
}

func CountEachWord(s string) map[string]int {
	s = strings.ToLower(s)
	words := strings.Fields(s)
	counts := make(map[string]int)

	for _, w := range words {
		counts[w]++
	}

	return counts
}

func main() {
	text := "the quick brown fox jumps over the lazy dog and fox"
	text2 := "Go is fun and go is powerful"

	fmt.Printf("count vowels from \"%s\", %d\n", text, CountVowels(text))
	fmt.Println(CountWords(text2))
	fmt.Println(CountEachWord(text2))
}
