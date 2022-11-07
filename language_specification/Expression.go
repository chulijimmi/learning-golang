package main

import (
	"fmt"
)

type T0 struct {
	x int
}

func (*T0) M0() {}

type T1 struct {
	y int
}

func (T1) M1() {}

type T2 struct {
	z int
	T1
	*T0
}

func (*T2) M2() {}

type Q *T2

var t T2
var p *T2
var q Q = p

func expSelector() {
	fmt.Println("====== Selector ======")
	a := t.z
	fmt.Println("a: ", a)
	b := t.T1.y
	fmt.Println("b: ", b)
}

type TA struct {
	name string
}

func (t TA) Mv(b string) string {
	return "finish"
}

func methodExpression() {
	fmt.Println("====== Method Expression ======")
	a := &TA{name: "test"}
	var a1 = a.Mv("quick")
	fmt.Println("a1: ", a1)
}

type S struct{ *TS }
type TS int

func (t TS) M() { fmt.Println(t) }

type TB struct {
	a int
}

func (tv TB) Mv(a int) int          { return a }
func (tp *TB) Mp(f float32) float32 { return 1 }

var tb TB
var ptb *TB

func maketb() (TB int) { return 0 }

func methodValue() {
	fmt.Println("========= method value =========")
	t := new(TS)
	s := S{TS: t}
	fmt.Println("s: ", s)
	f := t.M
	fmt.Println("f: ", f)
	g := s.M
	fmt.Println("g: ", g)
	*t = 42
	fmt.Println("t: ", t)
	o := S{TS: *&t}
	fmt.Println("o: ", o)

	ex := tb.Mv(1291)
	fmt.Println("ex: ", ex)

	ex1 := tb.Mv
	ex1(1231)
	fmt.Printf("ex1: %v\n", ex1)
	fmt.Printf("ptb.Mp(51): %v\n", ptb.Mp(51))
}

func sliceExpression() {
	aK := make([]int, 0, 100)
	for i := 0; i < 10; i++ {
		aK = append(aK, i)
	}
	fmt.Println("aK: ", aK)
	// low:high
	fmt.Printf("aK[0:2]: %v\n", aK[0:2])
	// set ak[3]
	aK[3] = 99
	fmt.Printf("aK: %v\n", aK)
	fmt.Printf("aK: %v\n", aK[2:]) // same as aK[2 : len(aK)]
	fmt.Printf("aK: %v\n", aK[:3]) // same as aK[0 : 3]
	fmt.Printf("aK: %v\n", aK[:])  // same as aK[0: len(aK)]
	fmt.Printf("aK: %v\n", aK[0:2])

	ptraK := &aK
	fmt.Printf("ptraK: %v\n", (*ptraK)[0:2])
}

func typeAssertion() {
	var x interface{} = 7 // x has dynamic type int and value 7
	i := x.(int)          // i has type int and value 7
	fmt.Println('i', i)
	type MyInterface interface{ m() }
}

func callFunc() {
	fmt.Println("========== callFunc ==========")
	split := func(s string, pos int) (string, string) {
		return s[0:pos], s[pos:]
	}
	fmt.Printf("split: %v\n", split)
	aS1 := "abcdef"
	rs1, rs2 := split(aS1, 2)
	fmt.Printf("aS1: %v\n", rs1, rs2)
}

func variadic() {
	greeting := func(prefix string, who ...string) string {
		for i := 0; i < len(who); i++ {
			fmt.Println("greeting", who[i])
		}
		return "_" + prefix
	}
	fmt.Printf("greeting: %v\n", greeting)
	f2 := greeting("hello", "a", "b", "c")
	fmt.Printf("f2: %v\n", f2)

	sayFriends := func(s string, w ...string) []string {
		fmt.Println("sayFriend", s)
		return w
	}

	friends := []string{"bryan", "adam", "clark"}
	fmt.Printf("friends: %v\n", friends)
	f3 := sayFriends("hi", friends...)
	fmt.Printf("f3: %v\n", f3)

}

func main() {
	expSelector()
	methodExpression()
	methodValue()
	sliceExpression()
	typeAssertion()
	callFunc()
	variadic()
}
