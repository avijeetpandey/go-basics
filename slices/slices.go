// A simple program to understand slices
package main

import "fmt"

/**
A slice is basically a data structure to store records of similar type. It can grow/shrink at will
**/

func main() {
	// a simple slice of 3 elements
	socialMediaSites := []string{"FB", "Instagram", "Twitter"}

	// appending elements in a slice
	socialMediaSites = append(socialMediaSites, "SnapChat")

	// iterating over the slices
	for index, site := range socialMediaSites {
		fmt.Println(index, site)
	}
}
