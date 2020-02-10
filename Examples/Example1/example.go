package main

import "fmt"

// This is a comment
func main() {
	// Variable initialisation
	var message1 string
	message1 = "Hello World!"

	// Go can infer types, this is the shorthand
	message2 := "Hi Cambridge"

	fmt.Println(message1)
	fmt.Println(message2)
}
