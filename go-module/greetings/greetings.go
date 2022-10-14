package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// If no name was given, return an error message
	if name == "" {
		return name, errors.New("empty name")
	}
	// If name was received, return a value that embeds the name
	// Return greetings embedded in a message
	// message := fmt.Sprintf("Hi, %v. Welcome !", name)
	// Let's change message with random string
	message := fmt.Sprintf(randomFormat(), name)
	// This code will result test FAIL
	// message := fmt.Sprintf(randomFormat())
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with the name
		messages[name] = message
	}
	return messages, nil
}

// Init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat return one of a set greeting message.
// The returned mesage is selected at random
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome !",
		"Great to see you, %v!",
		"Hail, %v. ! Well met",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]

}
