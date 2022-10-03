// Go offers built-in support for JSON encoding and decoding,
// including to and from built-in and custom data types.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// We’ll use these two structs to demonstrate encoding
// and decoding of custom types below.
type responseA struct {
	Page   int
	Fruits []string
}

// Only exported fields will be encoded/decoded in JSON.
// Fields must start with capital letters to be exported.
type responseB struct {
	Page   int      //`json: "page"`
	Fruits []string //`json: "fruits"`
}

func main() {
	// First we’ll look at encoding basic data types to JSON strings.
	// Here are some examples for atomic values.
	bolA, error := json.Marshal(true)
	fmt.Println(string(bolA))
	fmt.Println(error)

	intA, _ := json.Marshal(1)
	fmt.Println(string(intA))

	floatA, _ := json.Marshal(2.34)
	fmt.Println(string(floatA))

	stringA, _ := json.Marshal("gopher")
	fmt.Println(string(stringA))

	// And here are some for slices and maps, which encode to
	// JSON arrays and objects as you’d expect.
	sliceA := []string{"apple", "peach", "pear"}
	sliceAToJson, _ := json.Marshal(sliceA)
	fmt.Println(string(sliceAToJson))

	mapA := map[string]int{"apple": 5, "banana": 7}
	mapAToJson, _ := json.Marshal(mapA)
	fmt.Println(string(mapAToJson))

	// The JSON package can automatically encode your custom data types.
	// It will only include exported fields in the encoded output and will
	// by default use those names as the JSON keys.
	resA := &responseA{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	resAToJson, _ := json.Marshal(resA)
	fmt.Println(string(resAToJson))

	// You can use tags on struct field declarations to customize the encoded
	// JSON key names. Check the definition of response2 above to see an example of such tags.
	resB := &responseB{Page: 2, Fruits: []string{"a", "b", "c"}}
	resBToJson, _ := json.Marshal(resB)
	fmt.Println(string(resBToJson))

	// Now let’s look at decoding JSON data into Go values.
	// Here’s an example for a generic data structure.
	byt := []byte(`{"num": 6.13, "str": ["xyz", "def"]}`)

	// We need to provide a variable where the JSON package can put the decoded data.
	// This map[string]interface{} will hold a map of strings to arbitrary data types.
	var dataMap map[string]interface{}

	// Here’s the actual decoding, and a check for associated errors.
	if err := json.Unmarshal(byt, &dataMap); err != nil {
		panic(err)
	}
	fmt.Println(dataMap)

	// In order to use the values in the decoded map, we’ll need to convert them to
	// their appropriate type. For example here we convert the value in num to
	// the expected float64 type.
	num := dataMap["num"].(float64)
	fmt.Println("num", num)

	// Accessing nested data requires a series of conversions.
	// strs := dataMap["strs"].([]interface{})
	// str1 := strs[0].(string)
	// fmt.Println("str1", str1)

	// We can also decode JSON into custom data types. This has the advantages of adding
	// additional type-safety to our programs and eliminating the need
	// for type assertions when accessing the decoded data.
	fmt.Println("========= Decode JSON =========")
	str := `{"page": 1, "fruits": ["apple","peach"]}`
	res := responseB{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// In the examples above we always used bytes and strings as intermediates between
	// the data and JSON representation on standard out. We can also stream JSON
	// encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	fmt.Println("========= Encoding =========")
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "banana": 10}
	enc.Encode(d)

}
