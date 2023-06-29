package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Print() {
	fmt.Println(p.name, p.age)
}

func main() {
	person := Person{name: "chandini", age: 22}
	var p Person
	p.Print()
	person.Print()
	fmt.Println(person)
}
