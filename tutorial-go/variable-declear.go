package main

import "fmt"

func display(b int) (a int) {
	a = b
	return
}

func main() {
	var a int = 5
	fmt.Println(a)
	print(display(a))
}
