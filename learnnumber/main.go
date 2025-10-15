package main

import (
	"fmt"
	"math"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// func isPrimeNumber(n int) bool {
func isPrimeNumber(n int) string {
	if n < 2 {
		// return false
		return "not a prime number"
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			// return false
			return "not a prime number"
		}
	}
	// return true
	return "a prime number"
}

func maxNumber() {

}

func removeDuplicateNumber() {

}

// hitung jumlah digitnya lalu pangkatkan tiap angka dengan jumlah digit tersebut
func isArmstrong(n int) bool {
	sum, temp := 0, n
	digits := int(math.Log10(float64(n))) + 1

	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}

	return sum == n
}

func main() {
	num := int(5)
	num2 := int(11)
	num3 := int(50)

	numArm1 := int(153)
	numArm2 := int(525)

	fmt.Printf("factorial from %d is %d\n", num, factorial(num))
	fmt.Printf("%d is %s\n", num, isPrimeNumber(num))
	fmt.Printf("%d is %s\n", num2, isPrimeNumber(num2))
	fmt.Printf("%d is %s\n", num3, isPrimeNumber(num3))

	fmt.Printf("%d is armstrong %v\n", numArm1, isArmstrong(numArm1))
	fmt.Printf("%d is armstrong %v\n", numArm2, isArmstrong(numArm2))
}
