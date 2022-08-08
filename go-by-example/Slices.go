// Slices are a key data type in Go,
// giving a more powerful interface to sequences than arrays.

// Unlike arrays, slices are typed only by the elements they
// contain (not the number of elements). To create an empty
// slice with non-zero length, use the builtin make.
// Here we make a slice of strings of length 3 (initially zero-valued).

// Reference: https://go.dev/blog/slices-intro

package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	setArr := func(index int) {
		switch index {
		case 1:
			s[1] = "B"
		case 2:
			s[2] = "C"
		default:
			s[0] = "A"
		}
	}
	fmt.Println("set s array:")
	setArr(1)
	setArr(2)
	setArr(3)
	fmt.Println(s)
	fmt.Println("get:s[0]", s[0])

	// In addition to these basic operations,
	// slices support several more that make them richer than arrays.
	// One is the builtin append, which returns a slice containing
	// one or more new values. Note that we need to accept
	// a return value from append as we may get a new slice value.
	s = append(s, "D")
	s = append(s, "E", "F")
	fmt.Println("append:", s)

	// Slices can also be copy’d.
	// Here we create an empty slice c of the same
	// length as s and copy into c from s.
	c := make([]string, len(s))
	fmt.Println("len c", c)
	copy(c, s)
	fmt.Println("copy:", c)

	// Slices support a “slice” operator with the syntax slice[low:high].
	// For example, this gets a slice of the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("slice s by 1", l)

	//This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("slice exclude 5", l)

	//And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("slice from including [2]", l)

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary,
	// unlike with multi-dimensional arrays.
	twoDimmension := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDimmension[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDimmension[i][j] = i + j
		}
	}
	fmt.Println("2d array:", twoDimmension)
}
