Certainly! Here's an example program that incorporates the concepts covered in the lesson on user-defined types in Go:

```go
package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Address struct {
	street  string
	city    string
	country string
}

type Employee struct {
	Person
	company string
}

type Age int

func main() {
	// Declaring variables of struct types
	var person1 Person
	person1.name = "John Doe"
	person1.age = 30

	// Creating a struct value using a struct literal
	person2 := Person{name: "Jane Smith", age: 35}

	// Creating a struct value without declaring field names
	person3 := Person{"Alice Johnson", 25}

	// Declaring fields based on other struct types
	employee := Employee{
		Person:  Person{name: "Mark Thompson", age: 40},
		company: "ACME Inc.",
	}

	// Using struct literals to create values for fields
	address := Address{
		street:  "123 Main Street",
		city:    "New York",
		country: "USA",
	}

	// Pointers to structs
	employeePtr := &employee

	// Declaration of a new type based on the int type
	var myAge Age = 35

	// Printing the struct values
	fmt.Println("Person 1:", person1)
	fmt.Println("Person 2:", person2)
	fmt.Println("Person 3:", person3)
	fmt.Println("Employee:", employee)
	fmt.Println("Address:", address)
	fmt.Println("Employee (via pointer):", *employeePtr)
	fmt.Println("Age:", myAge)
}
```

This program demonstrates various aspects of user-defined types in Go:
- Declaration and initialization of struct variables.
- Creating struct values using struct literals.
- Creating struct values without explicitly declaring field names.
- Declaring fields based on other struct types.
- Using pointers to structs.
- Declaring a new type based on an existing type.

When you run the program, it will output the values of the struct variables, showcasing the different aspects discussed in the lesson.

Note: Make sure to save the code in a file with a ".go" extension (e.g., "main.go") and run it using the `go run` command: `go run main.go`.