package main

import ("fmt"
		"example.com/greetings"
)

func main() {
	//Get greeting message and print
	message := greetings.Hello("Wasae")
	fmt.Println(message)
}

