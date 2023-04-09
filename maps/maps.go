package main

import "fmt"

// Maps are basically the store of the key and value of pairs
func main() {

	// declaring and assigning values at the same time
	colors := map[string]string{
		"red":    "#ff0000",
		"yellow": "#ff9x0d",
	}

	// declaring an empty map
	var emptyColorsMap map[string]string

	emptyColorsMap = make(map[string]string)
	emptyColorsMap["yellow"] = "jlsd"

	// making a map
	madeColorsMap := make(map[string]string)
	madeColorsMap["while"] = "white"

	fmt.Println(emptyColorsMap)
	fmt.Println(madeColorsMap)
	fmt.Println(colors)
}
