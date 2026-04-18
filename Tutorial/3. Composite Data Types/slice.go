package main

import "fmt"

func main() {
	var a =[]int{}
	// add
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)

	// remove
	a = a[1:3]


	var result int
	for _, v := range a {
		result += v
	}
	fmt.Printf("Sum of elements in slice: %d\n", result)

	fmt.Println("Length of slice:", len(a))
	fmt.Println("Capacity of slice:", cap(a))
	fmt.Println(a)
}
