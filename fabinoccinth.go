package main

import "fmt"

func main() {
	a := 0
	b := 1
	n := 10
	count := 0
	if n <= 0 {
		fmt.Println("enter a positive number")
	} else if n == 1 {
		fmt.Println(0)
	} else {
		for count < n {
			fmt.Println(a)
			c := a + b
			a = b
			b = c
			count++
		}
	}
}
