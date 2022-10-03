package main

import "fmt"

func permutationEquation(p []int) []int {
	n := make([]int, 0)
	for i := range p {
		n = append(n, p[i]+1)
	}
	fmt.Println(n)

	res := make([]int, 0)
	for j := range n {
		res = append(res, p[j]+1)
	}
	fmt.Println(res)
	return p
}

func main() {
	var input = []int{1, 2, 3, 4, 5}
	result := permutationEquation(input)

	fmt.Println("result", result)
}
