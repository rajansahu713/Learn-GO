package main

import "fmt"

func main() {
	// String
	var name string = "Alice"
	fmt.Println("Name:", name)

	// Integer
	var age int = 30
	fmt.Println("Age:", age)

	// Float
	var height float64 = 5.9
	fmt.Println("Height:", height)

	// Boolean
	var isStudent bool = true
	fmt.Println("Is Student:", isStudent)

	// Default values
	var defaultString string
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool

	fmt.Println("Default String:", defaultString)
	fmt.Println("Default Int:", defaultInt)
	fmt.Println("Default Float:", defaultFloat)
	fmt.Println("Default Bool:", defaultBool)

	// other ways to create variables
	city := "Guwahati" // type is inferred
	fmt.Println("City:", city)
	age2 := 25
	fmt.Println("Age2:", age2)
}