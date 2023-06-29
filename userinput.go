package main

import (
	"fmt"
)

func main() {
	var i, j, z string
	fmt.Println("enter a number:")
	fmt.Scanln(&i,&j)
	fmt.Println("number is:", i, j)
	fmt.Println("enter a number:")
	fmt.Scanf("%d", &z)
	fmt.Println(z)
}
