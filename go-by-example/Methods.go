package main

import "fmt"

type rectangle struct {
	width, height int
}

// This area method has a receiver type of *rect.
func (r *rectangle) area() int {
	return r.width * r.height
}

// Methods can be defined for either pointer or
// value receiver types. Here’s an example of a value receiver.
func (r *rectangle) perim() int {
	return 2*r.width + 2*r.height
}

// Let's mutate the struct
func (r *rectangle) mutate() int {
	r.width = r.width * r.width
	r.height = r.height * r.height
	return r.width + r.height
}

func main() {
	r := rectangle{width: 10, height: 5}
	fmt.Println("rectangle:", r)

	// Here we call the 2 methods defined for our struct.
	fmt.Println("area method:", r.area())
	fmt.Println("rectangle:", r)
	fmt.Println("perim method:", r.perim())
	fmt.Println("rectangle:", r)

	// Go automatically handles conversion between values and
	// pointers for method calls. You may want to use a pointer
	// receiver type to avoid copying on method calls or to
	// allow the method to mutate the receiving struct.
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
	fmt.Println("mutate:", rp.mutate())
	fmt.Println("struct:", r)
}
