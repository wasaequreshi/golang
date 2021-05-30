package greetings

import "fmt"

//Hello returns greeting for person

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	/*
		You can also do the following
		var message string
		message = fmt.Sprintf(\"Hi, %v. Welcome!\",name)
	*/
	return message
}