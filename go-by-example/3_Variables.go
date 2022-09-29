package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var declare a as initial
	var a = "initial"
	fmt.Println(a)

	//var declare b,c as int
	var b, c int = 1, 2
	fmt.Println(b, c)

	//var declare d as boolean
	var d = true
	fmt.Println(d)

	//var declare e as integer
	var e int
	fmt.Println(e)

	//shorthand var declaration :=
	f := "apple"
	fmt.Println(f)

	//Print data type
	fmt.Printf("a-type: %s\n", reflect.TypeOf(a))
	fmt.Printf("b-type: %s\n", reflect.TypeOf(a))
	fmt.Printf("d-type: %s\n", reflect.TypeOf(d))
	fmt.Printf("e-type: %s\n", reflect.TypeOf(e))

	//Error: undefined: unknownVariables
	//fmt.Println(unknownVariables)
}
