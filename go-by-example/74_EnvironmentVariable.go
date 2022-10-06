// Environment variables are a universal mechanism for conveying configuration
// information to Unix programs. Let’s look at how to set, get, and list environment variables.
// Reference: https://en.wikipedia.org/wiki/Environment_variable
// Reference: https://www.12factor.net/config

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// To set a key/value pair, use os.Setenv. To get a value for a key, use os.Getenv.
	// This will return an empty string if the key isn’t present in the environment.\
	os.Setenv("alpha", "http://alpha.domain.com")
	os.Setenv("beta", "http:///beta.domain.com")
	fmt.Println("ENV alpha: ", os.Getenv("alpha"))
	fmt.Println("ENV beta: ", os.Getenv("beta"))
	fmt.Println("ENV prod: ", os.Getenv("prod"))

	// Use os.Environ to list all key/value pairs in the environment.
	// This returns a slice of strings in the form KEY=value.
	// You can strings.SplitN them to get the key and value. Here we print all the keys.
	fmt.Println("============== os.Environ ==============")
	for _, el := range os.Environ() {
		fmt.Println("el", el)
		pair := strings.SplitN(el, "=", 2)
		fmt.Println("Pair[0]", pair[0])
	}
}
