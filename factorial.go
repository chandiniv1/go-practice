package main

import "fmt"

func factorial(n int) int {
	fact := n
	for i := n - 1; i > 0; i-- {
		fact = fact * i
	}
	return fact
}

func main() {
	fmt.Println(factorial(5))
}
