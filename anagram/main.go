package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(s string) string {
	runes := strings.Split(s, "")
	fmt.Println(runes)
	sort.Strings(runes)
	fmt.Println(runes)
	return strings.Join(runes, "")
}

func isAnagram(a, b string) bool {
	return sortString(strings.ToLower(a)) == sortString(strings.ToLower(b))
}

func main() {
	fmt.Println(isAnagram("listen", "silent"))
	fmt.Println(isAnagram("listen", "leisen"))
}
