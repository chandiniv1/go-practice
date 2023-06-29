package main

import (
	"fmt"
)

func decimaltobinary(n int) int {
	sum := 0
	for n > 0 {
		rem := n % 2
		sum = sum*10 + rem
		n = n / 2
	}
	return sum
}

func main() {
	fmt.Println(decimaltobinary(5))
}
