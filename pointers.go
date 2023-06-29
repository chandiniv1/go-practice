package main

import "fmt"

//result without using pointers will be 0,0
/*func increase(i int) {
	i = i + 1
}
func main() {
	i := 0
	fmt.Println(i)
	increase(i)
	fmt.Println(i)
}*/

func increase(i *int) {
	*i = *i + 1
}
func main() {
	i := 0
	fmt.Println(i)
	increase(&i)
	fmt.Println(i)
}
