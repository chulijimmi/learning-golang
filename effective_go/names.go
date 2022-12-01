package main

import (
	"encoding/base64"
	"fmt"
)

type authors struct {
	name string
	dob  string
}

type Interface interface {
	Size() int
}

func (a authors) Login() string {
	fmt.Println("Name: ", a.name)
	return a.name
}

func (a authors) SetLogin(name string) {
	a.name = name
}

func getter(obj authors) string {
	user := "test"
	auth := obj.Login()
	if auth != user {
		obj.SetLogin("test")
	}
	return obj.name
}

func count(data Interface) int {
	x := data.Size()
	return x
}

func main() {
	// Package name
	data := []byte("Hello, world")
	dst := make([]byte, base64.RawStdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(dst, data)
	fmt.Println(string(dst))

	// Getters
	objType := authors{
		name: "bob",
		dob:  "10 June 2010",
	}
	x := getter(objType)
	fmt.Printf("x: %v\n", x)
	// Interface Name
	data := Interface{}
	i := count()

	// MixedCaps
}
