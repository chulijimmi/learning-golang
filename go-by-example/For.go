package main

import "fmt"

func main() {
	i := 1
	count := 5
	fmt.Print("result for < 3")
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Print("----------------------")
	fmt.Print("result condition")
	for n := 0; n <= count; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	fmt.Print("----------------------")
	fmt.Print("result stop break")
	for {
		fmt.Println(i)
		break
	}
}
