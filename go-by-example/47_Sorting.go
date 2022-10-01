// Go’s sort package implements sorting for builtins and user-defined types.
// We’ll look at sorting for builtins first.

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sort methods are specific to the builtin type; here’s an example for strings.
	// Note that sorting is in-place, so it changes the given slice and doesn’t return a new one.
	chars := []string{"c", "a", "b"}
	sort.Strings(chars)
	fmt.Println("Strings", chars)

	// An example of sorting ints.
	nums := []int{7, 2, 4}
	sort.Ints(nums)
	fmt.Println("Numbers", nums)

	// We can also use sort to check if a slice is already in sorted order.
	isSortNums := sort.IntsAreSorted(nums)
	fmt.Println("Sorted", isSortNums)
}
