// Go provides built-in support for base64 encoding/decoding.
// This syntax imports the encoding/base64 package with the b64 name instead of the default base64.
// It’ll save us some space below.

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// Here’s the string we’ll encode/decode.
	data := "abc123!?$*&()'-=@~"

	// Go supports both standard and URL-compatible base64.
	// Here’s how to encode using the standard encoder.
	// The encoder requires a []byte so we convert our string to that type.
	dataEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Data encoding: ", dataEnc)

	// Decoding may return an error, which you can check if you
	// don’t already know the input to be well-formed.
	dataDec, err := base64.StdEncoding.DecodeString(dataEnc)
	fmt.Println("Data decoding: ", string(dataDec))
	fmt.Println(err)

	// This encodes/decodes using a URL-compatible base64 format.
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("Url encoding: ", uEnc)
	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println("Url decoding: ", string(uDec))
}
