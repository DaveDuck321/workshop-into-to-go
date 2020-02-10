package main

import "fmt"

func main() {
	// A 'standard' for loop:
	// 3 parts: init statement, condition expression, post statement
	for i := 0; i < 3; i++ {
		fmt.Println("For ; ; ;", i)
	}

	// A 'while' for loop
	// One condition expression
	count := 6
	for count > 0 {
		count -= 2
		fmt.Println("For <bool>", count)
	}

	// A FORever loop
	for {
		fmt.Println("Forever!")
		count++

		// Stop the program running forever
		if count > 3 {
			break
		}
	}
}
