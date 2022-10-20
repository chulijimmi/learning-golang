package main

import (
	"testing"
)

func TestSumInts(t *testing.T) {
	n := make(map[string]int64)
	n["a"] = 1
	sum := SumInts(n)
	if sum != 1 {
		t.Fatalf("Error sum %v", sum)
	}
}
