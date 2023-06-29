---
marp: true
theme: uncover

---
Title: The Nature of Types in Go

---
Objective: Understand the different types in Go and their usage.

---
Topic: Built-in Types

- Numeric Types: `int`, `float64`, `complex128`
- String Type: `string`
- Boolean Type: `bool`

---
Topic: Reference Types

- Slice Type: `[]T`
- Map Type: `map[K]V`
- Channel Type: `chan T`
- Interface Type: `interface{}`
- Function Type: `func(params) returnType`

---
Topic: Struct Types

- User-defined composite types
- Combination of primitive and non-primitive data
- Struct with Primitive Nature: Fields of built-in types
- Struct with Non-Primitive Nature: Fields of reference types or other struct types

---
Example: Person Struct

```go
// Person struct represents a person's information.
type Person struct {
	Name  string
	Age   int
	Email string
}

// PrintPerson prints the details of a Person.
func PrintPerson(p Person) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Email:", p.Email)
	fmt.Println()
}
```

---
Example: Book Struct

```go
// Book struct represents a book's information.
type Book struct {
	Title  string
	Author string
	Year   int
}

// PrintBook prints the details of a Book.
func PrintBook(b Book) {
	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
	fmt.Println("Year:", b.Year)
	fmt.Println()
}
```

---
Best Practices:
- Choose built-in types for basic data values
- Use reference types for complex data structures
- Employ struct types for composite types

---
Additional Resources:
- A Tour of Go: [tour.golang.org/basics/11](https://tour.golang.org/basics/11)
- Go by Example: [gobyexample.com/types](https://gobyexample.com/types)
- Effective Go: [golang.org/doc/effective_go.html#types_and_interfaces](https://golang.org/doc/effective_go.html#types_and_interfaces)
- Go Programming Language Specification: [golang.org/ref/spec#Types](https://golang.org/ref/spec#Types)

---
Conclusion:
- Types are fundamental in Go for representing data.
- Built-in types represent basic values.
- Reference types include slices, maps, channels, interfaces, and functions.
- Struct types combine primitive and non-primitive data.
- Choose types based on their nature and requirements.

Note: The code examples on slides 6 and 7 are formatted using code blocks. Ensure that the necessary import statements are included, and the code is properly indented according to Go's coding conventions.