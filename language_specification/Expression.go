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

func instantiations() {
	min := func(x, y int) int {
		return x + y
	}
	fmt.Println(min(2, 3))
}

func typeUnification() {
	type a []map[int]bool
	b := make(a, 2)
	fmt.Printf("b: %v\n", b)
}

func operators() {
	// Expression = UnaryExpr | Expression binary_op Expression .
	// UnaryExpr  = PrimaryExpr | unary_op UnaryExpr .

	// binary_op  = "||" | "&&" | rel_op | add_op | mul_op .
	// rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
	// add_op     = "+" | "-" | "|" | "^" .
	// mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

	// unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .

	var a [8]byte
	fmt.Printf("a: %v\n", a)
	var s uint = 33
	fmt.Printf("s: %v\n", s)

	var s1 = 1 << s
	fmt.Printf("i: %v\n", s1)
	var s2 int32 = 1 << 2
	fmt.Printf("s2: %v\n", s2)
	var s3 = uint64(1 << s)
	fmt.Printf("s3: %v\n", s3)
	var s4 int = 1.0 << s
	fmt.Printf("s4: %v\n", s4)
	var s5 = 1.0<<s == s2
	fmt.Printf("s5: %v\n", s5)

	// Operato predecence
	var n1 = 5
	var n2 = 10
	if n1 < 5 || n2 > n1 {
		fmt.Println("result predecence 1")
	}

	// floating point operator
	r := 0
	x := 1
	y := 2
	z := 3

	r = x*y + z
	fmt.Printf("r: %v\n", r)
	r = z
	fmt.Printf("r: %v\n", r)
	r += x * y
	fmt.Printf("r: %v\n", r)

	// string concatenation
	stra := "kitty"
	strb := "hi " + string(stra)
	strb += " and good bye"
	fmt.Printf("strb: %v\n", strb)
}

func comparisonOperator() {
	// ==    equal
	// !=    not equal
	// <     less
	// <=    less or equal
	// >     greater
	// >=    greater or equal
	eq := 1 == 1
	fmt.Printf("eq: %v\n", eq)
	neq := !true
	fmt.Printf("neq: %v\n", neq)
}

func logicalOperator() {
	and := true && false
	fmt.Printf("and: %v\n", and)
	or := true || false
	fmt.Printf("or: %v\n", or)
	not := !false
	fmt.Printf("not: %v\n", not)
}

func addressOperator() {
	var x int = 1
	fmt.Printf("x: %v\n", x)
	var y *int = &x
	fmt.Printf("y: %v\n", y)
	fmt.Printf("*y: %v\n", *y)
	z := *&*y + 3
	fmt.Printf("z: %v\n", z)
	fmt.Printf("y: %v\n", *y)
	fmt.Printf("x: %v\n", x)

	// Change data source pointer
	var number = 5
	fmt.Printf("number before: %v\n", number)
	change := func(src *int, value int) {
		*src = value
	}
	change(&number, 10)
	fmt.Printf("number after: %v\n", number)
}

func conversion() {
	type Person struct {
		Name    string
		Address *struct {
			Street string
			City   string
		}
	}

	var data *struct {
		Name    string `json:"name"`
		Address *struct {
			Street string `json:"street"`
			City   string `json:"city"`
		} `json:"address"`
	}

	var person = (*Person)(data) // ignoring tags, the underlying types are identical
	fmt.Printf("person: %v\n", person)

	hello := []byte("hellø")
	fmt.Printf("hello: %v\n", hello)

	a := 10
	b := string(a)
	fmt.Printf("b: %v\n", b)
}

func constExpression() {
	const b = 15 / 4
	fmt.Printf("b: %v\n", b)
}

func orderOfEvaluation() {
	a := 1
	f := func() int { a++; return a }
	x := []int{a, f()} // x may be [1, 2] or [2, 2]: evaluation order between a and f() is not specified
	fmt.Printf("x: %v\n", x)
	m := map[int]int{a: 1, a: 2} // m may be {2: 1} or {2: 2}: evaluation order between the two map assignments is not specified
	fmt.Printf("m: %v\n", m)
	n := map[int]int{a: f()} // n may be {2: 3} or {3: 3}: evaluation order between the key and the value is not specified
	fmt.Printf("n: %v\n", n)
}

func main() {
	expSelector()
	methodExpression()
	methodValue()
	sliceExpression()
	typeAssertion()
	callFunc()
	variadic()
	instantiations()
	typeUnification()
	operators()
	comparisonOperator()
	logicalOperator()
	addressOperator()
	conversion()
	constExpression()
	orderOfEvaluation()
}
