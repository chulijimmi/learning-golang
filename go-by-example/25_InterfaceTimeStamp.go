// The real world: getting a proper timestamp out of the Twitter API
// The Twitter API represents timestamps using a string of the following format:

// "Thu May 31 00:00:01 +0000 2012"
// of course, timestamps can be represented in any number of ways in a JSON document,
// because timestamps aren’t a part of the JSON spec. For the sake of brevity,
// I won’t put in the entire JSON representation of a tweet, but let’s take a look
// at how the created_at field would be handled by encoding/json:

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

var input = `
{
	"created_at": "Thu May 31 00:00:01 +0000 2012"
}
`

type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

type Timestamp time.Time

// By implementing this method we satisfy the json.Unmarshaler interface,
// causing json.Unmarshal to call our custom unmarshalling code when seeing
// a Timestamp value. For this case, we use a pointer method, because we want
// the caller to the see the changes made to the receiver. In order to set
// the value that a pointer points to, we dereference the pointer manually
// using the * operator. Inside of the UnmarshalJSON method, t represents a
// pointer to a Timestamp value. By saying *t, we dereference the pointer t
// and we are able to access the value that t points to.
// Remember: everything is pass-by-value in Go. That means that inside of
// the UnmarshalJSON method, the pointer t is not the same pointer as
// the pointer in its calling context; it is a copy. If you were to
// assign t to another value directly, you would just be reassigning a
// function-local pointer; the change would not be seen by the caller.
// However, the pointer inside of the method call points to the same data
// as the pointer in its calling scope; by dereferencing the pointer,
// we make our change visible to the caller.

// We can make use of the time.Parse method, which has the signature
// func(layout, value string) (Time, error). That is, it takes two strings:
// the first string is a layout string that describes how we are formatting
// our timestamps, and the second is the value we wish to parse.
// It returns a time.Time value, as well as an error
// (in case we failed to parse the timestamp for some reason).
// You can read more about the semantics of the layout strings in the
// time package documentation, but in this example we won’t need to
// figure out the layout string manually because this layout string
// already exists in the standard library as the value time.RubyDate.
// So in effect, we can resolve the string “Thu May 31 00:00:01 +0000 2012”
// to a time.Time value by invoking the function
// time.Parse(time.RubyDate, "Thu May 31 00:00:01 +0000 2012").
// The value we will receive is of type time.Time. In our example,
// we’re interested in values of the type Timestamp. We can convert the
// time.Time value to a Timestamp value by saying Timestamp(v),
// where v is our time.Time value. Ultimately, our UnmarshalJSON
// function winds up looking like this:

// we take a subslice of the incoming byte slice because the incoming
// byte slice is the raw data of the JSON element and contains the
// quotation marks surrounding the string value; we want to chop
// those off before passing the string value into time.Parse.

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Timestamp(v)
	return nil
}

func main() {
	// our target will be of type map[string]interface{}, which is a
	// pretty generic type that will give us a hashtable whose keys
	// are strings, and whose values are of type interface{}
	var val map[string]Timestamp

	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
	fmt.Println(time.Time(val["created_at"]))
}
