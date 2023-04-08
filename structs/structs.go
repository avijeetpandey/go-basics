package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
}

type personWithContact struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// firstWay of defining a struct
	alex := person{"Alex", "Anderson"}

	// another way of creating a struct
	amy := person{firstName: "Amy", lastName: "Miles"}

	fmt.Println(alex)
	fmt.Println(amy)

	//  struct contains zero values if left un-assigned
	var aiko person
	fmt.Println(aiko)

	// creating person with contactInfo
	jim := personWithContact{
		firstName: "Jim",
		lastName:  "Kick",
		contact: contactInfo{
			email: "test@g.cc",
			zip:   122002,
		},
	}

	jim.print()
}

func (p personWithContact) print() {
	fmt.Printf("Full Name => %v", p.firstName+" "+p.lastName)
	fmt.Printf("Email => %v", p.contact.email)
	fmt.Printf("Zip Code => %d", p.contact.zip)
}
