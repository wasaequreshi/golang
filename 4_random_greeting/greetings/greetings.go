package greetings

import ("fmt"
		"errors"
		"math/rand"
		"time"
)

//Hello returns greeting for person

func Hello(name string) (string, error) {
	// If no name given, return blank with an error
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	/*
		You can also do the following
		var message string
		message = fmt.Sprintf(\"Hi, %v. Welcome!\",name)
	*/
	return message, nil
}

//Initial values set for variables
func init() {
	rand.Seed(time.Now().UnixNano())
}

//Returns one of the greeting messages
func randomFormat() string {
	// Different string outputs
	formats := []string {
		"Hi, %v. Welcome!",
        "Great to see you, %v!",
		"Hello %v. Nice to meet you",
	}
	// Picks an int randomly between 0 and 1
	return formats[rand.Intn(len(formats))]
}