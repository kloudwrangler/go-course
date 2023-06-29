Exercise: Creating and Manipulating Structs in Go

Instructions:
1. Create a new Go file named "structs_exercise.go" on your local machine.
2. Implement the following tasks based on the concepts taught in the lesson.
3. Test your code by running the program and verifying the expected output.
4. Once you have completed the exercise, you can refer to the provided solution and explanation.

Tasks:
1. Define a struct type named "Book" with the following fields:
   - Title (string)
   - Author (string)
   - YearPublished (int)

2. Create a function named "NewBook" that takes the title, author, and year published as parameters and returns a Book struct value.

3. Create a function named "PrintBookDetails" that takes a Book struct as a parameter and prints its title, author, and year published.

4. In the main function:
   a. Declare a variable "book1" of type Book using the NewBook function. Provide your favorite book details as arguments.
   b. Declare a variable "book2" of type Book using a struct literal and initialize it with values for title, author, and year published of another book of your choice.
   c. Call the PrintBookDetails function with "book1" and "book2" as arguments to display their details.

Solution:

```go
package main

import "fmt"

type Book struct {
	Title         string
	Author        string
	YearPublished int
}

func NewBook(title, author string, yearPublished int) Book {
	return Book{
		Title:         title,
		Author:        author,
		YearPublished: yearPublished,
	}
}

func PrintBookDetails(book Book) {
	fmt.Printf("Title: %s\nAuthor: %s\nYear Published: %d\n", book.Title, book.Author, book.YearPublished)
}

func main() {
	book1 := NewBook("The Great Gatsby", "F. Scott Fitzgerald", 1925)
	book2 := Book{
		Title:         "To Kill a Mockingbird",
		Author:        "Harper Lee",
		YearPublished: 1960,
	}

	PrintBookDetails(book1)
	fmt.Println()
	PrintBookDetails(book2)
}
```

Explanation:
- The solution begins by defining a struct type "Book" with the specified fields: Title, Author, and YearPublished.
- The "NewBook" function is implemented to create a new Book struct value by accepting the title, author, and year published as arguments and returning a Book struct with those values.
- The "PrintBookDetails" function takes a Book struct as a parameter and uses the "fmt.Printf" function to display the book's title, author, and year published.
- In the main function, "book1" is declared using the "NewBook" function, passing the title, author, and year published as arguments.
- "book2" is declared using a struct literal, providing values for title, author, and year published.
- Finally, the "PrintBookDetails" function is called twice, passing "book1" and "book2" as arguments to display their details.

When you run the program, it should output:
```
Title: The Great Gatsby
Author: F. Scott Fitzgerald
Year Published: 1925

Title: To Kill a Mockingbird
Author: Harper Lee
Year Published: 1960
```

This exercise helps reinforce the concepts of creating struct types, initializing struct values using functions and struct literals, and accessing struct fields for further operations.