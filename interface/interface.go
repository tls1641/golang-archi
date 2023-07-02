package interfaces

import "fmt"

type person struct {
	first string
}

func (p person) speak() {
	fmt.Println("from a person", p.first)
}

type secretAgent struct {
	person
	ltk bool
}

func (p secretAgent) speak() {
	fmt.Println("from a person", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func RunInterfaces() {
	p1 := person{
		first: "Jenny",
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	saySomething(p1)
	saySomething(sa1)
}
