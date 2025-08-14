package main

import (
	"fmt"
)

// input n integer
// output
// the function will return the word "Fizz" when the number is divisible by 3
// the function will return the word "Buzz" when the number is divisible by
// the function will return the word "FizzBuzz" when the number is divisible by 3 & five
func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 && i%5 != 0 {
			fmt.Println("Fizz")
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	FizzBuzz(30)
}
