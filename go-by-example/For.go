// for is Go’s only looping construct. Here are some basic types of for loops.

package main

import "fmt"

func main() {
	i := 1
	count := 5
	// The most basic type, with a single condition.
	fmt.Print("result for < 3")
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Print("----------------------")

	// A classic initial/condition/after for loop.
	for j := 0; j < count; j++ {
		fmt.Println("j:", j)
	}
	fmt.Print("----------------------")

	// You can also continue to the next iteration of the loop.
	fmt.Print("result condition")
	for n := 0; n <= count; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	fmt.Print("----------------------")

	// for without a condition will loop repeatedly until you
	// break out of the loop or return from the enclosing function.
	fmt.Print("result stop break")
	for {
		fmt.Println(i)
		break
	}

	// If you're looping over an array, slice, string, or map,
	// or reading from a channel, a range clause can manage the loop.
	arr := map[int]string{1: "abce", 2: "def", 3: "ghi"}
	for key, value := range arr {
		println("key, value", key, value)
	}

}
