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
