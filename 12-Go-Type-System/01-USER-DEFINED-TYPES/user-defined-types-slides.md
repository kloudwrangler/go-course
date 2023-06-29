---
marp: true
theme: uncover

---
Title: User Defined Types in Go

---
Objective: Understand the concept of user-defined types in Go and how to work with them effectively.

---
## Struct Types
- Structs are composite data types that group together zero or more values with different types.
- Declaration syntax:
```go
type StructName struct {
    field1 fieldType1
    field2 fieldType2
    // ...
}
```

---
## Struct Variables
- Declare a variable of struct type and set it to its zero value:
```go
type Person struct {
    name string
    age  int
}

var p Person // p is set to the zero value of the Person struct type
```

---
## Struct Literals
- Declare a variable of struct type using a struct literal:
```go
p := Person{name: "John", age: 30} // p is assigned a struct literal with specific field values
```

---
## Creating Struct Values
- Create a struct value directly using a struct literal:
```go
john := Person{name: "John", age: 30} // Creating a Person struct value without assigning it to a variable
```

---
## Struct Value without Field Names
- Creating a struct value without declaring field names (using the order of fields):
```go
p := Person{"John", 30} // Creating a Person struct value without declaring field names
```

---
## Fields Based on Other Struct Types
- Declare fields based on other struct types to compose complex data structures:
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

---
## Pointers to Structs
- Use pointers to structs for indirect access and modification of underlying struct values:
```go
p := &Person{name: "John", age: 30} // p is a pointer to a Person struct
```

---
## New Types Based on Base Types
- Declare new types based on base types to provide clarity and semantic meaning:
```go
type Age int // Declaring a new type Age based on the int type
```

---
## Compiler Errors
- Compiler error when assigning values of different types:
```go
age := 30
name := "John"
age = name // Compiler error: cannot use name (type string) as type int in assignment
```

---
## Linked List Example
- Implementing a linked list using structs and pointers:
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

---
Additional Resources:
- The Go Tour: More Types - Pointers to Struct