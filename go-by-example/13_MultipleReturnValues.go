package main

import "fmt"

func vals() (int, int, int) {
	return 3, 5, 7
}

func main() {
	// Here we use the 2 different return values from
	// the call with multiple assignment.
	a, b, c := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// If you only want a subset of the returned values,
	// use the blank identifier _.
	_, d, _ := vals()
	fmt.Println(d)

	_, _, e := vals()
	fmt.Println(e)
}
