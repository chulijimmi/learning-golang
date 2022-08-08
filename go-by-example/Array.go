// In Go, an array is a numbered sequence of elements of a specific length.
package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("decimal", b)

	var twoDimension [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimension[i][j] = i + j
		}
	}
	fmt.Println(twoDimension)
	fmt.Println("-------------------------------------")

	//runtime error: index out of range [2] with length 2
	var threeDimension [2][3][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 4; k++ {
				threeDimension[i][j][k] = i + j + k
			}
		}
	}
	fmt.Println(threeDimension)
}
