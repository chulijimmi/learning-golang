package main

import (
	"fmt"
)

func main() {
	{
		a := 1
		fmt.Println("a: ", a)

		{
			b := 2
			fmt.Println("b: ", b)
		}
	}
}
