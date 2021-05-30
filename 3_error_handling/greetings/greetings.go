package greetings

import ("fmt"
		"errors"
)

//Hello returns greeting for person

func Hello(name string) (string, error) {
	// If no name given, return blank with an error
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	/*
		You can also do the following
		var message string
		message = fmt.Sprintf(\"Hi, %v. Welcome!\",name)
	*/
	return message, nil
}