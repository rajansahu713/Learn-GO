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
	var operation, a, b int
	fmt.Println("Please choose appropiate operations \n1. Multiple \n2. Addition \n3. Dividing \n4. Substract")
	fmt.Scan(&operation)
	fmt.Println("Please Enter 1st Number: ")
	fmt.Scan(&a)
	fmt.Println("Please Enter 2nd Number: ")
	fmt.Scan(&b)

	switch operation {
	case 1:
		fmt.Printf("multiplication: %d \n", mul_values(a, b))
	case 2:
		fmt.Printf("addition: %d \n", add_values(a, b))
	case 3:
		fmt.Printf("dividing: %d \n", div_values(a, b))
	case 4:
		fmt.Printf("substract: %d \n", sub_values(a, b))
	}
}
