package main

import "fmt"

func constantDeclaration() {
	const Pi float32 = 5.321
	fmt.Println("Pi: ", Pi)

	const zero = 0.0
	fmt.Println("zero: ", zero)
	const (
		size int64 = 1024
		eof        = -1
	)
	fmt.Println("size: ", size)
	fmt.Println("eof: ", eof)
	const a, b, c = 3, 4, "foo"
	fmt.Println("a, b, c", a, b, c)
	const u, v float32 = 0, 3
}

func iotaDeclaration() {
	const (
		c0 = iota
		c1 = iota
		c2 = iota
	)
	fmt.Printf("c0: %v\n", c0)
	fmt.Printf("c1: %v\n", c1)
	fmt.Printf("c2: %v\n", c2)
}

func typeDeclaration() {
	type TreeNode struct {
		left, right *TreeNode
		value       any
	}

	a := *&TreeNode{}
	a.value = 3
	fmt.Println("a.value: ", a.value)
}

type Point3D struct {
	x, y, z float32
}

type Line struct{ p, q Point3D }

func compositeLiterals() {
	fmt.Println("========= Composite Literals =========")
	origin := Point3D{}
	line := Line{origin, Point3D{y: -4, z: 12.3}}
	fmt.Println("line: ", line)

	// Taking the address of a composite literal generates a pointer to a unique variable initialized with the literal's value.
	var ptr *Point3D = &Point3D{y: 1000}
	fmt.Println("ptr: ", ptr)
	fmt.Println("origin: ", origin)
	fmt.Println("line: ", line)

	p1 := &[]int{}
	fmt.Println("p1: ", p1)
	p2 := new([]int)
	fmt.Println("p2: ", p2)
	ptrP2 := &p2
	fmt.Println("ptrp2: ", ptrP2)

	// The notation ... specifies an array length equal to the maximum element index plus one.
	buffer := [10]string{}
	fmt.Println("buffer: ", buffer)
	intSet := [6]int{1, 2, 3, 5}
	fmt.Println("intSet: ", intSet)
	days := [...]string{"Sat", "Sun"}
	fmt.Println("days: ", days)

	// and is shorthand for a slice operation applied to an array
	tmp := [10]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	tmp_a := tmp[0:3]
	fmt.Println("tmp_a", tmp_a)

	// frequencies in Hz for equal-tempered scale (A4 = 440Hz)
	noteFrequency := map[string]float32{
		"C0": 16.35,
		"D0": 18.35,
		"E0": 20.60,
		"F0": 21.83,
		"G0": 24.50,
		"A0": 27.50,
		"B0": 30.87,
	}
	fmt.Println("noteFrequency: ", noteFrequency)
}

func functionLiterals() {
	fmt.Println("============ Function Literal ============")
	f := func(x, y int) int { return x + y }
	res := f(2, 4)
	fmt.Println("res f: ", res)

}

func main() {
	constantDeclaration()
	iotaDeclaration()
	typeDeclaration()
	compositeLiterals()
	functionLiterals()
}
