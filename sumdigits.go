package main

import "fmt"

func sumDigits(n int) int {
	sum := 0
	for n > 0 {
		rem := n % 10
		sum = sum + rem
		n = n / 10
	}
	return sum
}

func main() {
	fmt.Println(sumDigits(123))
}
