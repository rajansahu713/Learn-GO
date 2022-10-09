package main

import "fmt"

func main() {

	var a, b int

	fmt.Println("Enter 1st Number")
	fmt.Scan(&a)

	fmt.Println("Enter 2nd Number")
	fmt.Scan(&b)
	if a > b {
		fmt.Println("a is greater than b")
	} else if a == b {
		fmt.Println("a and b both are equal")

	} else {
		fmt.Println("a is less than b")
	}
}
