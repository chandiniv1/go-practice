package main

import "fmt"

func main() {
	sli := []int{1, 2, 3, 4, 5}
	fmt.Println(sli)

	mysli := sli[:3]
	fmt.Println(mysli)
	mysli1 := sli[:]
	fmt.Println(mysli1)

}
