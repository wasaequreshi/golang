package main

import ("fmt"
		"log"
		"example.com/greetings"
)

// Creating function to print results of greetings.Hello with good and bad example
func print(message string, err error) {
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

func main() {
	// Setting properites of logger
	
	// Print at the start of the logs the prefix
	log.SetPrefix("greetings: ")
	// Don't print time, source file, and line number
	log.SetFlags(0)

	//Get greeting message and if there was an error (in this case no)
	message, err := greetings.Hello("Wasae")
	print(message, err)

	//Get greeting message and if there was an error (in this case yes)
	message, err = greetings.Hello("")
	print(message, err)
	
}

