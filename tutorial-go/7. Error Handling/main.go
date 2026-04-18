package main

import (
	"fmt"
)

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "denominator cannot be zero"
	}
	return a / b, ""
}



func main() {
	// Example of error handling in Go
	fmt.Println("Starting the program...")

	result, err := divide(10, 0)
	if err != "" {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result: %f\n", result)
	}
}