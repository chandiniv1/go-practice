package main

import (
	"fmt"
	"math"
)

func digitCount(n int) int {
	c := 0
	for c > 0 {
		c++
		n = n / 10
	}
	return c
}

func isArmstrong(n int) bool {
	temp := n
	sum1 := 0
	count := digitCount(temp)
	for temp > 0 {
		rem := temp % 10
		d := math.Pow(float64(rem), float64(count))
		sum1 = sum1 + int(d)
		temp = temp / 10
	}
	if sum1 == n {
		return true
	}
	return false
}
func main() {
	fmt.Println(isArmstrong(8208))
}
