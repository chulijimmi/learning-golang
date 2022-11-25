package main

import "fmt"

// Send the sequence 2, 3, 4 ... to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // send i to channel ch
	}
}

// Copy the values from channel 'src' to channel 'dst'
// removing those divisible by 'prime'
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dst <- i // send i to channel dst
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func sieve() {
	ch := make(chan int)
	fmt.Printf("ch: %v\n", ch)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Printf("prime: %v\n", prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	fmt.Printf("end ch: %v\n", ch)
}

func main() {
	sieve()
}
