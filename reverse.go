package main

import "fmt"

func reverse(n int) int {
	sum := 0
	for n > 0 {
		rem := n % 10
		sum = sum*10 + rem
		n = n / 10
	}
	return sum

}

func main() {
	fmt.Println(reverse(123))
}
