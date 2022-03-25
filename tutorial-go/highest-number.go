package main

import "fmt"

func add_values(a int, b int, c int) int {

	var d = a + b + c

	return d
}

func main() {

	var a, b, c, d int

	fmt.Println("Enter 1st Number: ")

	fmt.Scan(&a)

	fmt.Println(a)

	fmt.Println("Enter 2nd Number: ")

	fmt.Scan(&b)

	fmt.Println("Enter 3rd Number: ")

	fmt.Scan(&c)
	fmt.Println(b)

	d = add_values(a, b, c)

	fmt.Printf("Sum: %d \n", d)

}
