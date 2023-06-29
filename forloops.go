package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	i := 0
	for i != 10 {
		fmt.Println(i)
		i++
	}
	for {
		fmt.Println("helloo")
		break
	}
	arr := []int{1, 2, 3, 4, 5}
	for i := range arr {
		fmt.Println(i)
	}
}
