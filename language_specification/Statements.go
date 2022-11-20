package main

import (
	"errors"
	"fmt"
)

func terminatingStatements() {
	fmt.Println("============ Terminating Statements ============")
}

func emptyStatements() {
	fmt.Println("============ Empty Statements ============")
}

func labeledStatements() {
	fmt.Println("============ Labeled Statements ============")
}

func expressionStatements() {
	fmt.Println("============ Expression Statements ============")
}

func sendStatements() {
	fmt.Println("============ Send Statements ============")
	x := make(chan int, 1)
	x <- 3
	fmt.Printf("x: %v\n", x)
}

func incdecStatements() {
	fmt.Println("============ IncDec Statements ============")
	x := 1
	x++
	x += 1
	fmt.Printf("x increment: %v\n", x)
	x--
	x -= 1
	fmt.Printf("x decrement: %v\n", x)
}

func assignmentStatements() {
	fmt.Println("============ Assignment Statements ============")
	// exchange a and b
	a := 1
	b := 2
	a, b = b, a
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	// set x[i] of []
	x := []int{1, 2, 3, 4, 5}
	i := 1
	i, x[i] = 1, 10
	fmt.Printf("i: %v\n", i)
	fmt.Printf("x: %v\n", x)

	// sequence set x[i]
	i = 0
	x[i], x[i] = 1, -1
	fmt.Printf("i: %v\n", i)
	fmt.Printf("x: %v\n", x)

	// panic set x[i]
	// x[4], x[5] = 4, 5 // panic because index out of range[5]
	// type Point struct{ x, y int }
	// var p *Point
	// x[2], p.x = 6, 7 // panic: runtime error: invalid memory address or nil pointer dereference

	// set after loop
	i = 2
	x = []int{3, 5, 7}

	// assigntment
	for i, x[i] = range x { // set i, x[2] = 0, x[0]
		break
	}
	fmt.Printf("i: %v\n", i)
	fmt.Printf("x: %v\n", x)
	fmt.Println("After loop 1", i, x)
	// after this loop, i == 0 and x == []int{3, 5, 3}
	for i, x[i] = range x { // set i, x[0], x[1], x[2] = 0, x[2], x[1], x[0]
		fmt.Printf("x[i]: %v\n", i, x[i])
	}
	// after this loop, i == 2 and x == []int{5, 3, 3}
	fmt.Println("After loop 2", i, x)
	fmt.Printf("i: %v\n", i)
	fmt.Printf("x: %v\n", x)

	// For loop
	for i, v := range x {
		fmt.Printf("index as i: %v\n", i)
		fmt.Printf("value as v: %v\n", v)
		x[i] = v
	}
	fmt.Println("After loop 3", i, x)
	fmt.Printf("i: %v\n", i)
	fmt.Printf("x: %v\n", x)
}

func ifStatements() int {
	fmt.Println("============ If Statements ============")
	fx := func() int { return 1 }
	y := 2
	z := 3
	if x := fx(); x < y {
		return x
	} else if x > z {
		return z
	} else {
		return y
	}
}

func switchStatements() {
	fmt.Println("============ Switch Statements ============")
	var x interface{}
	x = 1
	switch x.(type) {
	case nil:
		fmt.Println("x is nil")
	case int:
		fmt.Println("x is int")
	default:
		fmt.Println("don't know the type")
	}
}

func forStatements() {
	fmt.Println("============ For Statements ============")

	// For statements with single condition
	a := 1
	b := 2

	for a < b {
		fmt.Printf("a: %v\n", a)
		a *= 2
	}

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	// For statements with for clause
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
	}

	// For statements with range clause
	var testdata *struct {
		a *[7]int
	}
	f := func(i int) int {
		fmt.Printf("i: %v\n", i)
		return i * 10
	}

	for i, _ := range testdata.a {
		f(i)
	}

	var key string
	var val interface{} // element type of m is assignable to val
	m := map[string]int{"mon": 0, "tue": 1, "wed": 2, "thu": 3, "fri": 4, "sat": 5, "sun": 6}
	for key, val = range m {
		fmt.Printf("key: %v\n", key)
		fmt.Printf("val: %v\n", val)
	}
}

func goStatements() {
	fmt.Println("============ Go Statements ============")
	server := func() int {
		fmt.Println("server run")
		return 1
	}

	server()

	// go server()
	// go list.Sort()  // run list.Sort concurrently; don't wait for it.
}

