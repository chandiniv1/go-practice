package main

import "fmt"

func isPrime(a int) bool {
	if a == 1 || a == 2 {
		return true
	} else {
		for i := 2; i < a; i++ {
			if a%i == 0 {
				return false
			}
		}
		return true
	}
}

func printPrime(n int) {
	for i := 1; i < n; i++ {
		if isPrime(i) == true {
			fmt.Println(i)
		}
	}
}

func main() {
	printPrime(100)
}
