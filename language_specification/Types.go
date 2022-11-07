package main

import (
	"fmt"
)

func booleanType() {
	fmt.Println("========== Boolean ==========")
	type question struct {
		isAnswer bool
	}
	var a = *&question{}
	a.isAnswer = false
	fmt.Println("isAnswer: ", a.isAnswer)
}

func numericType() {
	fmt.Println("========== Numberic ==========")
	var a uint8 // unsigned  8-bit integers
	a = 255     // 0 to 255
	fmt.Println("uint8: ", a)
	var b uint16 // unsigned 16-bit integers
	b = 65535    // 0 to 65535
	fmt.Println("uint16: ", b)
	var c uint32   // unsigned 32-bit integers
	c = 4294967295 // 0 to 4294967295
	fmt.Println("uint32: ", c)
	var d uint64             // unsigned 64-bit integers
	d = 18446744073709551615 // 0 to 18446744073709551615
	fmt.Println("uint64: ", d)
	var e int8 // signed  8-bit integers
	e = -128   // -128 to 127
	fmt.Println("int8: ", e)
	var f int16 // signed 16-bit integers
	f = -32768  // -32768 to 32767
	fmt.Println("int16: ", f)
	var g int32     // signed 32-bit integers
	g = -2147483648 // -2147483648 to 2147483647
	fmt.Println("int32: ", g)
	var h int64              // signed 64-bit integers
	h = -9223372036854775808 // -9223372036854775808 to 9223372036854775807
	fmt.Println("int64: ", h)
	var i float32 // IEEE-754 32-bit floating-point numbers
	i = 2.34
	fmt.Println("float32: ", i)
	var j float64 // IEEE-754 64-bit floating-point numbers
	j = 1920.012910
	fmt.Println("float64: ", j)
	var k complex64 // complex numbers with float32 real and imaginary parts
	k = -123.23
	fmt.Println("complex64: ", k)
	var l complex128 // complex numbers with float64 real and imaginary parts
	l = -1009.203
	fmt.Println("complex128: ", l)
	var m byte // alias for uint8
	m = 255
	fmt.Println("byte: ", m)
	var n rune // alias for int32
	n = 1212
	fmt.Println("rune: ", n)
	var o uint // either 32 or 64 bits
	o = 12
	fmt.Println("uint: ", o)
	var q int // same size as uint
	q = 123
	fmt.Println("int: ", q)
	var r uintptr // an unsigned integer large enough to store the uninterpreted bits of a pointer value
	r = 123
	fmt.Println("uintptr: ", r)
}

func arrayType() {
	fmt.Println("========== Array Type ==========")
	var a [32]byte
	a[0] = 1
	fmt.Println("arr a: ", a)
	var b [2 * 1]struct{ x, y int32 }
	b[0].x = 12
	fmt.Println("arr b: ", b)

	// https://www.tutorialspoint.com/go/go_array_of_pointers.htm
	const MAX int = 3
	c := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &c[i] /* assign the address of integer. */
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *ptr[i])
	}

	var d [2][3]int
	d[0][1] = 2
	fmt.Println("arr d: ", d)
	var e [2]([2]([2]float64)) // [2][2][2]float64
	e[0][1][1] = 56.32
	fmt.Println("arr e: ", e)
}

func sliceType() {
	fmt.Println("========== Slice Type ==========")
	a := make([]int, 1, 10)
	a = append(a, 1)
	fmt.Println("slice a: ", a)
}

func structType() {
	// Empty struct
	type a struct{}
	type b struct {
		x, y int
		u    float32
		_    float32
		A    *[]int
		F    func()
	}
	type c struct {
		x, y float32 ""
		name string
		_    [4]byte
	}
	type d struct {
		microsec  uint64 `protobuf:"1"`
		serverIP6 uint64 `protobuf:"2"`
	}
}

func pointerType() {
	fmt.Println("========== Pointer Type ==========")
	type P struct {
		name string
	}
	a := P{"foo"}
	b := *&P{"bar"}
	println("a: ", a.name)
	fmt.Printf("b: %v\n", b.name)
}

func a() {
	fmt.Println("empty return function")
}

func b(x int) int {
	fmt.Println("int return function", x)
	return x
}

func c(a, _ int, z float32) bool {
	fmt.Println("bool return function", a, z)
	return true
}

func d(a, b int, z float32) bool {
	fmt.Println("bool return function", a, b, z)
	return false
}

func e(prefix string, values ...int) {
	fmt.Println("empty return function", prefix, values)
}

func f(a, b int, z float64, opt ...interface{}) (success bool) {
	fmt.Println("arg bool return function", a, b, z, opt)
	return success
}

func g(int, int, float64) (float64, *[]int) {
	fmt.Println("float64, int return function")
	return 121.243, nil
}

type T struct {
	name string
}

func h(n int) func(p *T) {
	fmt.Println("N", n)
	a := T{name: "foo"}
	fmt.Println(a)
	return nil
}

func mapType() {
	fmt.Println("========== Map Type ==========")
	a := make(map[string]int)
	a["test"] = 1
	fmt.Println("Map a: ", a)
}

func chanType() {
	fmt.Println("========== Chan Type ==========")
	a := make(chan int, 100)
	fmt.Printf("a: %v\n", a)
}

func coreType() {
	type Celsius float32
	type Kelvin float32

	type a interface{ int }
	type b interface{ Celsius | Kelvin }
}

// Array type consist of:
// - Boolean types
// - Numeric types
// - String types
// - Array types
// - Slice types
// - Struct types
// - Pointer types
// - Function types
// - Interface types
// - Map types
// - Channel types

func main() {
	booleanType()
	numericType()
	arrayType()
	sliceType()
	structType()
	pointerType()
	mapType()
	chanType()
	coreType()
}
