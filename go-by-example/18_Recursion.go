// Go supports recursive functions. Here’s a classic example.
package main

import "fmt"

// This fact function calls itself until it reaches
// the base case of fact(0).
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// Closures can also be recursive, but this requires the closure
// to be declared with a typed var explicitly before it’s defined.
func main() {
	// Equal to 5 * 4 * 3 * 2 * 1 * 1
	fmt.Println(fact(5))
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		// Since fib was previously declared in main,
		// Go knows which function to call with fib here.
		println("N", n)
		numA := fib(n - 1)
		numB := fib(n - 2)
		println("Num:", numA, numB)
		return numA + numB
	}

	fmt.Println("res fib(1): ", fib(1))
	fmt.Println("res fib(2): ", fib(2))
	fmt.Println("res fib(3): ", fib(3))
	fmt.Println("res fib(4): ", fib(4))
	fmt.Println("res fib(5): ", fib(5))
	fmt.Println("res fib(6): ", fib(6))
	fmt.Println("res fib(7): ", fib(7))
	fmt.Println("res fib(8): ", fib(8))
	fmt.Println("res fib(9): ", fib(9))
	fmt.Println("res fib(10): ", fib(10))
}
