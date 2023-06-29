Exercise: Working with Struct Types

Objective: The objective of this exercise is to reinforce the understanding of struct types in Go and their usage in representing data values with primitive or non-primitive nature.

Instructions:
1. Create a struct type called "Person" with the following fields:
   - Name (string): represents the person's name.
   - Age (int): represents the person's age.
   - Email (string): represents the person's email address.

2. Create two instances of the "Person" struct, initializing them with different values for each field.

3. Write a function called "PrintPerson" that takes a "Person" struct as a parameter and prints out the values of all its fields.

4. Call the "PrintPerson" function for each of the created instances of the "Person" struct.

5. Create a new struct type called "Book" with the following fields:
   - Title (string): represents the title of the book.
   - Author (string): represents the author of the book.
   - Year (int): represents the publication year of the book.

6. Create an instance of the "Book" struct, initializing it with values for each field.

7. Write a function called "PrintBook" that takes a "Book" struct as a parameter and prints out the values of all its fields.

8. Call the "PrintBook" function for the created instance of the "Book" struct.

Solution:

```go
package main

import "fmt"

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

func main() {
	// Create two instances of Person.
	person1 := Person{Name: "John Doe", Age: 30, Email: "johndoe@example.com"}
	person2 := Person{Name: "Jane Smith", Age: 35, Email: "janesmith@example.com"}

	// Call PrintPerson for each instance of Person.
	fmt.Println("Person 1:")
	PrintPerson(person1)

	fmt.Println("Person 2:")
	PrintPerson(person2)

	// Create an instance of Book.
	book := Book{Title: "The Go Programming Language", Author: "Alan A. A. Donovan and Brian W. Kernighan", Year: 2015}

	// Call PrintBook for the instance of Book.
	fmt.Println("Book:")
	PrintBook(book)
}
```

Explanation:
In this exercise, we start by creating a struct type called "Person" with fields for name, age, and email. We then create two instances of the "Person" struct, initializing them with different values. We also define a function called "PrintPerson" that takes a "Person" struct as a parameter and prints out the values of its fields.

Next, we create a struct type called "Book" with fields for title, author, and year. We create an instance of the "Book" struct, initializing it with values. We define a function called "PrintBook" that takes a "Book" struct as a parameter and prints out the values of its fields.

In the `main()` function, we call the "PrintPerson" function for each