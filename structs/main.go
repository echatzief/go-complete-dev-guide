package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	lastName  string
	firstName string
	contact   ContactInfo
}

func (p Person) print() {
	fmt.Println(p)
}

func (p *Person) updateName(firstName string) {
	(*p).firstName = firstName
}

func main() {
	jim := Person{firstName: "Jim", lastName: "Jones", contact: ContactInfo{email: "jimjones@gmail.com", zipCode: 38321}}
	jim.print()
	jimPointer := &jim
	jimPointer.updateName("Jack")
	jim.print()
	jimPointer.print()
}
