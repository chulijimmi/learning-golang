package main

import (
	"fmt"
)

// If
func doif(a, b int) (bool, []int) {
	if a < b {
		return true, nil
	}

	// if accept an initialization statement
	if err := []int{1, 2}; err != nil {
		return false, err
	}

	return false, nil
}

// Redeclaration and reassignment
// Example:
// f, err := os.Open(name)

// For
func doFor() {
	// Like a C for
	// for init; condition; post { }

	// Like a C while
	// for condition { }

	// Like a C for(;;)
	// for { }

	// Short declarations make it easy to declare the index variable right in the loop.
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("sum: %v\n", sum)

	// If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.
	arr := []int{1, 2, 3}
	newArr := []int{4, 5, 6}
	for key, value := range arr {
		newArr[key] = value
	}
	fmt.Printf("newArr: %v\n", newArr)

	// For string
	for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}

	// no comma operator multiple variable
	str := [6]byte{65, 66, 67, 226, 130, 172}
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	fmt.Printf("str: %v\n", str)
}

// Switch
func doSwitch(src []byte, size int) {
Loop:
	for n := 0; n < 10; n += size {
		fmt.Printf("n: %v\n", n)
		switch {
		case src[n] == 10:
			break
		case src[n] == 100:
			fmt.Println("Break Loop !")
			break Loop
		}
	}
}

// Type switch
func doTypeSwitch() {
	var t interface{}
	t = func() bool { return true }
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected t: %v\n", t)
	case bool:
		fmt.Printf("bool expected t: %v\n", t)
	}
}

func main() {
	// If statement
	a, _ := doif(2, 5)
	fmt.Printf("a: %v\n", a)

	doFor()

	src := []byte{10, 20, 30, 40, 50, 60, 100, 80, 90, 100}
	doSwitch(src, 1)

	doTypeSwitch()
}
