// Pointer VS Value

package method

type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
	// Body exactly the same as the Append function defined above.
	slice = append(slice, data...)
	return slice
}

// This still requires the method to return the updated slice. We can eliminate that clumsiness
// by redefining the method to take a pointer to a ByteSlice as its receiver,
// so the method can overwrite the caller's slice.
func (p *ByteSlice) Add(data []byte) {
	slice := *p
	*p = slice
}

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	*p = slice
	return len(data), nil
}
