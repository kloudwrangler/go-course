# Lesson: User Defined Types in Go

Objective: In this lesson, you will learn about user-defined types in the Go programming language. We will cover the declaration of struct types, creating struct values, pointers to structs, declaring new types based on base types, and understanding compiler errors. Additionally, we will explore an example of using user-defined types to implement a linked list.

1. Declaration of a Struct Type:
   - Go allows you to define your own custom types using structs. A struct is a composite data type that groups together zero or more values with different types.
   - The syntax to declare a struct type is as follows:
     ```go
     type StructName struct {
         field1 fieldType1
         field2 fieldType2
         // ...
     }
     ```

2. Declaration of a Variable of the Struct Type Set to Its Zero Value:
   - Once you have defined a struct type, you can declare variables of that type.
   - When a variable of a struct type is declared without explicitly initializing its fields, it is set to its zero value. For numeric types, the zero value is 0, for strings, it is an empty string, and so on.
   - Example:
     ```go
     type Person struct {
         name string
         age  int
     }
     
     var p Person // p is set to the zero value of the Person struct type
     ```

3. Declaration of a Variable of the Struct Type Using a Struct Literal:
   - You can also declare a variable of a struct type using a struct literal. A struct literal specifies values for the fields of the struct type within braces.
   - Example:
     ```go
     p := Person{name: "John", age: 30} // p is assigned a struct literal with specific field values
     ```

4. Creating a Struct Type Value Using a Struct Literal:
   - Apart from declaring variables, you can create struct values directly using a struct literal without assigning it to a variable.
   - Example:
     ```go
     john := Person{name: "John", age: 30} // Creating a Person struct value without assigning it to a variable
     ```

5. Creating a Struct Type Value Without Declaring the Field Names:
   - If you know the order of the fields in the struct type, you can create a struct value without explicitly declaring the field names. However, this approach is not recommended as it may lead to confusion.
   - Example:
     ```go
     p := Person{"John", 30} // Creating a Person struct value without declaring field names
     ```

6. Declaring Fields Based on Other Struct Types:
   - You can declare fields in a struct type based on other struct types. This allows you to compose complex data structures by combining multiple smaller struct types.
   - Example:
     ```go
     type Address struct {
         street  string
         city    string
         country string
     }
     
     type Person struct {
         name    string
         age     int
         address Address // Using the Address struct type as a field in the Person struct type
     }
     ```

7. Pointers to Structs:
   - Go supports pointers to structs, which allow you to indirectly access and modify the underlying struct values.
   - Example:
     ```go
     p := &Person{name: "John", age: 30} // p is a pointer to a Person struct
     ```

8. Declaration of a New Type Based on a Base Type:
   - Go allows you to declare a new type based on a base type to provide additional clarity or semantic meaning to your code. This is often useful when you want to add specific behaviors or methods to an existing type.
   -

 Example:
     ```go
     type Age int // Declaring a new type Age based on the int type
     ```

9. Compiler Error Assigning Value of Different Types:
   - When assigning values of different types to variables, the Go compiler will raise an error.
   - Example:
     ```go
     age := 30
     name := "John"
     age = name // Compiler error: cannot use name (type string) as type int in assignment
     ```

10. Actual Compiler Error:
    - It is important to understand and interpret the compiler error messages correctly when working with Go.
    - Example:
      ```
      ./main.go:7:9: cannot use name (type string) as type int in assignment
      ```

11. Example: Linked List:
    - Let's explore an example that demonstrates the usage of user-defined types in Go.
    - Implementing a linked list using structs and pointers can be a great example to illustrate these concepts.
    - Here's a simplified implementation of a linked list:
      ```go
      type Node struct {
          value int
          next  *Node
      }
      
      type LinkedList struct {
          head *Node
      }
      
      // Insert adds a new node with the given value at the end of the linked list.
      func (ll *LinkedList) Insert(value int) {
          newNode := &Node{value: value}
          if ll.head == nil {
              ll.head = newNode
          } else {
              current := ll.head
              for current.next != nil {
                  current = current.next
              }
              current.next = newNode
          }
      }
      
      // Display prints the values of the linked list.
      func (ll *LinkedList) Display() {
          current := ll.head
          for current != nil {
              fmt.Println(current.value)
              current = current.next
          }
      }
      ```

Additional Resources:
- The Go Tour: More Types - Pointers to Structs: [https://go.dev/tour/moretypes/4](https://go.dev/tour/moretypes/4)
- Go by Example: Structs: [https://gobyexample.com/structs](https://gobyexample.com/structs)
- A Tour of Go: [https://tour.golang.org/welcome/1](https://tour.golang.org/welcome/1)
- Go Documentation: [https://golang.org/doc/](https://golang.org/doc/)
- Effective Go: [https://golang.org/doc/effective_go.html](https://golang.org/doc/effective_go.html)

Note: It's always beneficial to practice writing code and experimenting with examples to solidify your understanding of user-defined types in Go.