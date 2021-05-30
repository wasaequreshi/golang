package main

import ("fmt"
		"log"
		"example.com/greetings"
)

func main() {
	// Setting properites of logger
	
	// Print at the start of the logs the prefix
	log.SetPrefix("greetings: ")
	// Don't print time, source file, and line number
	log.SetFlags(0)

	names := []string {
		"Wasae",
		"Bilal",
	}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
	
}

