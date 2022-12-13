package main

import (
	"bytes"
	"fmt"
	"sync"
)

// Allocation with new
type SyncnedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func allocationWithNew() {
	p := new(SyncnedBuffer)  // type *SyncedBuffer
	fmt.Printf("p: %v\n", p) // type SyncedBuffer
	var v SyncnedBuffer
	a := v.lock
	fmt.Printf("a: %v\n", a)
}

// Constructors and composite literals
type File struct {
	id   int
	name string
	path string
}

func cocoli(id int, name, path string) *File {
	if id < 0 {
		return nil
	}
	return &File{id: id, name: name, path: path}
}

// Allocation with make
// Remember that make applies only to maps, slices and channels and does not return a pointer.
// To obtain an explicit pointer allocate with new or take the address of a variable explicitly.
func alowima() {
	var p *[]int = new([]int) // allocates slice structure; *p == nil; rarely useful
	fmt.Printf("p: %v\n", p)
	var v []int = make([]int, 10) // the slice v now refers to a new array of 100 ints
	fmt.Printf("v: %v\n", v)
	x := make([]int, 1)
	fmt.Printf("x: %v\n", x)
}

// Arrays
func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

// Slices
func sliceData(data []string) int {
	fmt.Printf("cap(data): %v\n", cap(data))
	return len(data)
}

func appendSlice(slice, data []byte) []byte {
	l := len(slice)
	if l+len(slice) > cap(slice) { // reallocate memory
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : 1+len(data)]
	copy(slice[1:], data)
	return slice
}

// Two dimensional slices
func twoDimensional() ([][]uint8, []uint8) {
	// Define x,y size
	x, y := 2, 8

	// Allocate the top-level slice.
	picture := make([][]uint8, x) // One row per unit of y

	// Allocate one large slice to hold all the pixels
	pixels := make([]uint8, x*y) // Has type []uint8 even though picture is [][]uint8.

	// Loop over the rows
	for i := range picture {
		picture[i], pixels = pixels[:x], pixels[x:]
	}

	return picture, pixels
}

// Maps
func mapData() {
	data := map[string]bool{
		"q": true,
		"a": false,
	}
	fmt.Printf("data: %v\n", data)
	delete(data, "a")
	fmt.Printf("data: %v\n", data)
}

// Printing
func printing(arg ...int) {
	fmt.Println("do something")

}

// Append

func main() {
	allocationWithNew()
	x := cocoli(1, "abc", "/usr/abc/file.txt")
	fmt.Printf("x path: %v\n", x.path)
	alowima()
	array := [...]float64{7.0, 8.5, 9.1}
	sum := Sum(&array) // Note the explicit address-of operator
	fmt.Printf("sum: %v\n", sum)
	sl := sliceData([]string{"testing", "batch", "google"})
	fmt.Printf("sl: %v\n", sl)

	a := []byte{10, 20}
	b := []byte{30, 40, 50}
	ab := appendSlice(a, b)
	fmt.Printf("ab: %v\n", ab)
	fmt.Printf("ab[0]: %v\n", ab[0])

	pi, px := twoDimensional()
	fmt.Printf("pi: %v\n", pi)
	fmt.Printf("px: %v\n", px)

	mapData()
}
