// In this step, you’ll add a function to reverse a string.
package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

// func Reverse(s string) string {
// 	b := []byte(s)
// 	fmt.Println(b)
// 	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
// 		b[i], b[j] = b[j], b[i]
// 	}
// 	return string(b)
// }

// The entire seed corpus used strings in which every character was a single byte.
// However, characters such as 泃 can require several bytes.
// Thus, reversing the string byte-by-byte will invalidate multi-byte characters.
// func Reverse(s string) string {
// 	r := []rune(s)
// 	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return string(r)
// }

// Look closely at the reversed string to spot the error. In Go, a string is a read
// only slice of bytes, and can contain bytes that aren’t valid UTF-8.
// The original string is a byte slice with one byte, '\x91'.
// When the input string is set to []rune, Go encodes the byte slice to UTF-8, and
// replaces the byte with the UTF-8 character �. When we compare the replacement
// UTF-8 character to the input byte slice, they are clearly not equal.
func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, errRev := Reverse(input)
	doubleRev, errDoubleRev := Reverse(rev)
	fmt.Printf("Original: %q\n", input)
	fmt.Printf("Reversed: %q, err: %v\n", rev, errRev)
	fmt.Printf("Reverse Again: %q, err: %v\n", doubleRev, errDoubleRev)
}
