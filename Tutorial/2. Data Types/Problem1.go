// adding two numbers 
package main

import (
	"fmt"
)

func main() {

	var a, b, c int

	fmt.Println("Enter 1st Number: ")

	fmt.Scan(&a)

	fmt.Println(a)

	fmt.Println("Enter 2nd Number: ")

	fmt.Scan(&b)
	fmt.Println(b)

	c = a + b

	fmt.Printf("Sum: %d \n", c)

}
