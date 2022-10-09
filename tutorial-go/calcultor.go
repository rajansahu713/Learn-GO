package main

import (
	"fmt"
)

func add_values(a int, b int) int {
	return a + b
}

func sub_values(a int, b int) int {
	return a - b
}

func div_values(a int, b int) int {
	return a / b
}

func mul_values(a int, b int) int {
	return a * b
}

func main() {

	var a, b int

	fmt.Println("Enter 1st Number: ")

	fmt.Scan(&a)

	fmt.Println("First Number Enter by You ", a)

	fmt.Println("Enter 2nd Number: ")

	fmt.Scan(&b)
	fmt.Println("Second Number Enter by You ", b)

	fmt.Printf("add: %d \n", add_values(a, b))
	fmt.Printf("Sub: %d \n", sub_values(a, b))
	fmt.Printf("mul: %d \n", mul_values(a, b))
	fmt.Printf("div: %d \n", div_values(a, b))

}
