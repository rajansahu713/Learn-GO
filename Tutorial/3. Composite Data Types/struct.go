package main
import ("fmt")


type Student struct {
	name string
	age  int
	address string
}


func main() {
	var s1 = Student{}

	// ways to assign values to struct fields
	s1.name = "Alice"
	s1.age = 20
	s1.address = "123 Main St"
	
	fmt.Println("Student:", s1)

	s2 := Student{name: "Bob", age: 22, address: "456 Elm St"}
	fmt.Println("Student: ", s2)
}