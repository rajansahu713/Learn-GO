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

whenever you declear a variable in golang you will a default value to it.
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