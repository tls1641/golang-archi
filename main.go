package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("from a person - this is my name", p.first)
}

// any type that has the methods specified by an interface, is also of the interface type
// hey baby if you have these methods, then you're my type

type human interface {
	speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	fmt.Printf("%T\n", p1)

	// in go a value can be of more than one Type
	// in this example, p1 is both type person, and human.
	var x human
	x = p1
	fmt.Printf("%T\n", x)
}
