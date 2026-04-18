package main
import (
	"fmt"
)


func main() {
	//create a map
	m := make(map[string]int)

	// add key-value pairs to the map
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	fmt.Println("Map:", m["a"]) // Output: Map: 1
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	
	// delete a key-value pair from the map
	delete(m, "b")
	fmt.Println("Map after deletion:", m) // Output: Map after deletion: map[a:1 c:3]

	// check if a key exists in the map
	if value, exists := m["d"] 
	exists {
		fmt.Printf("Key 'd' exists with value: %d\n", value)
	} else {
		fmt.Println("Key 'd' does not exist in the map")
	}

	// nested map
	nestedMap := make(map[string]map[int]int)
	nestedMap["first"] = make(map[int]int)
	nestedMap["first"][1] = 10
	nestedMap["first"][2] = 20

	fmt.Println("Nested Map:", nestedMap["first"]) // Output: Nested Map: map[first:map[1:10 2:20]]
}