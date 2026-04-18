package main

import (
	"fmt"
)

func getting() string {
	return "hello"
}

func main(){
	message := getting()

	fmt.Println("message: ", message)
}