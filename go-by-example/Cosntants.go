// A const statement can appear anywhere a var statement can.

package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	const n = 500000000
	const d = 32e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
