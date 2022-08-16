// Reference:
// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
// https://research.swtch.com/interfaces

package main

import "fmt"

type Animal interface {
	speak() string
}

type Dog struct {
}

func (d Dog) speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) speak() string {
	return "Meow!"
}

type Human struct {
}

func (h Human) speak() string {
	return "Helloooo!"
}

type Engineer struct {
}

func (e Engineer) speak() string {
	return "Design pattern!"
}

func printAnimal() {
	animals := []Animal{Dog{}, Cat{}, Human{}, Engineer{}}
	for _, animal := range animals {
		fmt.Println(animal.speak())
	}
}

// The interface{} type, the empty interface, is the source of much confusion.
// The interface{} type is the interface that has no methods. Since there is
// no implements keyword, all types implement at least zero methods, and
// satisfying an interface is done automatically, all types satisfy the
// empty interface. That means that if you write a function that takes an
// interface{} value as a parameter, you can supply that function with any value.

func doSomething(v interface{}) {
	// Here’s where it gets confusing: inside of the DoSomething function,
	// what is v’s type? Beginner gophers are led to believe that
	// “v is of any type”, but that is wrong. v is not of any type; it
	// is of interface{} type. Wait, what? When passing a value into
	// the DoSomething function, the Go runtime will perform a
	// type conversion (if necessary), and convert the value to an interface{} value.
	// All values have exactly one type at runtime, and v’s one static type is interface{}.

}

func printVals(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

// By running this, you can see that we encounter the following error:
// cannot use names (type []string) as type []interface {} in function
// argument. If we want to actually make that work,
// we would have to convert the []string to an []interface{}:
// Code:
// names := []string{"stanley", "bob", "alice"}
// printVals(names)
// Solve covert the [] string to an []interface{}
func printName() {
	names := []string{"stanley", "bob", "alice"}
	vals := make([]interface{}, len(names))
	for key, value := range names {
		vals[key] = value
	}
	printVals(vals)
}

type Bird struct {
}

func (b *Bird) speak() string {
	return "Vue vue !"
}

func printAnimalPointer() {
	// go-by-example\InterfacesAnimal.go:93:22: cannot use Bird{} (value of type Bird) as type Animal in array or slice literal:
	// Bird does not implement Animal (speak method has pointer receiver)
	// animals := []Animal{Bird{}}
	animals := []Animal{new(Bird), Cat{}, new(Human)}
	for _, animal := range animals {
		fmt.Println("pointer", animal)
		fmt.Println("speak", animal.speak())
	}
}

func main() {
	printAnimal()
	printName()
	printAnimalPointer()
}
