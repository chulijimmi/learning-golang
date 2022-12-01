package main

import "fmt"

func main() {
	// Formatting
	type T struct {
		name  string
		value int
	}

	fmt.Printf("T: %v\n", T)

	// Commentary
	// Example Block comment /* */
	/*
	* Put documentation here
	 */
	a := "hello word"
	fmt.Printf("a: %v\n", a)

	// Put documentation
	b := "hi, user"
	fmt.Printf("b: %v\n", b)
}
