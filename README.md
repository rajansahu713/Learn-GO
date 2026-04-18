<h1 align="center">Go</h1>

1. What is Go?
* It is an open-source, compiled, and statically typed programming language.
* It is designed by Google
* A statically-typed language is a language (such as Java, C, or C++) where variable types are known at compile time.

2. why go?
* It is a fast, statically-typed(which means attribute type must be know before compilation), compiled language.
* Go has fast run time and compilation time
* Go supports concurrency
* Go has memory management

3. How to run go program?
* write the program and save it as .go extension file.
* run the program by *go run filename.go* 


4. How to print Hello World in Go?

```go
    package main
    import "fmt"
    func main() {
        fmt.Println("Hello World")
    }
```




* How to take a input from user?
* How to add two numbers?


We can declear variable two ways 


1. You always have to specify either type or value (or both).
```go
var variablename type = value
```


2. With the := sign:
Use the := sign, followed by the variable value:

```go
variablename := value
```
In the second case it will inferred from the `variable` (means that the compiler decides the type of the variable, based on the value)


Example
```go
var student1 string = "John" //type is string
var student2 = "Jane" //type is inferred
x := 2 //type is inferred
```

Whenever you declare a variable without assigning a value, it will be assigned a default value based on its type. 
```go
var a string
var b int
var c bool
```

if you try to print you will see
```
""
0
false
```

```go
var student1 string
student1 = "John"
```


Multiple variable
```go
a,b,c int = 1,2,3
```

Multi-Word Variable Names

1. Camel Case
Each word, except the first, starts with a capital letter:
myVariableName = "John"

2. Pascal Case
Each word starts with a capital letter:
MyVariableName = "John"

3. Snake Case
Each word is separated by an underscore character:
my_variable_name = "John"

4. How to create a variable in go that is not changable

* for this use constant key word in go
* it is read only field
```go
const CONSTNAME type = value
```

Go has three basic data types:

bool: represents a boolean value and is either true or false
Numeric: represents integer types, floating point values, and complex types
string: represents a string value


### Rule of variable declaration in go
1. Variable names must start with a letter or an underscore.
2. Variable names can only contain letters, numbers, and underscores.
3. Variable names cannot be a reserved keyword in Go.
4. if a variable is declared but not used, the compiler will throw an error. This is to encourage clean and efficient code.

### different between var and const
1. var is used to declare a variable that can be changed, while const is used to declare a constant that cannot be changed.
2. var can be assigned a value at runtime, while const must be assigned a value at compile time.
3. var can be used to declare variables of any type, while const can only be used to declare constants of basic types (such as int, float, string, etc.).

### different between array and slice
1. An array is a fixed-size collection of elements of the same type, while a slice is a dynamically-sized, flexible view into the elements of an array.
2. An array has a fixed length that is determined at compile time, while a slice can grow and shrink dynamically as elements are added or removed.

Declaration of array
```go
var arr [5]int // declares an array of 5 integers
```

Declaration of slice
```go
var s []int // declares a slice of integers
```


## map in go
A map is a collection of key-value pairs, where each key is unique and maps to a single value. Maps are also known as hash tables or dictionaries in other programming languages.

Declaration of map
```go
var m map[string]int // declares a map with string keys and int values
```

if you try to search for a key that does not exist in the map

