package main

import (
	"fmt"
	"math"
)

func digitCount(n int) int {
	c := 0
	for n > 0 {
		c++
		n = n / 10
	}
	return c
}

func binarytodecimal(n int) int {
	count := digitCount(n)
	sum := 0
	for n > 0 {
		for i := 0; i < count; i++ {
			rem := n % 10
			d := math.Pow(float64(2), float64(i))
			sum = sum + rem*int(d)
			n = n / 10
		}

	}
	return sum
}

func main() {
	fmt.Println(binarytodecimal(101))
}
