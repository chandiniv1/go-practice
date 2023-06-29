package main

import "fmt"

func swap(x , y int) (int, int) {
	return y, x
}

func main() {
	p, q := swap(4, 3)
	fmt.Println(p, q)
	//fmt.Println(swap(4,3))
}

/*func main() {
	a, b := 2, 3
	c := float32(a + b)
	fmt.Println(c)

}*/
