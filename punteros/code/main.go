package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	c         *contact
	a         address
}

type contact struct {
	email string
	a     *address
}

type address struct {
	numero int
}

func main() {
	bito := person{
		firstName: "Bito",
		lastName:  "Logic",
	}

	if bito == (person{}) {
		bito.updateName("Mate Cocido")

	}

	bito.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("firstNme:  %+v \n", p.firstName)
	fmt.Printf("LastName:  %+v \n", p.lastName)
}
