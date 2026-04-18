// while will be look like same as foor loop

package main

import (
	"fmt"
)

func main(){

	count := 5

	for count > 0 {
		fmt.Println("count", count)
		count -= 1
	}
}