package main

import "fmt"

// Here’s a function that takes two ints and returns their sum as an int.
func addition(a int, b int) int {
	return a + b
}

// Go requires explicit returns, i.e. it won’t
// automatically return the value of the last expression.
// When you have multiple consecutive parameters of the same type,
// you may omit the type name for the like-typed parameters up
// to the final parameter that declares the type.
func multipleAdditionSameType(a, b, c int) int {
	return a + b + c
}

// Call a function just as you’d expect, with name(args).
func main() {
	res := addition(1, 2)
	fmt.Println("1+2 eq: ", res)

	res = multipleAdditionSameType(1, 2, 3)
	fmt.Println("1+2+3 eq:", res)

	// command-line-arguments
	// .\Functions.go:27:22: cannot use "a" (untyped string constant)
	// as int value in argument to addition
	// fail := addition(1, "a")
	// fmt.Println("Fail", fail)
}
