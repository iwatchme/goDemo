package main

import (
	"awesomeProject/greetings"
	"fmt"
	"log"
)

func main() {
	messages, err := greetings.Hello("Frank")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
