package main

import (
	"fmt"
)

func BifClose() {
	fmt.Println("====== Close ======")
}

func BifLengthAndCapacity() {
	fmt.Println("====== Length And Capacity ======")
	a := "abcde"
	fmt.Printf("a: %v\n", len(a))
	s := []int{1, 2}
	fmt.Printf("s: %v\n", cap(s))
}

func BifAllocation() {
	fmt.Println("====== Allocation ======")
	type S struct {
		a string
		b int
	}
	x := new(S)
	x.a = "test"
	x.b = 11
	fmt.Printf("x: %v\n", x)
}

func BifMakingSlice() {
	fmt.Println("====== Making slices, maps and channels ======")
	s := make([]int, 10, 100)
	fmt.Printf("s: %v\n", s)
	c := make(chan int, 10)
	fmt.Printf("c: %v\n", c)

}

func BifAppending() {
	fmt.Println("====== Appending to and copying slices ======")
	s0 := []int{0, 0}

	// append a single element
	s1 := append(s0, 2)
	fmt.Printf("s1: %v\n", s1)

	// append multiple elements
	s2 := append(s1, 3, 5, 7)
	fmt.Printf("s2: %v\n", s2)

	// append a slice
	s3 := append(s2, s0...)
	fmt.Printf("s3: %v\n", s3)

	// append overlapping slice
	s4 := append(s3[3:6], s3[2:]...)
	fmt.Printf("s4: %v\n", s4)

	var t []interface{}
	t = append(t, 42, 3.1415, "foo")
	fmt.Printf("t: %v\n", t)

	// append string contents
	var b []byte
	b = append(b, "bar"...)
	fmt.Printf("b: %v\n", b)

	// copy copies slice elements
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("a: %v\n", a)

	var ab = make([]int, 6)
	fmt.Printf("ab: %v\n", ab)

	var abc = make([]byte, 5)
	fmt.Printf("abc: %v\n", abc)

	n1 := copy(ab, a[:]) // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("n1: %v\n", n1)
	fmt.Printf("ab: %v\n", ab)

	// n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
	n2 := copy(ab, ab[2:])
	fmt.Printf("n2: %v\n", n2)

	n3 := copy(abc, "Hello, world!")
	fmt.Printf("n3: %v\n", n3)
}

func BifDeletion() {
	fmt.Println("====== Deletion of map elements ======")
	a := make(map[string]string)
	a["name"] = "shoes"
	a["brand"] = "zara"
	fmt.Printf("a: %v\n", a)
	delete(a, "brand")
	fmt.Printf("a: %v\n", a)
}

func BifManipulating() {
	fmt.Println("====== Manipulating complex numbers ======")
	var a = complex(2, -2)
	fmt.Printf("a: %v\n", a)
	const b = complex(1.0, -1.4)
	fmt.Printf("b: %v\n", b)
}

func BifHandlingPanic() {
	fmt.Println("====== Handling Panics ======")
}

func BifBootstrapping() {
	fmt.Println("====== Bootstrapping ======")
}

func main() {
	BifClose()
	BifLengthAndCapacity()
	BifAllocation()
	BifMakingSlice()
	BifAppending()
	BifDeletion()
	BifManipulating()
	BifHandlingPanic()
	BifBootstrapping()
}
