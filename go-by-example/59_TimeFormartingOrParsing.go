// Go supports time formatting and parsing via pattern-based layouts.

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// Here’s a basic example of formatting a time according to RFC3339, using the corresponding layout constant.
	fmt.Println(t.Format(time.RFC3339))

	// Time parsing uses the same layout values as Format.
	t1, err := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	fmt.Println(t1)
	fmt.Println(err)

	// Format and Parse use example-based layouts. Usually you’ll use a
	// constant from time for these layouts, but you can also supply custom layouts.
	// Layouts must use the reference time Mon Jan 2 15:04:05 MST 2006 to show the
	// pattern with which to format/parse a given time/string. The example time must be
	// exactly as shown: the year 2006, 15 for the hour, Monday for the day of the week, etc.
	fmt.Println(t.Format("3:04PM"))
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	fmt.Println(t2)

	// For purely numeric representations you can also use standard string formatting
	// with the extracted components of the time value.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	// Parse will return an error on malformed input explaining the parsing problem.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, err = time.Parse(ansic, "8:41PM")
	fmt.Println(err)
}
