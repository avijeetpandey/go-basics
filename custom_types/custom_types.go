// A simple program to demonstrate create and
// working with new types since Go is not an object oriented language
package main

import "fmt"

// create a new type `fruits` , which is essentially a slice of string
type fruits []string

func (fr fruits) print() {
	for index, fruit := range fr {
		fmt.Println(index, fruit)
	}
}

func main() {
	fruits := fruits{"Apple", "Payapa", "Grapes", "Honey"}
	fruits = append(fruits, "Banana")
	fruits.print()
}
