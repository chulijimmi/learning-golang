package main

import (
	"fmt"
	"sort"
)

type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

func (s Sequence) String() string {
	s = s.Copy() // Make a copy; don't overwrite argument.
	fmt.Println("Copy s String()")
	sort.Sort(s)
	str := "["
	for i, elem := range s { // Loop is O(N²); will fix that in next example.
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

func (s Sequence) run() string {
	s = s.Copy()
	fmt.Println("Copy s run()")
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

type Stringer interface {
	String() string
}

// To accept multiple argement we can use interface
func typeSwithc(value interface{}) string {
	switch str := value.(type) {
	case string:
		return str
	case int32:
		return string(str)
	case Stringer:
		return str.String()
	}
	return "nil"
}

type HandlerFunc func(a string) string

func (f HandlerFunc) Serve(a string) string {
	s := f(a)
	fmt.Printf("s: %v\n", s)
	return "Two " + a
}

func main() {
	fmt.Println("=========== Method & Interface ===========")
	s := Sequence{1, 2, 3, 4}
	str := s.String()
	fmt.Printf("str: %v\n", str)
	fmt.Println(" =========== Conversion =========== ")
	run := s.run()
	fmt.Printf("run: %v\n", run)
	fmt.Println(" =========== Interface conversions and type assertions	=========== ")
	tS := typeSwithc(123)
	fmt.Printf("tS: %v\n", tS)
	var is interface{} = int32(111)
	tIs := typeSwithc(is)
	fmt.Printf("tIs: %v\n", tIs)
	fmt.Println(" =========== Generality =========== ")
	fmt.Println(" =========== Interface & Method =========== ")
	var x HandlerFunc = func(a string) string { return "One " + a }
	fmt.Printf("fb: %v\n", x.Serve("test"))
}
