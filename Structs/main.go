package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type newPerson struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// var tejas person
	// fmt.Printf("%+v", tejas)

	tejas := person{
		firstName: "Tejas",
		lastName:  "Namjoshi",
		contact: contactInfo{
			email: "tejasnamjoshi9@gmail.com",
			zip:   411009,
		},
	}
	// fmt.Printf("%+v", tejas)
	// tejas.print()
	// tejas.updateName("Tejas New")
	// tejas.print()

	// tejasPointer := &tejas
	// tejasPointer.updateNameReal("Tejas Real New")

	tejas.updateNameRealShortcut("Tejas New real shortcut")
	tejas.print()

	// newTejas := newPerson{
	// 	firstName: "Tejas",
	// 	lastName:  "Namjoshi",
	// 	contactInfo: contactInfo{
	// 		email: "tejasnamjoshi9@gmail.com",
	// 		zip:   411009,
	// 	},
	// }
	// fmt.Printf("%+v", newTejas)
	// tejas.contact.email = "tnamjoshi@salucro.com"
	// tejas.firstName = "New Tejas"
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) updateNameReal(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p *person) updateNameRealShortcut(newFirstName string) {
	p.firstName = newFirstName
}
