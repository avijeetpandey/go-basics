package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// firstWay of defining a struct
	alex := person{"Alex", "Anderson"}

	// another way of creating a struct
	amy := person{firstName: "Amy", lastName: "Miles"}

	fmt.Println(alex)
	fmt.Println(amy)
}
