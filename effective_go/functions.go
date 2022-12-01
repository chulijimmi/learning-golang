package main

import "fmt"

// Multiple return values
func multiple(a int, b bool) (int, bool) {
	return a, b
}

// Named result parameters
// Must clear return
func read(a string, b int) (i int, j string) {
	i = i + b
	a = "_" + j
	return
}

// Defer
// Go's defer statement schedules a function call
// (the deferred function) to be run immediately before
// the function executing the defer returns.
func deferProc() {
	closeMysql := func() {
		fmt.Println("Closed mysql")
	}
	defer closeMysql() // This will run after end of function
	// run query
	for i := 0; i < 5; i++ {
		fmt.Println("run query", i)
	}
	fmt.Println("end of function")
}

func main() {
	multiple(1, true)
	read("abc", 3)
	deferProc()
}
