package main

import "fmt"

func hurdleRace(k int, height []int) int {
	max := 0
	for _, val := range height {
		if val > k && val > max {
			max = val
		}
	}

	if max == 0 {
		return 0
	}
	return max - k
}

func main() {
	k := 4
	height := []int{1, 6, 3, 5, 2}
	res := hurdleRace(k, height)

	fmt.Println(res)
}
