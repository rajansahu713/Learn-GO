// Building simple calculator

package main

import (
	"fmt"
)

func main() {
	
	a := 10
	b := 20

	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b

	fmt.Println("Addition: %d, Subtraction: %d, Multiplication: %d, Division: %d", sum, sub, mul, div)
}