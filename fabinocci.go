package main

import (
	"fmt"
)
func main() {
	x := 0
	y := 1
	c := 0
	n := 5

	for c < n {
		z := x + y
		x = y
		y = z
		c++
	}
	fmt.Println(x)
	// fmt.Println(x, "\n", y)
	// for i := 3; i <= s; i++ {
	// 	fmt.Println(z)
	// 	x = y
	// 	y = z
	// 	z = x + y
	// }

}
