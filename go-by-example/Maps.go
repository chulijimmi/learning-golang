// Maps are a convenient and powerful built-in data structure that
// associate values of one type (the key) with values of
// another type (the element or value). The key can be of any type
// for which the equality operator is defined, such as integers,
// floating point and complex numbers, strings, pointers,
// interfaces (as long as the dynamic type supports equality),
// structs and arrays. Slices cannot be used as map keys, because
// equality is not defined on them. Like slices, maps hold references
// to an underlying data structure. If you pass a map to a function
// that changes the contents of the map, the changes will be visible in the caller.
package main

import "fmt"

func main() {
	m := make(map[string]int)
	// Set key value pair
	m["k1"] = 10
	m["k2"] = 20
	fmt.Println(m)
	v1 := m["k1"]
	fmt.Println("v11:", v1)
	fmt.Println("len:", len(m))
	// Remove key k2
	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional second return value when getting
	// a value from a map indicates if the key was
	// present in the map. This can be used to
	// disambiguate between missing keys and keys
	// with zero values like 0 or "". Here we
	// didn’t need the value itself,
	// so we ignored it with the blank identifier _.
	_, prs := m["k2"]
	fmt.Println("prs", prs)

	// You can also declare and initialize a
	// new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// An attempt to fetch a map value with a key that is not
	// present in the map will return the zero value for the
	// type of the entries in the map. For instance, if the
	// map contains integers, looking up a non-existent key
	// will return 0. A set can be implemented as a map with
	// value type bool. Set the map entry to true to put the
	// value in the set, and then test it by simple indexing.
	personAttended := map[string]bool{"john": true, "ben": false}
	if personAttended["john"] {
		fmt.Println("john", "was at the meeting")
	}

}
