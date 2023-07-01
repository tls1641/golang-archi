package main

import "fmt"

type human interface {
	speak()
}

// concrete type
type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("from a person - this is my name", p.first)
}

type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I'm a secret agent - this is my name", sa.first)
}

func foo(h human) {
	h.speak()
}

// any type that has the methods specified by an interface, is also of the interface type
// hey baby if you have these methods, then you're my type

func main() {
	p1 := person{
		first: "MissMoneypoenny",
	}

	sa1 := secretAgent{
		person: person{
			first: "Miss moneyPenny",
		},
		ltk: true,
	}

	// in go a value can be of more than one Type
	// in this example, p1 is both type person, and human.
	var x, y human
	x = p1
	y = sa1
	x.speak()
	y.speak()

	foo(x)
	foo(y)
	foo(p1)
	foo(sa1)
}
