package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	num := 8%4 == 0
	if num {
		fmt.Println("8 devided by 4")
	}

	if x := 9; x < 0 {
		fmt.Println(x, "is negative")
	} else if x < 10 {
		fmt.Println(x, "has 1 digits")
	} else {
		fmt.Println(x, "has multiple digits")
	}
}
