package main

import (
	"fmt"
	"time"
)

func announceManager(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println(message + delay.String())
}

func announceVp(message string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println("Ping VP")
	}
}

func announceDirector(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println(message)
}

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func main() {
	fmt.Println("================== Share by communicating ==================")
	// Go encourages a different approach in which shared values are passed around on
	// channels and, in fact, never actively shared by separate threads of execution
	fmt.Println("================== Goroutines ==================")
	go announceDirector("Ping Director", 1)
	go announceVp("Ping VP", 2)
	go announceManager("Ping Manager", 1)
	time.Sleep(10)
	fmt.Println("goroutine done")

	list := []int{1, 2, 3, 4, 5}
	c := make(chan int) // allocate channel
	// Start the sort in a goroutine; when it completes, signal on the channel.
	go func() {
		fmt.Printf("list: %v\n", list)
		c <- 1
	}()
	time.Sleep(10)
	<-c
	fmt.Println("================== Channels ==================")
	fmt.Println("================== Channels of channels ==================")
	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	// send request
	close(request.resultChan)
	// wait for response
	fmt.Printf("answer: %d\n", <-request.resultChan)
	fmt.Println("================== Parallelization ==================")
	fmt.Println("================== A leaky buffer ==================")
}
