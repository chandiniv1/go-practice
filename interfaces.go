package main

import "fmt"

type Rect struct {
	len int
	bre int
}

type Shape interface {
	area()
	perimeter() int
}

func (r Rect) area() {
	fmt.Println(r.len * r.bre)
}
func (r Rect) perimeter() int {
	return 2 * (r.len + r.bre)
}
func main() {
	var shape Shape
	shape = Rect{10, 20}
	shape.area()
	fmt.Println(shape.perimeter())
}
