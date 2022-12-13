package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	// Import for side effect
	_ "net/http/pprof"
)

// Unused imports and variables
var _ = fmt.Printf

func main() {
	fmt.Println("=============== The blank identifier in multiple assignment ===============")
	path := "/etc/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", path)
	}
	fmt.Println("=============== Unused imports and variables ===============")
	fd, er := os.Open("test.go")
	if er != nil {
		log.Fatal(er)
	}
	// TODO: use fd.
	_ = fd

	// Interface check
	var val interface {
		name() string
	}

	if _, ok := val.(json.Marshaler); ok {
		fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
	}
}
