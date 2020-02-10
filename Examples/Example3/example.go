package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	// Go supports recursion
	return n * factorial(n-1)
}

// Go allows multiple returns
func sortNumbers(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

func main() {
	fac10, fac11 := factorial(10), factorial(11)

	// Use _ when you don't need a result. Go will complain about unused variables
	bigger, _ := sortNumbers(fac10, fac11)

	fmt.Println("The biggest result was: ", bigger)
}