func selectStatements() {
	fmt.Println("============ Select Statements ============")
	var a []int
	fmt.Printf("a: %v\n", a)
	var c, c1, c2, c3, c4 chan int
	fmt.Printf("c: %v\n", c)
	fmt.Printf("c1: %v\n", c1)
	fmt.Printf("c2: %v\n", c2)
	fmt.Printf("c3: %v\n", c3)
	fmt.Printf("c4: %v\n", c4)

	var i1, i2 int
	fmt.Printf("i1: %v\n", i1)
	fmt.Printf("i2: %v\n", i2)

	i1 = 1
	i2 = 2
	// select {
	// case i1 = <-c1:
	// 	print("received", i1, " from c1\n")
	// case c2 <- i2:
	// 	print("sent ", i2, " to c2\n")
	// case i3, ok := <-c3:
	// 	if ok {
	// 		fmt.Printf("ok: %v\n", ok)
	// 	} else {
	// 		fmt.Printf("i3: %v\n", i3)
	// 	}
	// case t := <-c4:
	// 	fmt.Printf("t: %v\n", t)
	// default:
	// 	println("no communication")
	// }

	// for {
	// 	select {
	// 	case c <- 0:
	// 		fmt.Println("Channel 0", c)
	// 	case c <- 1:
	// 		fmt.Println("Channel 1", c)
	// 	}
	// }

	// select {} // block forever
}

func returnStatements() {
	fmt.Println("============ Return Statements ============")
	// return ExpressionList
	check := func() bool {
		return true
	}

	a := check()
	fmt.Printf("a: %v\n", a)

	// empty return
	destroy := func() {
		fmt.Println("destroy")
		return
	}
	destroy()

	// restriction
	sum := func(a, b int) (int, error) {
		if a < 1 {
			return -1, errors.New("a less than 1")
		}
		return a + b, nil
	}
	b, bError := sum(-1, 2)
	fmt.Printf("b: %v\n", b)
	if bError != nil {
		fmt.Printf("bError: %v\n", bError)
	}
}

func breakStatements() {
	fmt.Println("============ Break Statements ============")
	// A "break" statement terminates execution of the innermost
	// "for", "switch", or "select" statement within the same function.
	i := 0
	for {
		i++
		if i > 2 {
			fmt.Printf("i: %v\n", i)
			break
		}
	}
}

func continueStatements() {
	fmt.Println("============ Continue Statements ============")
	for i := 0; i < 10; i++ {
		if i > 0 {
			fmt.Printf("i continue: %v\n", i)
			continue
		}
		fmt.Printf("i break: %v\n", i)
	}
}

func gotoStatements() {
	fmt.Println("============ Goto Statements ============")
	n := 10
L1:
	fmt.Println("go to L1")

	if n%2 == 1 {
		goto L1
	}
}

func falthroughStatements() {
	fmt.Println("============ Falthrough Statements ============")
	// A "fallthrough" statement transfers control to the first
	// statement of the next case clause in an expression "switch" statement. It

	f := func() {
		gfg := "Geek"
		// Use switch on the day variable.
		switch {
		case gfg == "Geek":
			fmt.Println("Geek")
			fallthrough
		case gfg == "For":
			fmt.Println("For")
			fallthrough
		case gfg == "Geeks":
			fmt.Println("Geeks")
		}
	}

	f()
}

// A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns,
// either because the surrounding function executed a return statement, reached the end of its function body,
// or because the corresponding goroutine is panicking.

func deferStatements() {
	fmt.Println("============ Defer Statements ============")
	// prints 3 2 1 0 before surrounding function returns
	for i := 0; i <= 3; i++ {
		defer fmt.Print(i)
	}

	// f returns 42
	f := func() (result int) {
		defer func() {
			// result is accessed after it was set to 6 by the return statement
			result *= 7
		}()
		return 6
	}

	fmt.Printf("f(): %v\n", f())
}

func main() {
	fmt.Println("============ Statements ============")
	terminatingStatements()
	emptyStatements()
	labeledStatements()
	expressionStatements()
	sendStatements()
	incdecStatements()
	assignmentStatements()
	i := ifStatements()
	fmt.Printf("ifStatements: %v\n", i)
	switchStatements()
	forStatements()
	goStatements()
	selectStatements()
	returnStatements()
	breakStatements()
	continueStatements()
	gotoStatements()
	falthroughStatements()
	deferStatements()
}
