package main

import (
	"fmt"
)


func main(){
	// simple loop creation
	var sum int

	for i :=0; i < 10 ; i++{
		sum = sum + i
	}
	fmt.Println("sum: ", sum)

	// iterating a arr
	nums := [5]int{4,6,2,6}

	for index, num := range nums {
		fmt.Println("index: %d, num: %d", index, num)
	}

}