package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Jimbo",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 30318,
		},
	}

	jim.updateName("jimmy")
	jim.print()
}

// By adding *person, the type can be either the pointer OR the actual value and will self resolve
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person) print() {
	fmt.Printf("%+v", *p)
}
